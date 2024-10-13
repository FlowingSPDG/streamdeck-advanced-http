package sdhttp

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/FlowingSPDG/streamdeck"
)

// ButtonWillAppearHandler WillAppear handler
func (s *SDHTTP) ButtonWillAppearHandler(ctx context.Context, client *streamdeck.Client, event streamdeck.Event) error {
	return willAppearHandler[*KeyDownPI](ctx, client, event)
}

// ButtonWillDisapperHandler WillAppear handler
func (s *SDHTTP) ButtonWillDisapperHandler(ctx context.Context, client *streamdeck.Client, event streamdeck.Event) error {
	return WillDisappearHandler[*KeyDownPI](ctx, client, event)
}

// KeyDownHandler Handles "KeyDown" action
func (s *SDHTTP) KeyDownHandler(ctx context.Context, client *streamdeck.Client, event streamdeck.Event) error {
	payload := streamdeck.KeyDownPayload[KeyDownPI]{}
	if err := json.Unmarshal(event.Payload, &payload); err != nil {
		msg := fmt.Sprintf("Failed to unmarshal KeyDown event payload: %s", err)
		client.LogMessage(ctx, msg)
		client.ShowAlert(ctx)
		return err
	}

	r := request{
		body:              payload.Settings.Body,
		method:            payload.Settings.Method,
		url:               payload.Settings.URL,
		showAlert:         payload.Settings.ShowAlert,
		basicAuthID:       payload.Settings.BasicAuthID,
		basicAuthPassword: payload.Settings.BasicAuthPassword,
	}

	if err := s.do(ctx, client, r); err != nil {
		return err
	}

	msg := fmt.Sprintf("Request succeeded :%v", payload.Settings)
	client.LogMessage(ctx, msg)

	if payload.Settings.ShowOK {
		client.ShowOk(ctx)
	}
	return nil
}
