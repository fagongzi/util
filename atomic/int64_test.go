package atomic

import (
	"testing"
)

func TestInt64SetAndGet(t *testing.T) {
	var v Int64
	v.Set(10)

	got := v.Get()

	if got != int64(10) {
		t.Errorf("failed, got=<%d> expect=<%d>",
			got,
			10)
	}
}
