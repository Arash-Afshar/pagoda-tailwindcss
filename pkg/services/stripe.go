package services

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"time"

	stripe "github.com/stripe/stripe-go/v81"
	"github.com/stripe/stripe-go/v81/client"
)

type stripeKVData map[string]*PaymentData

type StripeClient struct {
	c             *client.API
	webhookSecret string
}

type PaymentData struct {
	PaymentID         string
	Status            stripe.PaymentIntentStatus
	Amount            int64
	PriceID           string
	CheckoutSessionID string
}

func (p stripeKVData) MarshalBinary() ([]byte, error) {
	json, err := json.Marshal(p)
	if err != nil {
		return nil, err
	}
	return json, nil
}

func (p *stripeKVData) UnmarshalBinary(data []byte) error {
	err := json.Unmarshal(data, p)
	if err != nil {
		return err
	}
	return nil
}

var allowedEvents = []stripe.EventType{
	stripe.EventTypeCheckoutSessionCompleted,
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

func (s *StripeClient) GetCustomer(ctx context.Context, cache *CacheClient, userId int, userEmail string) (*stripe.Customer, error) {
	stripeCustomerId, err := cache.Get().
		Key(kvUserKey(userId)).
		Fetch(ctx)
	if err != nil && !errors.Is(err, ErrCacheMiss) {
		return nil, fmt.Errorf("getting stripe customer from cache: %w", err)
	}
	if (err != nil && errors.Is(err, ErrCacheMiss)) || stripeCustomerId == nil {
		params := &stripe.CustomerParams{
			Email: stripe.String(userEmail),
			Metadata: map[string]string{
				"user_id": strconv.Itoa(userId),
			},
		}
		c, err := s.c.Customers.New(params)
		if err != nil {
			return nil, fmt.Errorf("creating stripe customer: %w", err)
		}
		stripeCustomerId = c.ID

		err = cache.Set().
			Key(kvUserKey(userId)).
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

func (s *StripeClient) paymentDataFromCheckoutSession(sessionID string) (*PaymentData, error) {
	params := &stripe.CheckoutSessionParams{
		Params: stripe.Params{
			Expand: []*string{
				stripe.String("line_items"),
			},
		},
	}

	sess, err := s.c.CheckoutSessions.Get(sessionID, params)
	if err != nil {
		return nil, err
	}

	if sess.LineItems != nil && len(sess.LineItems.Data) > 0 {
		p := sess.PaymentIntent
		return &PaymentData{
			PaymentID:         p.ID,
			Status:            p.Status,
			Amount:            p.Amount,
			PriceID:           sess.LineItems.Data[0].Price.ID,
			CheckoutSessionID: sess.ID,
		}, nil
	}
	return nil, nil
}

func (s *StripeClient) SyncStripeDataToKV(ctx context.Context, cache *CacheClient, stripeCustomerId string, paymentData *PaymentData) (stripeKVData, error) {
	currSer, err := cache.Get().Key(kvStripeCustomerKey(stripeCustomerId)).Fetch(ctx)
	if err != nil {
		return nil, fmt.Errorf("getting stripe payment data from cache: %w", err)
	}

	currSerStr, ok := currSer.(string)
	if !ok {
		return nil, fmt.Errorf("stripe payment data is not a string")
	}

	var newData stripeKVData
	err = json.Unmarshal([]byte(currSerStr), &newData)
	if err != nil {
		return nil, fmt.Errorf("marshalling stripe payment data: %w", err)
	}

	newData[paymentData.CheckoutSessionID] = paymentData

	err = cache.Set().
		Key(kvStripeCustomerKey(stripeCustomerId)).
		Data(newData).
		Expiration(time.Hour * 24).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("saving stripe payment data to cache: %w", err)
	}
	return newData, nil
}

func (s *StripeClient) GetStripeDataFromKV(ctx context.Context, cache *CacheClient, customerId string) (stripeKVData, error) {
	paymentsDataJson, err := cache.Get().
		Key(kvStripeCustomerKey(customerId)).
		Fetch(ctx)
	if err != nil {
		return nil, fmt.Errorf("getting stripe payment data from cache: %w", err)
	}
	paymentsDataJsonBytes, ok := paymentsDataJson.(string)
	if !ok {
		return nil, fmt.Errorf("stripe payment data is not a string")
	}
	var paymentsData stripeKVData
	err = json.Unmarshal([]byte(paymentsDataJsonBytes), &paymentsData)
	if err != nil {
		return nil, fmt.Errorf("unmarshalling stripe payment data: %w", err)
	}
	return paymentsData, nil
}

func (s *StripeClient) Success(ctx context.Context, cache *CacheClient, userId int, sessionID string) error {
	stripeCustomerId, err := cache.Get().
		Key(kvUserKey(userId)).
		Fetch(ctx)
	if err != nil {
		return fmt.Errorf("getting stripe customer from cache: %w", err)
	}

	paymentData, err := s.paymentDataFromCheckoutSession(sessionID)
	if err != nil {
		return fmt.Errorf("getting payment data from checkout session: %w", err)
	}

	_, err = s.SyncStripeDataToKV(ctx, cache, stripeCustomerId.(string), paymentData)
	if err != nil {
		return fmt.Errorf("syncing stripe data to cache: %w", err)
	}
	return nil
}

// func (s *StripeClient) WebhookHandler(ctx context.Context, cache *CacheClient) func(w http.ResponseWriter, req *http.Request) {
// 	return func(w http.ResponseWriter, req *http.Request) {
// 		const MaxBodyBytes = int64(65536)
// 		req.Body = http.MaxBytesReader(w, req.Body, MaxBodyBytes)
// 		payload, err := io.ReadAll(req.Body)
// 		if err != nil {
// 			fmt.Fprintf(os.Stderr, "Error reading request body: %v\n", err)
// 			w.WriteHeader(http.StatusServiceUnavailable)
// 			return
// 		}
//
// 		event, err := webhook.ConstructEvent(payload, req.Header.Get("Stripe-Signature"),
// 			s.webhookSecret)
// 		if err != nil {
// 			fmt.Fprintf(os.Stderr, "Failed to parse webhook body json: %v\n", err.Error())
// 			w.WriteHeader(http.StatusBadRequest)
// 			return
// 		}
// 		if err := s.processEvent(ctx, cache, event); err != nil {
// 			fmt.Fprintf(os.Stderr, "Error processing event: %v\n", err.Error())
// 			w.WriteHeader(http.StatusInternalServerError)
// 			return
// 		}
// 		w.WriteHeader(http.StatusOK)
// 	}
// }
//
// func (s *StripeClient) processEvent(ctx context.Context, cache *CacheClient, event stripe.Event) error {
// 	if !slices.Contains(allowedEvents, event.Type) {
// 		return fmt.Errorf("event type %s is not allowed", event.Type)
// 	}
// 	customerId := event.Data.Object["customer"].(string)
// 	if customerId == "" {
// 		return fmt.Errorf("customer ID not provided for event %s", event.Type)
// 	}
//
// 	_, err := s.SyncStripeDataToKV(ctx, cache, customerId)
// 	if err != nil {
// 		return fmt.Errorf("syncing stripe data to cache: %w", err)
// 	}
// 	return nil
// }
//
