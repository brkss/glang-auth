package token

import (
	"fmt"
	"time"

	"github.com/aead/chacha20poly1305"
	"github.com/o1egl/paseto"
)

type PasetoMaker struct {
	paseto 			*paseto.V2
	symetricKey 	string
}

func NewPasetoMaker(symetricKey string) (Maker, error) {
	if len(symetricKey) < chacha20poly1305.KeySize {
		return nil, fmt.Errorf("Invalid symteric key size must be %d\n", chacha20poly1305.KeySize)
	}
	maker := &PasetoMaker{
		paseto: paseto.NewV2(),
		symetricKey: symetricKey,
	}
	return maker, nil
}

func (paseto *PasetoMaker)CreateToken(userId string, duration time.Duration)(string, error){

	payload := NewPayload(userId, duration)
	token, err := paseto.paseto.Encrypt([]byte(paseto.symetricKey), payload, nil) 
	
	if err != nil {
		return "", err
	}

	return token, nil
}

func (paseto *PasetoMaker)VerifyToken(token string)(*Payload, error){

	var payload Payload
	err := paseto.paseto.Decrypt(token, []byte(paseto.symetricKey), &payload, nil)

	if err != nil {
		return nil, ErrInvalidToken 
	}

	valid, err := payload.Valid()
	if !valid {
		return nil, err 
	} 
	return &payload, nil 
}
