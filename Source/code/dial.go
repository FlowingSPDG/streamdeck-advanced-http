package sdhttp

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/FlowingSPDG/streamdeck"
)

// DialWillAppearHandler WillAppear handler
func (s *SDHTTP) DialWillAppearHandler(ctx context.Context, client *streamdeck.Client, event streamdeck.Event) error {
	return willAppearHandler[*DialPI](ctx, client, event)
}

// DialWillDisapperHandler WillAppear handler
func (s *SDHTTP) DialWillDisapperHandler(ctx context.Context, client *streamdeck.Client, event streamdeck.Event) error {
	return WillDisappearHandler[*DialPI](ctx, client, event)
}

// DialRotateHandler
func (s *SDHTTP) DialRotateHandler(ctx context.Context, client *streamdeck.Client, event streamdeck.Event) error {
	payload := streamdeck.DialRotatePayload[DialPI]{}
	if err := json.Unmarshal(event.Payload, &payload); err != nil {
		msg := fmt.Sprintf("Failed to unmarshal KeyDown event payload: %s", err)
		client.LogMessage(ctx, msg)
		client.ShowAlert(ctx)
		return err
	}

	url := ""
	if payload.Ticks > 0 {
		url = payload.Settings.URLRight
	} else {
		url = payload.Settings.URLLeft
	}

	r := request{
		body:              payload.Settings.Body,
		method:            payload.Settings.Method,
		url:               url,
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

// DialDownHandler
func (s *SDHTTP) DialDownHandler(ctx context.Context, client *streamdeck.Client, event streamdeck.Event) error {
	payload := streamdeck.DialDownPayload[DialPI]{}
	if err := json.Unmarshal(event.Payload, &payload); err != nil {
		msg := fmt.Sprintf("Failed to unmarshal KeyDown event payload: %s", err)
		client.LogMessage(ctx, msg)
		client.ShowAlert(ctx)
		return err
	}

	r := request{
		body:              payload.Settings.Body,
		method:            payload.Settings.Method,
		url:               payload.Settings.URLPush,
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
