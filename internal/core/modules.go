package core

import (
	"ohp/internal/api"
	"ohp/internal/domain/auth"
	"ohp/internal/domain/push"
	"ohp/internal/domain/user"
	"ohp/internal/infrastructure/db"

	"go.uber.org/fx"
)

var Modules = fx.Options(
	api.Module,

	// domain
	push.Module,
	auth.Module,
	user.Module,

	//infrastructure
	db.Module,
)
