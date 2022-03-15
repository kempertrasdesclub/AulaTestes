package jwtverify

import (
	"github.com/gbrlsnchs/jwt/v3"
	"time"
)

func (e *JwtVerify) BuildToken(userUID, tokenUID string, audience []string) (token []byte, err error) {
	now := time.Now()
	pl := CustomPayload{
		Payload: jwt.Payload{
			Issuer:         userUID,
			Subject:        "",
			Audience:       jwt.Audience(audience),
			ExpirationTime: jwt.NumericDate(now.Add(24 * 30 * 12 * time.Hour)), //fixme
			NotBefore:      jwt.NumericDate(now.Add(30 * time.Minute)),         //fixme
			IssuedAt:       jwt.NumericDate(now),
			JWTID:          tokenUID,
		},
	}

	token, err = jwt.Sign(pl, e.algorithmHmacSha256)
	return
}
