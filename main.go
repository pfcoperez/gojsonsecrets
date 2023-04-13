package main

import (
	"encoding/json"
	"fmt"
)

var JsonRedactSecrets bool = true

type Secret[T any] struct {
	HiddenValue   T
	redactedValue T
}

func AsSecretString(s string) Secret[string] {
	return Secret[string]{
		HiddenValue: s,
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

	sample := SampleStruct{
		Name:    "pablo",
		Age:     Secret[int]{HiddenValue: 36},
		Address: AsSecretString("earth"),
	}

	mv, _ := json.Marshal(sample)
	fmt.Println(string(mv))
}
