package token

import "time"

type Maker interface {
	CreateToken(id string, duration time.Time)
	VerifyToken(token string)
}
