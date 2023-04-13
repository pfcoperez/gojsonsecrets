package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

var JsonRedactSecrets bool = true

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

func (s Secret[T]) MarshalJSON() ([]byte, error) {
	var valueToMarshal interface{} = s.HiddenValue
	if JsonRedactSecrets {
		valueToMarshal = s.redactedValue
	}
	return json.Marshal(valueToMarshal)
}

type SampleStruct struct {
	Name    string
	Age     Secret[int]
	Address Secret[string]
}

func main() {

	envRedact, errorParsingEnv := strconv.ParseBool(os.Getenv("REDACT_SECRETS"))
	if errorParsingEnv == nil {
		JsonRedactSecrets = envRedact
	}

	sample := SampleStruct{
		Name:    "Hiro Protagonist",
		Age:     Secret[int]{HiddenValue: 30, redactedValue: -1},
		Address: AsSecretString("U-Store-It unit"),
	}

	mv, _ := json.Marshal(sample)
	fmt.Println(string(mv))
}
