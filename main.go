package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

var RedactSecrets bool = true

type Secret[T any] struct {
	HiddenValue   T
	redactedValue T
}

func AsSecretString(s string) Secret[string] {
	return Secret[string]{
		HiddenValue:   s,
		redactedValue: "REDACTED",
	}
}

// Common

func safeValue[T any](s Secret[T]) T {
	if RedactSecrets {
		return s.redactedValue
	} else {
		return s.HiddenValue
	}
}

// JSON

func (s Secret[T]) MarshalJSON() ([]byte, error) {
	return json.Marshal(safeValue(s))
}

// Strings

func (s Secret[T]) String() string {
	return fmt.Sprint(safeValue(s))
}

type SampleStruct struct {
	Name    string
	Age     Secret[int]
	Address Secret[string]
}

func main() {

	envRedact, errorParsingEnv := strconv.ParseBool(os.Getenv("REDACT_SECRETS"))
	if errorParsingEnv == nil {
		RedactSecrets = envRedact
	}

	sample := SampleStruct{
		Name:    "Hiro Protagonist",
		Age:     Secret[int]{HiddenValue: 30, redactedValue: -1},
		Address: AsSecretString("U-Store-It unit"),
	}

	mv, _ := json.Marshal(sample)
	fmt.Println(string(mv))
	fmt.Println(sample)
}
