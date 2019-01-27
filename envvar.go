package envvar

import (
	"fmt"
	"os"
	"strconv"
)

func DefineString(name string, defaultValue string) {
	if _, exists := os.LookupEnv(name); !exists {
		os.Setenv(name, defaultValue)
	}
}

func DefineInt(name string, defaultValue int) {
	if _, exists := os.LookupEnv(name); !exists {
		os.Setenv(name, strconv.Itoa(defaultValue))
	}
}

func DefineBool(name string, defaultValue bool) {
	if _, exists := os.LookupEnv(name); !exists {
		os.Setenv(name, strconv.FormatBool(defaultValue))
	}
}

func GetString(name string) string {
	v := os.Getenv(name)
	if v == "" {
		panic(fmt.Sprintf("missing environment variable: %s", name))
	}
	return v
}

func GetInt(name string) int {
	s := GetString(name)
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(fmt.Sprintf("invalid value for integer environment variable %s: %s", name, s))
	}
	return i
}

func GetBool(name string) bool {
	s := GetString(name)
	b, err := strconv.ParseBool(s)
	if err != nil {
		panic(fmt.Sprintf("invalid value for boolean environment variable %s: %s", name, s))
	}
	return b
}
