package main

import (
	"context"
	_ "embed"
	"log"
	"os"

	"github.com/FlowingSPDG/streamdeck"

	sdhttp "github.com/FlowingSPDG/streamdeck-advanced-http/Source/code"
)

func main() {
	ctx := context.Background()
	log.Println("Starting...")
	if err := run(ctx); err != nil {
		log.Fatalf("%v\n", err)
	}
}

func run(ctx context.Context) error {
	params, err := streamdeck.ParseRegistrationParams(os.Args)
	if err != nil {
		return err
	}

	client := sdhttp.NewSDHTTP(ctx, params)

	return client.Run(ctx)
}
