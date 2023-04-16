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

func AsSecret[T any](value T, redactedValue ...T) Secret[T] {
	var redacted T
	if len(redactedValue) > 0 {
		redacted = redactedValue[0]
	}
	return Secret[T]{
		HiddenValue:   value,
		redactedValue: redacted,
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

func (s *Secret[T]) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &s.HiddenValue)
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
		Address: AsSecret("U-Store-It unit", "REDACTED"),
	}

	mv, _ := json.Marshal(sample)
	var unmarshalled Secret[SampleStruct]
	json.Unmarshal(mv, &unmarshalled)
	fmt.Println(string(mv))
	fmt.Println(unmarshalled)
}
