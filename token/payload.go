package token

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

var (
	ErrInvalidToken = errors.New("token is invalid")
	ErrExpiredToken = errors.New("token has expired")
)

type Payload struct {
	ID        uuid.UUID        `json:"id"`
	Username  string           `json:"username"`
	Role      string           `json:"role"`
	IssuedAt  time.Time        `json:"issued_at"`
	ExpiredAt time.Time        `json:"expired_at"`
	Issuer    string           `json:"issuer,omitempty"`
	Subject   string           `json:"subject,omitempty"`
	Audience  jwt.ClaimStrings `json:"audience,omitempty"`
}

func NewPayload(username, role string, duration time.Duration) (*Payload, error) {
	tokenID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	payload := &Payload{
		ID:        tokenID,
		Username:  username,
		Role:      role,
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(duration),
	}

	return payload, nil
}

func (payload *Payload) Valid() error {
	fmt.Println("Expired at:", payload.ExpiredAt)
	if time.Now().After(payload.ExpiredAt) {
		return ErrExpiredToken
	}
	return nil
}

func (payload *Payload) GetExpirationTime() (*jwt.NumericDate, error) {
	return &jwt.NumericDate{payload.ExpiredAt}, nil
}

func (payload *Payload) GetIssuedAt() (*jwt.NumericDate, error) {
	return &jwt.NumericDate{payload.IssuedAt}, nil
}

func (payload *Payload) GetNotBefore() (*jwt.NumericDate, error) {
	return &jwt.NumericDate{payload.IssuedAt}, nil
}

func (payload *Payload) GetIssuer() (string, error) {
	return payload.Issuer, nil
}

func (payload *Payload) GetSubject() (string, error) {
	return payload.Subject, nil
}

func (payload *Payload) GetAudience() (jwt.ClaimStrings, error) {
	return payload.Audience, nil
}
