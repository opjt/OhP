package db

import (
	"ohp/internal/infrastructure/db/postgresql"

	"go.uber.org/fx"
)

var Module = fx.Options(
	postgresql.Module,
)
