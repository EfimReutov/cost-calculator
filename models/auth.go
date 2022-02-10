package models

var OTPMap = map[string]string{}

type TokenResponse struct {
	Token string `json:"token"`
}
