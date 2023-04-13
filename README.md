# gojsonsecrets

Sample project, and portantial seed for a library, implementing sensitive struct fields tagging and showing its effects on JSON Marshallers.

For a given harcoded example of the following type:

```go
type SampleStruct struct {
	Name    string
	Age     Secret[int]
	Address Secret[string]
}
```

With value:

```go
sample := SampleStruct{
		Name:    "Hiro Protagonist",
		Age:     Secret[int]{HiddenValue: 30, redactedValue: -1},
		Address: AsSecretString("U-Store-It unit"),
	}
```

We can get two run examples:

```bash
REDACT_SECRETS=true go run main.go | jq 
```

Yields

```json
{
  "Name": "Hiro Protagonist",
  "Age": -1,
  "Address": "REDACTED"
}
```

And

```bash
REDACT_SECRETS=true go run main.go
```

Yields

```json
{
  "Name": "Hiro Protagonist",
  "Age": 30,
  "Address": "U-Store-It unit"
}
```
