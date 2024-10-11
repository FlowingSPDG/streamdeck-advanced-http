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

func willAppearHandler[T propertyInspector](ctx context.Context, client *streamdeck.Client, event streamdeck.Event) error {
	payload := streamdeck.WillAppearPayload[T]{}
	if err := json.Unmarshal(event.Payload, &payload); err != nil {
		msg := fmt.Sprintf("Failed to unmarshal WillAppear event payload: %s", err)
		client.LogMessage(ctx, msg)
		client.ShowAlert(ctx)
		return err
	}

	msg := fmt.Sprintf("Context %s WillAppear with settings :%v", event.Context, payload.Settings)
	client.LogMessage(ctx, msg)
	return nil
}

func WillDisappearHandler[T propertyInspector](ctx context.Context, client *streamdeck.Client, event streamdeck.Event) error {
	payload := streamdeck.WillDisappearPayload[T]{}
	if err := json.Unmarshal(event.Payload, &payload); err != nil {
		msg := fmt.Sprintf("Failed to unmarshal WillAppear event payload: %s", err)
		client.LogMessage(ctx, msg)
		client.ShowAlert(ctx)
		return err
	}

	msg := fmt.Sprintf("Deleting setting for context %s", event.Context)
	client.LogMessage(ctx, msg)
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
