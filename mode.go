package core

import "strings"

//
// MARK: Full name funcs

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

//
// MARK: Shorten name funcs

func IsPM(mode ...string) bool {
	return IsProductionMode(mode...)
}
func IsDM(mode ...string) bool {
	return IsDevelopmentMode(mode...)
}
