package services

import (
	// "context"
	// "fmt"
	"testing"
	// "github.com/Arash-Afshar/pagoda-tailwindcss/config"
	// "github.com/Arash-Afshar/pagoda-tailwindcss/pkg/tests"
	// "github.com/stretchr/testify/require"
)

// Run `make stripe-mock` to start the stripe mock server
// TODO: disabling for now until an easy way of testing in CI is found

func TestCheckoutSession(t *testing.T) {
	// testKey := "sk_test_12345"
	// url := "http://localhost:12111"
	// successUrl := "/success"
	// cancelUrl := "/cancel"
	// webhookSecret := "whsec_12345"

	// client := NewStripeClient(testKey, url, webhookSecret)

	// // Set the environment to test
	// config.SwitchEnvironment(config.EnvTest)

	// // Create a new container
	// c := NewContainer()

	// // Create a web context
	// ctx, _ := tests.NewContext(c.Web, "/")
	// tests.InitSession(ctx)

	// // Create a test user
	// var err error
	// user, err := tests.CreateUser(c.ORM)
	// require.NoError(t, err)

	// customer, err := client.GetCustomer(context.Background(), c.Cache, user.ID, user.Email)
	// require.NoError(t, err)
	// fmt.Println(customer)

	// session, err := client.CheckoutSession(context.Background(), successUrl, cancelUrl, customer.ID, "price_123", 1)
	// require.NoError(t, err)
	// fmt.Println(session)
}
