# inv

Runtime assertions for your invariants.

## Installation

```sh
go get -u github.com/disintegrator/inv
```

## Quick start

Add this import line to your Go source file:

```go
import "github.com/disintegrator/inv"
```

In any functions where you have some invariants that might constitute
preconditions, postconditions or otherwise, check them like so:

```go
func DoSomething(x int, species string, data []byte) error {
	if err := inv.Require(
		"do-something-inputs",
		"non-zero-operand", x != 0,
		"species-is-cat", species == "cat",
		"non-empty-data", len(data) > 0,
	); err != nil {
		return err
	}

	// ... rest of function ...
}
```

When invariants are not met, the returned error will summarize the failures:

```
/tmp/sandbox800493091/assert.go:32: assertion failed: do-something-inputs: non-zero-operand
/tmp/sandbox800493091/assert.go:32: assertion failed: do-something-inputs: species-is-cat
/tmp/sandbox800493091/assert.go:32: assertion failed: do-something-inputs: non-empty-data
```
