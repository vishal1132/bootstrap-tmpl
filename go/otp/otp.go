package otp

import (
	"bytes"
	"encoding/base32"
	"math/rand"
	"time"

	"github.com/pquerna/otp"
	"github.com/pquerna/otp/totp"
)

type otpGenerator struct {
	secret     string
	randSource *rand.Rand
}

func NewOtpGenerator(secret string) *otpGenerator {

	buf := bytes.NewBuffer([]byte(""))
	encoder := base32.NewEncoder(base32.StdEncoding, buf)
	encoder.Write([]byte(secret))
	encoder.Close()
	return &otpGenerator{
		secret:     buf.String(),
		randSource: rand.New(rand.NewSource(time.Now().UnixNano() + int64(rand.Intn(1000)))),
	}
}

func (o *otpGenerator) Generate(digit int) (string, error) {

	return totp.GenerateCodeCustom(o.secret, time.Now().Add(time.Duration(o.randSource.Intn(1000))*time.Second), totp.ValidateOpts{
		Algorithm: otp.AlgorithmSHA256,
		Digits:    otp.Digits(digit),
	})
}
