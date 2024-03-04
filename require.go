package inv

import (
	"errors"
	"fmt"
	"runtime"
	"strings"
)

const dbgPrefix = "_ "

// Require checks a set of invariants and returns an error if any of them fail.
// Each invariant is passed as a pair: the name of the invariant as a string and
// a value that may be a boolean or error. If any value is false or a non-nil
// error across the pairs, then an error is returned with all the unmet
// invariants. Since this function checks multiple invariants, a group name is
// used to relate these checks together.
//
// If the invariant name starts with "_ " then it is only considered if debug
// mode is enabled (it is by default). Other invariants are always checked.
//
//go:noinline
func Require(group string, pairs ...any) error {
	if len(pairs)%2 != 0 {
		return errors.New("invariants must be passed as pairs")
	}

	var errs error
	_, file, no, ok := runtime.Caller(1)
	caller := "<unknown>:0"
	if ok {
		caller = fmt.Sprintf("%s:%d", file, no)
	}

	for i := 0; i < len(pairs); i += 2 {
		id, pred := pairs[i].(string), pairs[i+1]

		isdebug := strings.HasPrefix(id, dbgPrefix)
		if isdebug && !debugMode {
			continue
		}

		invid := id
		if isdebug {
			invid = id[len(dbgPrefix):]
		}

		switch val := pred.(type) {
		case error:
			if val != nil {
				errs = errors.Join(errs, &InvariantError{
					Group: group, Invariant: invid,
					Caller: caller, Debug: isdebug, Cause: val,
				})
			}
		case bool:
			if !val {
				errs = errors.Join(errs, &InvariantError{
					Group: group, Invariant: invid,
					Debug: isdebug, Caller: caller,
				})
			}
		}
	}

	return errs
}

// Must is similar to Require but panics if any invariant checks fail.
func Must(group string, pairs ...any) {
	if err := Require(group, pairs...); err != nil {
		panic(err)
	}
}

// InvariantError represents an invariant mismatch. If the invariant failed
// because of a non-nil error then Cause will be set.
type InvariantError struct {
	// Invariant is the name of the invariant that was not met.
	Invariant string
	// Group is a label that relates this error to others.
	Group string
	// The location in code where the invariant was checked.
	Caller string
	// Cause is the underlying error that caused the invariant check to fail.
	Cause error
	// Debug is true if the invariant was only checked in debug mode.
	Debug bool
}

// Error returns a string representation of the invariant error.
func (ae *InvariantError) Error() string {
	title := "invariant mismatch"
	if ae.Debug {
		title = "(debug) invariant mismatch"
	}

	msg := fmt.Sprintf("%s: %s: %s: %s", ae.Caller, title, ae.Group, ae.Invariant)
	if ae.Cause != nil {
		msg += ": " + ae.Cause.Error()
	}

	return msg
}

// Unwrap returns the underlying error that caused the invariant check to fail.
func (ae *InvariantError) Unwrap() error {
	return ae.Cause
}

func DoSomething(x int, species string, data []byte) error {
	if err := Require(
		"do-something-inputs",
		"non-zero-operand", x != 0,
		"species-is-cat", species == "cat",
		"non-empty-data", len(data) > 0,
	); err != nil {
		return err
	}

	return nil
}
