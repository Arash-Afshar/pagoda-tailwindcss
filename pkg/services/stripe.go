package services

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"slices"
	"strconv"
	"time"

	"github.com/Arash-Afshar/pagoda-tailwindcss/ent"
	stripe "github.com/stripe/stripe-go/v81"
	"github.com/stripe/stripe-go/v81/client"
	"github.com/stripe/stripe-go/v81/webhook"
)

type StripeClient struct {
	c             *client.API
	webhookSecret string
}

type SubscriptionData struct {
	SubscriptionID     string
	Status             stripe.SubscriptionStatus
	PriceID            string
	CurrentPeriodEnd   int64
	CurrentPeriodStart int64
	CancelAtPeriodEnd  bool
	PaymentMethod      *stripe.PaymentMethod
}

var allowedEvents = []stripe.EventType{
	stripe.EventTypeCheckoutSessionCompleted,
	stripe.EventTypeCustomerSubscriptionCreated,
	stripe.EventTypeCustomerSubscriptionUpdated,
	stripe.EventTypeCustomerSubscriptionDeleted,
	stripe.EventTypeCustomerSubscriptionPaused,
	stripe.EventTypeCustomerSubscriptionResumed,
	stripe.EventTypeCustomerSubscriptionPendingUpdateApplied,
	stripe.EventTypeCustomerSubscriptionPendingUpdateExpired,
	stripe.EventTypeCustomerSubscriptionTrialWillEnd,
	stripe.EventTypeInvoicePaid,
	stripe.EventTypeInvoicePaymentFailed,
	stripe.EventTypeInvoicePaymentActionRequired,
	stripe.EventTypeInvoiceUpcoming,
	stripe.EventTypeInvoiceMarkedUncollectible,
	stripe.EventTypeInvoicePaymentSucceeded,
	stripe.EventTypePaymentIntentSucceeded,
	stripe.EventTypePaymentIntentPaymentFailed,
	stripe.EventTypePaymentIntentCanceled,
}

func kvUserKey(userId int) string {
	return "stripe:user:" + strconv.Itoa(userId)
}

func kvStripeCustomerKey(customerId string) string {
	return "stripe:customer:" + customerId
}

func NewStripeClient(key, url, webhookSecret string) *StripeClient {
	config := &stripe.BackendConfig{}
	if url != "" {
		config.URL = stripe.String(url)
	}
	backends := &stripe.Backends{
		API: stripe.GetBackendWithConfig(stripe.APIBackend, config),
	}
	c := client.New(key, backends)
	return &StripeClient{c: c, webhookSecret: webhookSecret}
}

func (s *StripeClient) GetCustomer(ctx context.Context, cache *CacheClient, user *ent.User) (*stripe.Customer, error) {

	stripeCustomerId, err := cache.Get().
		Key(kvUserKey(user.ID)).
		Fetch(ctx)
	if err != nil && !errors.Is(err, ErrCacheMiss) {
		return nil, fmt.Errorf("getting stripe customer from cache: %w", err)
	}
	if (err != nil && errors.Is(err, ErrCacheMiss)) || stripeCustomerId == nil {
		params := &stripe.CustomerParams{
			Email: stripe.String(user.Email),
			Metadata: map[string]string{
				"user_id": strconv.Itoa(user.ID),
			},
		}
		c, err := s.c.Customers.New(params)
		if err != nil {
			return nil, fmt.Errorf("creating stripe customer: %w", err)
		}
		stripeCustomerId = c.ID

		err = cache.Set().
			Key(kvUserKey(user.ID)).
			Data(stripeCustomerId).
			Expiration(time.Hour * 24).
			Save(ctx)
		if err != nil {
			return nil, fmt.Errorf("saving stripe customer to cache: %w", err)
		}
		return c, nil
	} else {
		params := &stripe.CustomerParams{}
		c, err := s.c.Customers.Get(stripeCustomerId.(string), params)
		if err != nil {
			return nil, fmt.Errorf("getting stripe customer: %w", err)
		}
		return c, nil
	}
}

