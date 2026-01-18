package pkg

import (
	"ohp/internal/pkg/config"
	"ohp/internal/pkg/token"
	"time"

	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(
		func(env config.Env) *token.TokenProvider {
			// 직접 아규먼트 값 주입
			// TODO: env에서 넣어주도록 개선
			return token.NewTokenProvider(
				env.JWTSecret, // secret
				"ohp-api",     // issuer

				2*time.Hour,     // accessExpiry 2hour
				23*24*time.Hour, // refreshExpiry 23day
			)
		},
	),
)
