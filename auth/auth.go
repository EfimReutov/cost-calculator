package auth

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"os"
	"strings"
	"time"
)

const (
	claimAuthorized = "authorized"
	claimEmail      = "email"
	claimExp        = "exp"

	DurationUnAuthorized = time.Minute
	DurationAuthorized   = time.Hour * 24 * 30
)

type TokenClaim struct {
	Email      string
	Authorized bool
}

func CreateToken(tc *TokenClaim, atDuration time.Duration) (string, error) {
	atClaims := jwt.MapClaims{}
	atClaims[claimAuthorized] = tc.Authorized
	atClaims[claimEmail] = tc.Email
	atClaims[claimExp] = time.Now().Add(atDuration).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)

	return at.SignedString([]byte(os.Getenv("ACCESS_TOKEN_SECRET")))
}

// ExtractToken gets Token string from http string.
func ExtractToken(r *http.Request) string {
	bearToken := r.Header.Get("Authorization")
	strSlc := strings.Split(bearToken, " ")
	if len(strSlc) == 2 {
		return strSlc[1]
	}
	return ""
}

// VerifyToken parses, validates, and return auth token.
func VerifyToken(tokenString, key string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv(key)), nil
	})
	if err != nil {
		return nil, fmt.Errorf("failed to parse token: %v", err)
	}
	return token, nil
}

// ValidateAccessToken checks the Access token.
func ValidateAccessToken(r *http.Request) (*jwt.Token, error) {
	tokenString := ExtractToken(r)
	return VerifyToken(tokenString, "ACCESS_TOKEN_SECRET")
}

// ExtractTokenMetadata converts metadata from jwt token to structure AccessDetails.
func ExtractTokenMetadata(token *jwt.Token) (*TokenClaim, error) {
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		tc := &TokenClaim{}
		authorized, ok := claims[claimAuthorized]
		if ok {
			auth, ok := authorized.(bool)
			if !ok {
				return nil, errors.New("error convert authorized to bool")
			}
			tc.Authorized = auth
		}
		email, ok := claims[claimEmail]
		if ok {
			e, ok := email.(string)
			if !ok {
				return nil, errors.New("error convert email to string")
			}
			tc.Email = e
		}
		return tc, nil
	}
	return nil, errors.New("error not valid token")
}
