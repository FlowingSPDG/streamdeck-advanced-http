package sdhttp

import (
	"context"
	"net/http"

	"github.com/FlowingSPDG/streamdeck"
)

const (
	// Action Perform action
	ActionPerform = "dev.flowingspdg.advancedhttp.perform"
	ActionDial    = "dev.flowingspdg.advancedhttp.dial"
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

	buttonAction := ret.sd.Action(ActionPerform)
	buttonAction.RegisterHandler(streamdeck.WillAppear, ret.ButtonWillAppearHandler)
	buttonAction.RegisterHandler(streamdeck.WillDisappear, ret.ButtonWillDisapperHandler)
	buttonAction.RegisterHandler(streamdeck.KeyDown, ret.KeyDownHandler)
	buttonAction.RegisterHandler(streamdeck.KeyUp, ret.KeyUpHandler)

	dialAction := ret.sd.Action(ActionDial)
	dialAction.RegisterHandler(streamdeck.WillAppear, ret.DialWillAppearHandler)
	dialAction.RegisterHandler(streamdeck.WillDisappear, ret.DialWillDisapperHandler)
	dialAction.RegisterHandler(streamdeck.DialRotate, ret.DialRotateHandler)
	dialAction.RegisterHandler(streamdeck.DialDown, ret.DialDownHandler)
	dialAction.RegisterHandler(streamdeck.DialUp, ret.DialUpHandler)
	dialAction.RegisterHandler(streamdeck.TouchTap, ret.TouchTapHandler)

	return ret
}

// Run Start client
func (s *SDHTTP) Run(ctx context.Context) error {
	return s.sd.Run(ctx)
}
