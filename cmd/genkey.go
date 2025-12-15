package cmd

import (
	"crypto/ecdh"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"encoding/base64"
)

type VapidKey struct {
	PublicKey  string
	PrivateKey string
}

func GenKey() (vKey VapidKey, err error) {
	// ECDSA P-256 키 생성
	pk, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return vKey, err
	}

	// 기존 ECDSA 개인키 D를 이용해 ECDH PrivateKey 생성
	curve := ecdh.P256()
	privECDH, err := curve.NewPrivateKey(pk.D.Bytes())
	if err != nil {
		return vKey, err
	}

	// Public Key 추출 (uncompressed format)
	pubKeyBytes := privECDH.PublicKey().Bytes()

	// Private Key 추출
	privKeyBytes := privECDH.Bytes()

	// Base64 URL-safe encoding (padding 없음)
	vKey.PublicKey = base64.RawURLEncoding.EncodeToString(pubKeyBytes)
	vKey.PrivateKey = base64.RawURLEncoding.EncodeToString(privKeyBytes)

	return vKey, nil
}
