package sdnewtek

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/FlowingSPDG/streamdeck"
)

// WillAppearHandler WillAppear handler
func (s *SDHTTP) WillAppearHandler(ctx context.Context, client *streamdeck.Client, event streamdeck.Event) error {
	payload := streamdeck.WillAppearPayload[PI]{}
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

// KeyDownHandler Perform HTTP Request
func (s *SDHTTP) KeyDownHandler(ctx context.Context, client *streamdeck.Client, event streamdeck.Event) error {
	payload := streamdeck.KeyDownPayload[PI]{}
	if err := json.Unmarshal(event.Payload, &payload); err != nil {
		msg := fmt.Sprintf("Failed to unmarshal KeyDown event payload: %s", err)
		client.LogMessage(ctx, msg)
		client.ShowAlert(ctx)
		return err
	}

	// Perform HTTP Request
	hc := &http.Client{Timeout: time.Second}
	// リクエスト定義
	buf := bytes.NewBufferString(payload.Settings.Body)
	req, err := http.NewRequest(payload.Settings.Method, payload.Settings.URL, buf)
	if err != nil {
		msg := fmt.Sprintf("Failed to generate request: %s", err)
		client.LogMessage(ctx, msg)
		client.ShowAlert(ctx)
		return err
	}
	// 認証情報をセット
	if payload.Settings.BasicAuthID != "" && payload.Settings.BasicAuthPassword != "" {
		req.SetBasicAuth(payload.Settings.BasicAuthID, payload.Settings.BasicAuthPassword)
	}

	// リクエスト実行
	resp, err := hc.Do(req)
	if err != nil {
		msg := fmt.Sprintf("Failed to perform request: %s", err)
		client.LogMessage(ctx, msg)
		client.ShowAlert(ctx)
		return err
	}
	defer resp.Body.Close()

	// discard buffer
	io.Copy(io.Discard, resp.Body)

	msg := fmt.Sprintf("Request succeeded :%v", payload.Settings)
	client.LogMessage(ctx, msg)

	return client.ShowOk(ctx)
}
