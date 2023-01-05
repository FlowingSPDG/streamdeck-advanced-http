package sdhttp

import (
	"context"
	"net/http"

	"github.com/FlowingSPDG/streamdeck"
)

const (
	// Action Perform action
	Action = "dev.flowingspdg.advancedhttp.perform"
)

// SDHTTP StreamDeck client
type SDHTTP struct {
	sd *streamdeck.Client
	ht *http.Client
}

// NewSDHTTP Get New StreamDeck plugin instance pointer
func NewSDHTTP(ctx context.Context, params streamdeck.RegistrationParams) *SDHTTP {
	ret := &SDHTTP{
		sd: streamdeck.NewClient(ctx, params),
		ht: http.DefaultClient,
	}

	action := ret.sd.Action(Action)
	action.RegisterHandler(streamdeck.WillAppear, ret.WillAppearHandler)
	action.RegisterHandler(streamdeck.KeyDown, ret.KeyDownHandler)

	return ret
}

// Run Start client
func (s *SDHTTP) Run(ctx context.Context) error {
	return s.sd.Run(ctx)
}
