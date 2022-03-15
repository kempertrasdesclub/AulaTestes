package jwtverify

import (
	"github.com/gbrlsnchs/jwt/v3"
)

func (e *JwtVerify) Verify(token []byte) (tokenUID, userUI string, err error) {
	var pl = CustomPayload{}
	_, err = jwt.Verify(token, e.algorithmHmacSha256, &pl)
	if err != nil {
		return
	}

	tokenUID = pl.JWTID
	userUI = pl.Issuer
	return
}
