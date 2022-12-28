package sdhttp

import (
	"net/http"
	"reflect"
)

// PI PropertyInspector
type PI struct {
	Method            string `json:"method"`
	URL               string `json:"url"`
	Body              string `json:"body"`
	BasicAuthID       string `json:"basic_auth_id"`
	BasicAuthPassword string `json:"basic_auth_password"`
	AuthHeader        string `json:"auth_header"`
}

// IsDefault Check if its default
func (p PI) IsDefault() bool {
	return reflect.ValueOf(p).IsZero()
}

// Initialize PI
func (p *PI) Initialize() {
	p.Method = http.MethodGet
	p.URL = "https://www.elgato.com"
	p.Body = ""
	p.BasicAuthID = ""
	p.BasicAuthPassword = ""
	p.AuthHeader = ""
}
