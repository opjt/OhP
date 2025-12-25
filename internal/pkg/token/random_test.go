package token

import (
	"bytes"
	"regexp"
	"testing"
)

// basic test
func TestGenerateEndpointToken_Basic(t *testing.T) {
	lengths := []int{5, 8, 16, 32, 64}
	base62Regex := regexp.MustCompile(`^[0-9a-zA-Z]+$`)

	for _, l := range lengths {
		// 실제 rand.Reader를 사용하는 래퍼 함수 호출
		token, err := GenerateEndpointToken(l)
		if err != nil {
			t.Errorf("length %d, err: %v", l, err)
		}

		// 결과물 길이 검증
		if len(token) != l {
			t.Errorf("value mismatch, expected: %d, got: %d", l, len(token))
		}

		// 문자 구성 검증 (Base62 문자만 들어있는지)
		if !base62Regex.MatchString(token) {
			t.Errorf("invalid string included: %s", token)
		}
	}
}

// padding logic test
func TestGenerateEndpointToken_PaddingLogic(t *testing.T) {
	length := 10
	// 모든 바이트가 0인 데이터를 제공하는 가짜 리더
	// 0으로 채워진 바이트는 Base62 인코딩 시 ''(빈문자열) 값이 나옴.
	mockZeroReader := bytes.NewReader(make([]byte, 20))

	token, err := generateEndpointToken(mockZeroReader, length)
	if err != nil {
		t.Fatalf("test error: %v", err)
	}

	if len(token) != length {
		t.Errorf("Length mismatch after padding, expected: %d, got: %d", length, len(token))
	}

	expected := "0000000000"
	if token != expected {
		t.Errorf("result mismatch, expected: %s, got: %s", expected, token)
	}
}
