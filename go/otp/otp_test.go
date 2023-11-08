package otp

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOtpGenerate(t *testing.T) {
	otpGenerator := NewOtpGenerator("abcd")
	numTimes := 10
	for i := 0; i < numTimes; i++ {
		digits := rand.Intn(9) + 1
		otp, err := otpGenerator.Generate(digits)
		assert.NoError(t, err)
		if len(otp) != digits {
			t.Errorf("expected otp to be %d digits, but got %d digits", digits, len(otp))
		}
	}
}
