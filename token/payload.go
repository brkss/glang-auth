package token

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

var ErrExpiredToken = errors.New("this token has been expired")
var ErrInvalidToken = errors.New("invalid token !")

type Payload struct {
	ID string 
	UserId string
	ExpiredAt time.Time
	IssuedAt time.Time
}

func NewPayload(userId string, duration time.Duration) (*Payload){
	return &Payload{
		ID: uuid.New().String(),
		UserId: userId,
		ExpiredAt: time.Now().Add(duration),
		IssuedAt: time.Now(),
	}
}

func (p *Payload)Valid()(bool, error){

	if time.Now().After(p.ExpiredAt){
		return false, ErrExpiredToken
	}

	return true, nil
}

