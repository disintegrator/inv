package inv_test

import (
	"errors"
	"testing"

	"github.com/disintegrator/inv"
	"github.com/stretchr/testify/require"
)

func TestDebug(t *testing.T) {
	var err error
	inv.Debug("valid",
		"err function", func() error { return nil },
		"err value", err,
		"bool function", func() bool { return true },
		"bool value", true,
		"nil value", nil,
	)
}

func TestDebugPanic(t *testing.T) {
	cases := []struct {
		id    string
		value any
	}{
		{"err function", func() error { return errors.New("simulated") }},
		{"err value", errors.New("simulated")},
		{"bool function", func() bool { return false }},
		{"bool value", false},
	}

	for _, tc := range cases {
		require.Panics(t, func() { inv.Debug("expected failure", tc.id, tc.value) })
	}
}
