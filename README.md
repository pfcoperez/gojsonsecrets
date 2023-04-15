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
REDACT_SECRETS=true go run main.go | head -n 1 | jq '.'
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
REDACT_SECRETS=false go run main.go | head -n 1 | jq '.'
```

Yields

```json
{
  "Name": "Hiro Protagonist",
  "Age": 30,
  "Address": "U-Store-It unit"
}
```

Instance of `Stringer` interface is also provided making string values safe:

```bash
REDACT_SECRETS=true go run main.go | tail -n 1
```

Yields

```
{Hiro Protagonist -1 REDACTED}
```

And

```bash
REDACT_SECRETS=false go run main.go | tail -n 1
```

Yields

```
{Hiro Protagonist 30 U-Store-It unit}
```