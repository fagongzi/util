package atomic

import (
	"sync/atomic"
)

// Uint32 atomic uint64
type Uint32 struct {
	value uint32
}

// Set atomic set uint32
func (u *Uint32) Set(value uint32) {
	atomic.StoreUint32(&u.value, value)
}

// Get returns atomic uint32
func (u *Uint32) Get() uint32 {
	return atomic.LoadUint32(&u.value)
}

// Incr incr atomic uint32
func (u *Uint32) Incr() uint32 {
	return atomic.AddUint32(&u.value, 1)
}
