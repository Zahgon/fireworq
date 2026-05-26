package test

import (
	"testing"
)

func runWithMySQL(block func()) error { _ = "STUB: not implemented"; return nil }

// Run runs a TestMain for a single "driver" configuration value.
func Run(m *testing.M) (int, error) { _ = "STUB: not implemented"; return 0, nil }

// RunAll runs a TestMain for all "driver" configuration values.
func RunAll(m *testing.M) { _ = "STUB: not implemented"; return }

// If returns if the configuration value of a key matches with one of
// the specified values.
func If(key string, values ...string) bool { _ = "STUB: not implemented"; return false }
