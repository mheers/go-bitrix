package main

import (
	"log"

	"github.com/nightwriter/go-bitrix/client"
	"github.com/nightwriter/go-bitrix/types"
)

func main() {
	c, err := client.NewEnvClientWithWebhookAuth()

	if err != nil {
		log.Fatalf("Can't create client: %s", err)
	}

	c.SetInsecureSSL(true)
	c.SetDebug(true)

	resp, err := c.Methods(&types.MethodsRequest{
		Full:  true,
		Scope: "landing",
	})

	if err != nil {
		log.Fatalf("Request error: %s", err)
	}

	log.Print(resp)
}
