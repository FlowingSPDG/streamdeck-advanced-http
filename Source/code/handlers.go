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

	if err := s.do(ctx, client, payload.Settings); err != nil {
		return err
	}

	msg := fmt.Sprintf("Request succeeded :%v", payload.Settings)
	client.LogMessage(ctx, msg)

	if payload.Settings.ShowOK {
		client.ShowOk(ctx)
	}
	return nil
}

func (s *SDHTTP) do(ctx context.Context, client *streamdeck.Client, setetings PI) error {
	// Perform HTTP Request
	// リクエスト定義
	buf := bytes.NewBufferString(setetings.Body)
	req, err := http.NewRequest(setetings.Method, setetings.URL, buf)
	if err != nil {
		msg := fmt.Sprintf("Failed to generate request: %s", err)
		client.LogMessage(ctx, msg)
		if setetings.ShowAlert {
			client.ShowAlert(ctx)
		}
		return err
	}
	// 認証情報をセット
	if setetings.BasicAuthID != "" && setetings.BasicAuthPassword != "" {
		req.SetBasicAuth(setetings.BasicAuthID, setetings.BasicAuthPassword)
	}

	// リクエスト実行
	resp, err := s.ht.Do(req)
	if err != nil {
		msg := fmt.Sprintf("Failed to perform request: %s", err)
		client.LogMessage(ctx, msg)
		if setetings.ShowAlert {
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
