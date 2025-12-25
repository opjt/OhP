package token

import (
	"crypto/rand"
	"io"
)

func GenerateEndpointToken(length int) (string, error) {
	return generateEndpointToken(rand.Reader, length)
}

func generateEndpointToken(r io.Reader, length int) (string, error) {
	// base62에서 문자열 하나 당 6비트
	byteLen := length * 6 / 8
	if byteLen < 8 {
		byteLen = 8
	}

	b := make([]byte, byteLen)

	if _, err := io.ReadFull(r, b); err != nil {
		return "", err
	}

	token := Base62Encode(b)
	if len(token) > length {
		token = token[:length]
	}

	for len(token) < length {
		token = "0" + token
	}
	return token, nil
}
