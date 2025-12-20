package core

import (
	"ohp/internal/api"
	"ohp/internal/domain/auth"
	"ohp/internal/domain/push"

	"go.uber.org/fx"
)

var Modules = fx.Options(
	api.Module,

	// domain
	push.Module,
	auth.Module,
)
