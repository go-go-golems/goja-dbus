package dbuscore

// Policy is intentionally small in Phase 2. Later phases will expand this into
// the Go-side authority for bus connection, call, signal, and service export
// permissions.
type Policy struct {
	AllowSessionBus bool
	AllowSystemBus  bool
}

func DefaultPolicy() Policy {
	return Policy{AllowSessionBus: true}
}
