//+build consumer

package client_test

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/pact-foundation/pact-go/dsl"
	"github.com/stretchr/testify/assert"
	"gitlab.com/modanisatech/marketplace/service-template/internal/client"
	"go.uber.org/zap"
)

func TestConsumerGetExistingOrder(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}

	orderID := 1

	pact := createPactConfig("ServiceTemplate", "AnotherService")
	defer pact.Teardown()
	pact.Setup(true)

	pact.
		AddInteraction().
		Given(fmt.Sprintf("I have a order with id %d", orderID)).
		UponReceiving("A request for get a order").
		WithRequest(dsl.Request{
			Method: http.MethodGet,
			Path:   dsl.String(fmt.Sprintf("/orders/%d", orderID)),
		}).
		WillRespondWith(dsl.Response{
			Status: http.StatusOK,
			Headers: dsl.MapMatcher{
				"Content-Type": dsl.String("application/json"),
			},
			Body: map[string]interface{}{
				"id":     dsl.Like(orderID),
				"status": dsl.Like("delivered"),
			},
		})

	err := pact.Verify(func() error {
		orderClient := client.New(fmt.Sprintf("http://localhost:%d", pact.Server.Port), zap.NewNop())
		_, err := orderClient.GetOrder(context.Background(), orderID)
		return err
	})

	assert.Nil(t, err)
}

func TestGetNotExistingOrder(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}

	orderID := 1

	pact := createPactConfig("ServiceTemplate", "AnotherService")
	defer pact.Teardown()
	pact.Setup(true)

	pact.
		AddInteraction().
		Given(fmt.Sprintf("I have no order with id %d", orderID)).
		UponReceiving("A request for get a order").
		WithRequest(dsl.Request{
			Method: http.MethodGet,
			Path:   dsl.String(fmt.Sprintf("/orders/%d", orderID)),
		}).
		WillRespondWith(dsl.Response{
			Status: http.StatusNotFound,
		})

	err := pact.Verify(func() error {
		orderClient := client.New(fmt.Sprintf("http://localhost:%d", pact.Server.Port), zap.NewNop())
		_, err := orderClient.GetOrder(context.Background(), orderID)
		assert.Equal(t, client.ErrOrderNotExists, err)
		return nil
	})

	assert.Nil(t, err)
}

func createPactConfig(consumer, provider string) *dsl.Pact {
	return &dsl.Pact{
		Host:                     "0.0.0.0",
		Consumer:                 consumer,
		Provider:                 provider,
		DisableToolValidityCheck: true,
		LogDir:                   "../log",
		PactDir:                  "../pacts",
		PactFileWriteMode:        "merge",
	}
}
