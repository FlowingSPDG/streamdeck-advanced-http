package sdhttp

import (
	"net/http"
	"reflect"
)

type propertyInspector interface {
	IsDefault() bool
	Initialize()
}

type pi struct {
	Method            string `json:"method"`
	Body              string `json:"body"`
	BasicAuthID       string `json:"basic_auth_id"`
	BasicAuthPassword string `json:"basic_auth_password"`
	AuthHeader        string `json:"auth_header"`

	ShowOK    bool `json:"show_ok"`
	ShowAlert bool `json:"show_alert"`
}

// KeyDownPI PropertyInspector for KeyDown
type KeyDownPI struct {
	pi
	URL string `json:"url"`
}

// IsDefault Check if its default
func (p KeyDownPI) IsDefault() bool {
	return reflect.ValueOf(p).IsZero()
}

// Initialize PI
func (p *KeyDownPI) Initialize() {
	p.Method = http.MethodGet
	p.URL = "https://www.elgato.com"
	p.Body = ""
	p.BasicAuthID = ""
	p.BasicAuthPassword = ""
	p.AuthHeader = ""
	p.ShowOK = true
	p.ShowAlert = true
}

type DialPI struct {
	pi
	URLLeft  string `json:"url_left"`
	URLRight string `json:"url_right"`
	URLPush  string `json:"url_push"`
	URLTouch string `json:"url_touch"`
}

// IsDefault Check if its default
func (p DialPI) IsDefault() bool {
	return reflect.ValueOf(p).IsZero()
}

// Initialize PI
func (p *DialPI) Initialize() {
	p.Method = http.MethodGet
	p.URLLeft = "https://www.elgato.com"
	p.URLRight = "https://www.elgato.com"
	p.URLPush = "https://www.elgato.com"
	p.URLTouch = "https://www.elgato.com"
	p.Body = ""
	p.BasicAuthID = ""
	p.BasicAuthPassword = ""
	p.AuthHeader = ""
	p.ShowOK = true
	p.ShowAlert = true
}
