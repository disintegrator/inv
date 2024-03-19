# inv

Runtime assertions for your invariants.

## Why?

This package is inspired by the prior art around runtime assertions in other
languages. It tries to bridge the gap between the Go philosophy of proper
handling by not encouraging the use of panics around assertions and by capturing
useful, contextual information about invariants at key points of a program.

Some references:

- https://github.com/tigerbeetle/tigerbeetle/blob/main/docs/TIGER_STYLE.md#safety
- https://doc.rust-lang.org/std/macro.assert.html
- https://ziglang.org/documentation/master/std/#A;std:debug.assert
- https://www.npmjs.com/package/ts-invariant

## Installation

```sh
go get -u github.com/disintegrator/inv
```

## Quick start

Add this import line to your Go source files:

```go
import "github.com/disintegrator/inv"
```

In any functions where you have some invariants that might constitute
preconditions, postconditions or otherwise, check them like so:

```go
func DoSomething(x int, species string, data []byte) error {
	if err := inv.Check(
		"do-something-inputs",
		"non-zero-operand", x != 0,
		"species-is-cat", species == "cat",
		"non-empty-data", len(data) > 0,
	); err != nil {
		return err
	}

	return nil
}

func main() {
	fmt.Println(DoSomething(0, "nope.txt", "dog", nil))
}
```

When invariants are not met, the returned error will summarize the failures:

```
/tmp/sandbox/prog.go:12: invariant mismatch: do something inputs: non-zero operand
/tmp/sandbox/prog.go:12: invariant mismatch: do something inputs: species is cat
/tmp/sandbox/prog.go:12: invariant mismatch: do something inputs: non-empty data
/tmp/sandbox/prog.go:12: invariant mismatch: do something inputs: path can be opened: open nope.txt: no such file or directory
```

## Panics

There are two variations of the `Check` function that trigger panics based on `go build` tags.

`inv.Require` will always panic if any of the invariants are not met in a group:

```go
func DoSomething(x int, species string, data []byte) error {
	inv.Require(
		"do-something-inputs",
		"non-zero-operand", x != 0,
		"species-is-cat", species == "cat",
		"non-empty-data", len(data) > 0,
	)

	// ... rest of function ...
}
```

This is the equivalent of runtime assertions in other languages.

`inv.Debug` is a variation that has no effect in regular builds but it can be
enabled by setting the `inv.debug` build tag:

```
go build -tags inv.debug [build flags]
```

```go
func DoSomething(x int, species string, data []byte) error {
	inv.Debug(
		"do-something-inputs",
		"non-zero-operand", x != 0,
		"species-is-cat", species == "cat",
		"non-empty-data", len(data) > 0,
	)

	// ... rest of function ...
}
```

This is the equivalent of debug assertions in other languages.
