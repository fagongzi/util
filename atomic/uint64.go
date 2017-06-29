package atomic

import (
	"sync/atomic"
)

// Uint64 atomic uint64
type Uint64 struct {
	value uint64
}

// Set atomic set uint64
func (u *Uint64) Set(value uint64) {
	atomic.StoreUint64(&u.value, value)
}

// Get returns atomic uint64
func (u *Uint64) Get() uint64 {
	return atomic.LoadUint64(&u.value)
}
