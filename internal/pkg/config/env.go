package config

import (
	"context"
	"os"

	"github.com/joho/godotenv"
	"github.com/sethvargo/go-envconfig"
)

type Env struct {
	stageRaw string `ignored:"true"`
	Stage    Stage  `ignored:"true"`

	Log     Log     `env:", prefix=LOG_"`
	Service Service `env:", prefix=SERVICE_"`
}

type SwitApi struct {
	Url string `env:"URL"`
}

type Log struct {
	Level string `env:"LEVEL"`
}

type Service struct {
	Port int `env:"PORT,default=8282"`
}

func NewEnv() (Env, error) {
	var env Env

	envPath := os.Getenv("CONFPATH")
	if envPath != "" {
		_ = godotenv.Load(envPath)
	}

	if err := envconfig.Process(context.Background(), &env); err != nil {
		return env, err
	}
	env.stageRaw = os.Getenv("STAGE")
	env.Stage = parseStage(env.stageRaw)

	if err := validateEnv(&env); err != nil {
		return env, err
	}

	return env, nil
}

func validateEnv(e *Env) error {
	if e.Stage == StageUnknown {
		return errUnknownStage
	}

	return nil

}
