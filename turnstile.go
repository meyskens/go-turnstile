package turnstile

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type HCaptcha struct {
	Secret       string
	TurnstileURL string
}

type Response struct {
	// Success indicates if the challenge was passed
	Success bool `json:"success"`
	// ChallengeTs is the timestamp of the captcha
	ChallengeTs string `json:"challenge_ts"`
	// Hostname is the hostname of the passed captcha
	Hostname string `json:"hostname"`
	// ErrorCodes contains error codes returned by hCaptcha (optional)
	ErrorCodes []string `json:"error-codes"`
	// Action  is the customer widget identifier passed to the widget on the client side
	Action string `json:"action"`
	// CData is the customer data passed to the widget on the client side
	CData string `json:"cdata"`
}

func New(secret string) *HCaptcha {
	return &HCaptcha{
		Secret:       secret,
		TurnstileURL: "https://challenges.cloudflare.com/turnstile/v0/siteverify",
	}
}

// Verify verifies a "h-captcha-response" data field, with an optional remote IP set.
func (h *HCaptcha) Verify(response, remoteip string) (*Response, error) {
	values := url.Values{"secret": {h.Secret}, "response": {response}}
	if remoteip != "" {
		values.Set("remoteip", remoteip)
	}
	resp, err := http.PostForm(h.TurnstileURL, values)
	if err != nil {
		return nil, fmt.Errorf("HTTP error: %w", err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("HTTP read error: %w", err)
	}

	r := Response{}
	err = json.Unmarshal(body, &r)
	if err != nil {
		return nil, fmt.Errorf("JSON error: %w", err)
	}

	return &r, nil
}
