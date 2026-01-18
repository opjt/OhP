package config

import (
	"errors"
	"strings"
)

type Stage int

const (
	StageUnknown Stage = iota
	StageDev
	StageProd
)

var (
	errUnknownStage = errors.New("unknown stage")
)

// 문자열로 변환
func (s Stage) String() string {
	switch s {
	case StageDev:
		return "dev"
	case StageProd:
		return "prod"
	default:
		return "unknown"
	}
}
func IsDev(s Stage) bool { return s == StageDev }

// 문자열을 Stage로 파싱
func parseStage(s string) Stage {
	switch strings.ToLower(s) {
	case "dev":
		return StageDev
	case "prod":
		return StageProd
	default:
		return StageUnknown
	}
}
