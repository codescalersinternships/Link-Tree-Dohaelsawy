package utils

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncryptPassword(t *testing.T) {
	testcases := []struct {
		input string
	}{
		{
			input: "password",
		},
		{
			input: "12341234",
		},
	}

	for _, test := range testcases {
		t.Run("test: success Encrypt Password", func(t *testing.T) {

			output, err := EncryptPassword(test.input)

			assert.Len(t, output, 60)
			assert.NoError(t, err)
		})
	}
}

func TestComparePassword(t *testing.T) {
	testcases := []struct {
		userPass string
		reqPass  string
		expected bool
	}{
		{
			userPass: "$2a$14$hIKTHGAMqBumm2StP9JKGOgvznTapiZZzGytyta1r1sGWWhenGWuC",
			reqPass:  "password",
			expected: true,
		},
		{
			userPass: "$2a$14$ux4I8TPeMUX2vfrZiGShMu7qM8hC41XogJ2KfJhfK17wtasjdjJTa",
			reqPass:  "43214321",
			expected: false,
		},
	}

	for _, test := range testcases {
		t.Run("test: success compare password", func(t *testing.T) {

			ok := ComparePassword(test.reqPass, test.userPass)

			assert.Equal(t, test.expected, ok)
		})
	}
}

func TestCreateToken(t *testing.T) {
	testcases := []struct {
		uid         uint
		secretToken string
		err         error
	}{
		{
			uid:         4,
			secretToken: "secret",
			err:         nil,
		},
		{
			uid:         10,
			secretToken: "-",
			err:         nil,
		},
	}

	for _, test := range testcases {
		t.Run("test: success compare password", func(t *testing.T) {

			token, err := CreateToken(test.uid, 8, test.secretToken)
			tokenSections := strings.Split(token, ".")

			assert.Equal(t, 3, len(tokenSections))
			assert.Equal(t, test.err, err)
		})
	}
}