func (s *StripeClient) CheckoutSession(ctx context.Context, successUrl, cancelUrl, stripeCustomerId, priceId string, quantity int) (*stripe.CheckoutSession, error) {
	params := &stripe.CheckoutSessionParams{
		Customer:   stripe.String(stripeCustomerId),
		SuccessURL: stripe.String(successUrl),
		CancelURL:  stripe.String(cancelUrl),
		LineItems: []*stripe.CheckoutSessionLineItemParams{
			{
				Price:    stripe.String(priceId),
				Quantity: stripe.Int64(int64(quantity)),
			},
		},
		Mode: stripe.String(string(stripe.CheckoutSessionModePayment)),
	}
	session, err := s.c.CheckoutSessions.New(params)
	if err != nil {
		return nil, fmt.Errorf("creating stripe checkout session: %w", err)
	}
	return session, nil
}

func (s *StripeClient) SyncStripeDataToKV(ctx context.Context, cache *CacheClient, stripeCustomerId string) (*SubscriptionData, error) {
	// Fetch latest subscription data from Stripe
	subscriptions := s.c.Subscriptions.List(&stripe.SubscriptionListParams{
		Customer: stripe.String(stripeCustomerId),
		Status:   stripe.String("all"),
		Expand:   stripe.StringSlice([]string{"data.default_payment_method"}),
	})

	// Assume there's only one subscription
	sub := subscriptions.Subscription()
	count := 0
	for subscriptions.Next() {
		count++
	}
	if count != 1 {
		return nil, fmt.Errorf("expected 1 subscription, got %d", count)
	}

	subData := &SubscriptionData{
		SubscriptionID:     sub.ID,
		Status:             sub.Status,
		PriceID:            sub.Items.Data[0].Price.ID,
		CurrentPeriodEnd:   sub.CurrentPeriodEnd,
		CurrentPeriodStart: sub.CurrentPeriodStart,
		CancelAtPeriodEnd:  sub.CancelAtPeriodEnd,
		PaymentMethod:      sub.DefaultPaymentMethod,
	}
	err := cache.Set().
		Key(kvStripeCustomerKey(stripeCustomerId)).
		Data(subData).
		Expiration(time.Hour * 24).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("saving stripe subscription to cache: %w", err)
	}
	return subData, nil
}

func (s *StripeClient) Success(ctx context.Context, cache *CacheClient, user *ent.User) error {
	stripeCustomerId, err := cache.Get().
		Key(kvUserKey(user.ID)).
		Fetch(ctx)
	if err != nil {
		return fmt.Errorf("getting stripe customer from cache: %w", err)
	}

	_, err = s.SyncStripeDataToKV(ctx, cache, stripeCustomerId.(string))
	if err != nil {
		return fmt.Errorf("syncing stripe data to cache: %w", err)
	}
	return nil
}

func (s *StripeClient) WebhookHandler(ctx context.Context, cache *CacheClient) func(w http.ResponseWriter, req *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		const MaxBodyBytes = int64(65536)
		req.Body = http.MaxBytesReader(w, req.Body, MaxBodyBytes)
		payload, err := io.ReadAll(req.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading request body: %v\n", err)
			w.WriteHeader(http.StatusServiceUnavailable)
			return
		}

		event, err := webhook.ConstructEvent(payload, req.Header.Get("Stripe-Signature"),
			s.webhookSecret)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to parse webhook body json: %v\n", err.Error())
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if err := s.processEvent(ctx, cache, event); err != nil {
			fmt.Fprintf(os.Stderr, "Error processing event: %v\n", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}

func (s *StripeClient) processEvent(ctx context.Context, cache *CacheClient, event stripe.Event) error {
	if !slices.Contains(allowedEvents, event.Type) {
		return fmt.Errorf("event type %s is not allowed", event.Type)
	}
	customerId := event.Data.Object["customer"].(string)
	if customerId == "" {
		return fmt.Errorf("customer ID not provided for event %s", event.Type)
	}

	_, err := s.SyncStripeDataToKV(ctx, cache, customerId)
	if err != nil {
		return fmt.Errorf("syncing stripe data to cache: %w", err)
	}
	return nil
}
