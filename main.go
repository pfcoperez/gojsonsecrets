package main

import (
	"encoding/json"
	"fmt"
)

type Secret[T any] struct {
	HiddenValue T
}

type SecretString Secret[string]

func AsSecretString(s string) SecretString {
	return SecretString{
		HiddenValue: s,
	}
}

func (s SecretString) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.HiddenValue)
}

type SampleStruct struct {
	Name    string
	Age     int
	Address SecretString
}

func main() {

	sample := SampleStruct{
		Name:    "pablo",
		Age:     36,
		Address: AsSecretString("earth"),
	}

	mv, _ := json.Marshal(sample)
	fmt.Println(string(mv))
}
