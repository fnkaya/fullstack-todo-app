package pact

import (
	"backend/server"
	"fmt"
	"github.com/pact-foundation/pact-go/dsl"
	"github.com/pact-foundation/pact-go/types"
	"github.com/pact-foundation/pact-go/utils"
	"testing"
)

func TestProvider(t *testing.T) {
	port, _ := utils.GetFreePort()
	httpServer := server.NewServer(port)
	go httpServer.StartServer()

	pact := dsl.Pact{
		Host:                     "127.0.0.1",
		Provider:                 "todos-be",
		Consumer:                 "todos-fe",
		DisableToolValidityCheck: true,
	}

	request := types.VerifyRequest{
		ProviderBaseURL: fmt.Sprintf("http://localhost:%d", port),
		PactURLs:        []string{"https://fnkaya.pactflow.io/pacts/provider/todos-be/consumer/todos-fe/version/1.0.0"},
		BrokerToken:     "-oJHQxYzGtNuI9cZCuOogA",
	}

	verifyResponses, err := pact.VerifyProvider(t, request)

	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(len(verifyResponses), "pact tests run")
}
