package config

import (
	"os"
	"strconv"
)

var (
	Host    = os.Getenv("HOST")
	Port    = EnvToInt(os.Getenv("PORT"))
	Storage = os.Getenv("STORAGE")
	Debug   = EnvToBool(os.Getenv("DEBUG"))
	Secret  = os.Getenv("SECRET")
)

func EnvToBool(env string) bool {
	b, err := strconv.ParseBool(env)
	if err != nil {
		return false
	}
	return b
}

func EnvToInt(env string) int {
	i, err := strconv.Atoi(env)
	if err != nil {
		return 0
	}
	return i
}
