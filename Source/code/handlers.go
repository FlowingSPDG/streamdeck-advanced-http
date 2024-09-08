package sdhttp

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/FlowingSPDG/streamdeck"
)

// ButtonWillAppearHandler WillAppear handler
func (s *SDHTTP) ButtonWillAppearHandler(ctx context.Context, client *streamdeck.Client, event streamdeck.Event) error {
	return willAppearHandler[*KeyDownPI](ctx, client, event)
}

func willAppearHandler[T propertyInspector](ctx context.Context, client *streamdeck.Client, event streamdeck.Event) error {
	payload := streamdeck.WillAppearPayload[T]{}
	if err := json.Unmarshal(event.Payload, &payload); err != nil {
		msg := fmt.Sprintf("Failed to unmarshal WillAppear event payload: %s", err)
		client.LogMessage(ctx, msg)
		client.ShowAlert(ctx)
		return err
	}

	if payload.Settings.IsDefault() {
		payload.Settings.Initialize()
		client.SetSettings(ctx, payload.Settings)
	}

	msg := fmt.Sprintf("Context %s WillAppear with settings :%v", event.Context, payload.Settings)
	client.LogMessage(ctx, msg)
	return nil
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

// DialWillAppearHandler WillAppear handler
func (s *SDHTTP) DialWillAppearHandler(ctx context.Context, client *streamdeck.Client, event streamdeck.Event) error {
	return willAppearHandler[*DialPI](ctx, client, event)
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

type request struct {
	body   string
	method string
	url    string

	showAlert         bool
	basicAuthID       string
	basicAuthPassword string
}

// do perform HTTP request
func (s *SDHTTP) do(ctx context.Context, client *streamdeck.Client, r request) error {
	// Perform HTTP Request
	// リクエスト定義
	buf := bytes.NewBufferString(r.body)
	req, err := http.NewRequest(r.method, r.url, buf)
	if err != nil {
		msg := fmt.Sprintf("Failed to generate request: %s", err)
		client.LogMessage(ctx, msg)
		if r.showAlert {
			client.ShowAlert(ctx)
		}
		return err
	}
	// 認証情報をセット
	if r.basicAuthID != "" && r.basicAuthPassword != "" {
		req.SetBasicAuth(r.basicAuthID, r.basicAuthPassword)
	}

	// リクエスト実行
	resp, err := s.ht.Do(req)
	if err != nil {
		msg := fmt.Sprintf("Failed to perform request: %s", err)
		client.LogMessage(ctx, msg)
		if r.showAlert {
			client.ShowAlert(ctx)
		}
		return err
	}
	defer resp.Body.Close()

	// discard buffer
	if _, err := io.Copy(io.Discard, resp.Body); err != nil {
		return err
	}
	return nil
}
