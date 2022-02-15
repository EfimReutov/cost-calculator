package rest

import (
	"cost-calculator/auth"
	"cost-calculator/models"
	"fmt"
	"math/rand"
	"net/http"
	"net/mail"
	"time"
)

const subject = "OTP"

func generateOTP() string {
	rand.Seed(time.Now().Unix())
	n := rand.Intn(9999)
	return fmt.Sprintf("%04d", n)
}

func (h *Handler) OTPSend(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	email := q.Get("email")

	_, err := mail.ParseAddress(email)
	if err != nil {
		response(w, http.StatusBadRequest, "invalid email")
		return
	}

	otp := generateOTP()

	if err = h.sender.WritePlainEmail([]string{email}, subject, otp); err != nil {
		response(w, http.StatusInternalServerError, "cannot send email")
		return
	}

	tc := &auth.TokenClaim{
		Authorized: false,
		Email:      email,
	}

	token, err := auth.CreateToken(tc, auth.DurationUnAuthorized)
	if err != nil {
		response(w, http.StatusInternalServerError, "cannot create token")
		return
	}

	models.OTPMap[email] = otp

	response(w, http.StatusOK, models.TokenResponse{Token: token})
}

func (h *Handler) authorization(w http.ResponseWriter, r *http.Request) {
	tokenString, err := auth.ValidateAccessToken(r)
	if err != nil {
		response(w, http.StatusUnauthorized, err.Error())
		return
	}

	tc, err := auth.ExtractTokenMetadata(tokenString)
	if err != nil {
		response(w, http.StatusInternalServerError, err.Error())
		return
	}

	otp, ok := models.OTPMap[tc.Email]
	if !ok {
		response(w, http.StatusInternalServerError, "no otp for this email")
		return
	}

	q := r.URL.Query()
	gotOTP := q.Get("otp")

	if otp != gotOTP {
		response(w, http.StatusUnauthorized, "otp is not match")
		return
	}

	tc.Authorized = true

	token, err := auth.CreateToken(tc, auth.DurationAuthorized)
	if err != nil {
		response(w, http.StatusInternalServerError, "cannot create token")
		return
	}

	response(w, http.StatusOK, models.TokenResponse{Token: token})
}
