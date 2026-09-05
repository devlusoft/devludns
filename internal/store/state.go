// Package store provides state management for devludns.
// State is persisted to SQLite; filled in #3.
package store

// State holds the in-memory state of the DNS server.
// Filled in issue #3.
type State struct {
	// TODO(#3): add zone and record fields
}

// NewState returns an empty State placeholder.
func NewState() *State {
	return &State{}
}
