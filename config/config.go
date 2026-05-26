package config

import (
	"sync"
)

type config struct {
	sync.RWMutex
	c map[string]string
}

var cached = config{
	c: make(map[string]string),
}

// Get returns the current configuration value of a key.
//
// If it has no specific value, it falls back to a value of
// environment variable starting with "FIREWORQ_" and then a default
// value which is returned from GetDefault.
func Get(key string) string { _ = "STUB: not implemented"; return "" }

// GetDefault returns the default configuration value of a key.
func GetDefault(key string) string { _ = "STUB: not implemented"; return "" }

// Set sets the current configuration value of a key.
func Set(k, v string) { _ = "STUB: not implemented"; return }

// SetDefault sets the default configuration value of a key.
func SetDefault(k, v string) { _ = "STUB: not implemented"; return }

// Locally overrides the current configuration value of a key in a block.
//
// This is not goroutine safe and should only be used in tests.
func Locally(k, v string, block func()) { _ = "STUB: not implemented"; return }

// Keys returns a list of configuration keys.
func Keys() []string { _ = "STUB: not implemented"; return nil }
