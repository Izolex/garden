package app

import (
	"fmt"
	"os"
	"strconv"
)

func MustEnv(name string) string {
	val, ok := os.LookupEnv(name)
	if !ok {
		panic(fmt.Errorf("env \"%s\" not found", name))
	}
	return val
}

func MustEnvInt(name string) int {
	val, ok := os.LookupEnv(name)
	if !ok {
		panic(fmt.Errorf("env \"%s\" not found", name))
	}
	valInt, err := strconv.Atoi(val)
	if err != nil {
		panic(err)
	}
	return valInt
}
