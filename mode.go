package core

import "strings"

func IsProductionMode(mode ...string) bool {
	tMode := envMode
	if len(mode) > 0 {
		tMode = mode[0]
	}

	return strings.EqualFold(tMode, ValueEnvModeProduction)
}

func IsDevelopmentMode(mode ...string) bool {
	tMode := envMode
	if len(mode) > 0 {
		tMode = mode[0]
	}

	return strings.EqualFold(tMode, ValueEnvModeDevelopment)
}
