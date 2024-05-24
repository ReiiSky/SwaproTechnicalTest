package environments

import (
	"fmt"
	"strconv"
)

type Env struct {
	kv map[string]string
}

func (env Env) Get(k string) string {
	value, ok := env.kv[k]

	if !ok {
		panic(fmt.Errorf("config '%s' not found", k))
	}

	return value
}

func (env Env) GetString(k string) string {
	return env.Get(k)
}

func (env Env) GetInt(k string) int {
	valueStr := env.Get(k)

	valueInt, err := strconv.Atoi(valueStr)

	if err != nil {
		panic(fmt.Errorf("error parse conf key: %s", err))
	}

	return valueInt
}

type EnvVarContainer interface {
	GetString(string) string
	GetInt(string) int
}
