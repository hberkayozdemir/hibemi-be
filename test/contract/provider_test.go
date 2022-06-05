package contract

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/joho/godotenv"
	"github.com/pact-foundation/pact-go/dsl"
	"github.com/pact-foundation/pact-go/types"
	"github.com/pact-foundation/pact-go/utils"
	"gitlab.com/modanisatech/marketplace/service-template/internal/foo"
	"gitlab.com/modanisatech/marketplace/service-template/pkg/server"
	"go.uber.org/zap"
)

const providerName = "ServiceTemplate"

func TestProvider(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}

	port, _ := utils.GetFreePort()

	srv := server.New(fmt.Sprintf(":%d", port), []server.Handler{foo.Handler{}}, zap.NewExample())

	go srv.Run()

	verifyProvider(t, providerName, port, types.StateHandlers{})
}

func verifyProvider(t *testing.T, providerName string, providerPort int, stateHandlers types.StateHandlers) {
	ci := os.Getenv("CI")
	if ci == "" {
		_ = godotenv.Load("../../.dev/.env")
	}

	brokerBaseURL := os.Getenv("PACT_BROKER_BASE_URL")
	brokerUsername := os.Getenv("PACT_BROKER_USERNAME")
	brokerPassword := os.Getenv("PACT_BROKER_PASSWORD")
	consumerName := os.Getenv("PACT_CONSUMER_NAME")
	consumerVersion := os.Getenv("PACT_CONSUMER_VERSION")
	consumerTag := os.Getenv("PACT_CONSUMER_TAG")
	providerVersion := os.Getenv("PACT_PROVIDER_VERSION")

	fmt.Println(strings.Repeat("-", 20))
	fmt.Println("consumer-name: ", consumerName)
	fmt.Println("provider-version: ", providerVersion)
	fmt.Println("base-url: ", brokerBaseURL)
	fmt.Println("username: ", brokerUsername)
	fmt.Println("tag: ", consumerTag)
	fmt.Println(strings.Repeat("-", 20))

	pact := &dsl.Pact{
		Host:                     "127.0.0.1",
		DisableToolValidityCheck: true,
	}

	verifyRequest := types.VerifyRequest{
		ProviderBaseURL:            fmt.Sprintf("http://127.0.0.1:%d", providerPort),
		ProviderVersion:            providerVersion,
		BrokerURL:                  brokerBaseURL,
		BrokerUsername:             brokerUsername,
		BrokerPassword:             brokerPassword,
		Tags:                       []string{consumerTag},
		StateHandlers:              stateHandlers,
		PublishVerificationResults: true,
	}

	if ci == "" {
		fmt.Println("Not publishing pact verification results because not running in CI.")
		verifyRequest.PublishVerificationResults = false
	}

	// validate specific contract if consumerVersionSpecified
	if consumerName != "" {
		pactURL := ""
		if consumerVersion != "" {
			pactURL = fmt.Sprintf("%s/pacts/provider/%s/consumer/%s/version/%s.json", brokerBaseURL, providerName, consumerName, consumerVersion)
		} else {
			pactURL = fmt.Sprintf("%s/pacts/provider/%s/consumer/%s/latest/master.json", brokerBaseURL, providerName, consumerName)
		}

		fmt.Println("Running provider test just to validate following consumer contract..")
		fmt.Println("consumer-name: ", consumerName)
		fmt.Println("consumer-tag: ", consumerTag)
		fmt.Println("pact-url: ", pactURL)
		verifyRequest.BrokerURL = ""
		verifyRequest.Tags = []string{}
		verifyRequest.PactURLs = []string{pactURL}
		verifyRequest.FailIfNoPactsFound = true
	} else {
		fmt.Println("Running " + consumerTag + " provider tests for all contracts..")
		pact.Provider = providerName
		verifyRequest.Tags = []string{consumerTag}
	}

	verifyResponses, err := pact.VerifyProvider(t, verifyRequest)
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(len(verifyResponses), " pact tests runned.")
}
