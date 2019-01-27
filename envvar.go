// Package envvar provides functions to access typed
// environment variables and define default values.
package envvar

import (
	"fmt"
	"os"
	"strconv"
)

// DefineString defines a string environment variable
// with a defaultValue.
func DefineString(name string, defaultValue string) {
	if _, exists := os.LookupEnv(name); !exists {
		os.Setenv(name, defaultValue)
	}
}

// DefineInt defines an int environment variable
// with a defaultValue.
func DefineInt(name string, defaultValue int) {
	if _, exists := os.LookupEnv(name); !exists {
		os.Setenv(name, strconv.Itoa(defaultValue))
	}
}

// DefineBool defines a boolean environment variable
// with a defaultValue.
func DefineBool(name string, defaultValue bool) {
	if _, exists := os.LookupEnv(name); !exists {
		os.Setenv(name, strconv.FormatBool(defaultValue))
	}
}

// GetString returns the value of the environment
// variable with the given name. If the environment
// variable is not set and no defaultValue was defined,
// this func panics.
func GetString(name string) string {
	v := os.Getenv(name)
	if v == "" {
		panic(fmt.Sprintf("missing environment variable: %s", name))
	}
	return v
}

// GetInt returns the value of the environment
// variable with the given name converted to int.
// If the environment variable is not set and no
// defaultValue was defined, this func panics.
func GetInt(name string) int {
	s := GetString(name)
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(fmt.Sprintf("invalid value for integer environment variable %s: %s", name, s))
	}
	return i
}

// GetBool returns the value of the environment
// variable with the given name converted to bool.
// If the environment variable is not set and no
// defaultValue was defined, this func panics.
func GetBool(name string) bool {
	s := GetString(name)
	b, err := strconv.ParseBool(s)
	if err != nil {
		panic(fmt.Sprintf("invalid value for boolean environment variable %s: %s", name, s))
	}
	return b
}
