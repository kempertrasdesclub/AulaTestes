package jwtverify

import (
	"github.com/gbrlsnchs/jwt/v3"
)

func (e *JwtVerify) NewAlgorithm(secretKey []byte) (err error) {
	e.algorithmHmacSha256 = jwt.NewHS256(secretKey)
	return
}
