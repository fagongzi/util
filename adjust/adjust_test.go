package adjust

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInt(t *testing.T) {
	var value, adjust int
	value = 0
	adjust = 1

	got := Int(value, adjust)
	if got != adjust {
		t.Errorf("failed, got=<%d> expect=<%d>", got, adjust)
	}

	value = 1
	adjust = 2
	got = Int(value, adjust)
	if got != value {
		t.Errorf("failed, got=<%d> expect=<%d>", got, value)
	}
}

func TestInt8(t *testing.T) {
	var value, adjust int8
	value = 0
	adjust = 1

	got := Int8(value, adjust)
	if got != adjust {
		t.Errorf("failed, got=<%d> expect=<%d>", got, adjust)
	}

	value = 1
	adjust = 2
	got = Int8(value, adjust)
	if got != value {
		t.Errorf("failed, got=<%d> expect=<%d>", got, value)
	}
}

func TestInt16(t *testing.T) {
	var value, adjust int16
	value = 0
	adjust = 1

	got := Int16(value, adjust)
	if got != adjust {
		t.Errorf("failed, got=<%d> expect=<%d>", got, adjust)
	}

	value = 1
	adjust = 2
	got = Int16(value, adjust)
	if got != value {
		t.Errorf("failed, got=<%d> expect=<%d>", got, value)
	}
}

func TestInt32(t *testing.T) {
	var value, adjust int32
	value = 0
	adjust = 1

	got := Int32(value, adjust)
	if got != adjust {
		t.Errorf("failed, got=<%d> expect=<%d>", got, adjust)
	}

	value = 1
	adjust = 2
	got = Int32(value, adjust)
	if got != value {
		t.Errorf("failed, got=<%d> expect=<%d>", got, value)
	}
}

func TestInt64(t *testing.T) {
	var value, adjust int64
	value = 0
	adjust = 1

	got := Int64(value, adjust)
	if got != adjust {
		t.Errorf("failed, got=<%d> expect=<%d>", got, adjust)
	}

	value = 1
	adjust = 2
	got = Int64(value, adjust)
	if got != value {
		t.Errorf("failed, got=<%d> expect=<%d>", got, value)
	}
}

func TestUInt(t *testing.T) {
	var value, adjust uint
	value = 0
	adjust = 1

	got := UInt(value, adjust)
	if got != adjust {
		t.Errorf("failed, got=<%d> expect=<%d>", got, adjust)
	}

	value = 1
	adjust = 2
	got = UInt(value, adjust)
	if got != value {
		t.Errorf("failed, got=<%d> expect=<%d>", got, value)
	}
}

func TestUInt8(t *testing.T) {
	var value, adjust uint8
	value = 0
	adjust = 1

	got := UInt8(value, adjust)
	if got != adjust {
		t.Errorf("failed, got=<%d> expect=<%d>", got, adjust)
	}

	value = 1
	adjust = 2
	got = UInt8(value, adjust)
	if got != value {
		t.Errorf("failed, got=<%d> expect=<%d>", got, value)
	}
}

func TestUInt16(t *testing.T) {
	var value, adjust uint16
	value = 0
	adjust = 1

	got := UInt16(value, adjust)
	if got != adjust {
		t.Errorf("failed, got=<%d> expect=<%d>", got, adjust)
	}

	value = 1
	adjust = 2
	got = UInt16(value, adjust)
	if got != value {
		t.Errorf("failed, got=<%d> expect=<%d>", got, value)
	}
}

func TestUInt32(t *testing.T) {
	var value, adjust uint32
	value = 0
	adjust = 1

	got := UInt32(value, adjust)
	if got != adjust {
		t.Errorf("failed, got=<%d> expect=<%d>", got, adjust)
	}

	value = 1
	adjust = 2
	got = UInt32(value, adjust)
	if got != value {
		t.Errorf("failed, got=<%d> expect=<%d>", got, value)
	}
}

func TestUInt64(t *testing.T) {
	var value, adjust uint64
	value = 0
	adjust = 1

	got := UInt64(value, adjust)
	if got != adjust {
		t.Errorf("failed, got=<%d> expect=<%d>", got, adjust)
	}

	value = 1
	adjust = 2
	got = UInt64(value, adjust)
	if got != value {
		t.Errorf("failed, got=<%d> expect=<%d>", got, value)
	}
}

func TestMaxInt(t *testing.T) {
	assert.Equal(t, 2, MaxInt(1, 2))
}

func TestMaxInt8(t *testing.T) {
	assert.Equal(t, int8(2), MaxInt8(1, 2))
}

func TestMaxInt16(t *testing.T) {
	assert.Equal(t, int16(2), MaxInt16(1, 2))
}

func TestMaxInt32(t *testing.T) {
	assert.Equal(t, int32(2), MaxInt32(1, 2))
}

func TestMaxInt64(t *testing.T) {
	assert.Equal(t, int64(2), MaxInt64(1, 2))
}

func TestMaxUInt(t *testing.T) {
	assert.Equal(t, uint(2), MaxUInt(1, 2))
}

func TestMaxUInt8(t *testing.T) {
	assert.Equal(t, uint8(2), MaxUInt8(1, 2))
}

func TestMaxUInt16(t *testing.T) {
	assert.Equal(t, uint16(2), MaxUInt16(1, 2))
}

func TestMaxUInt32(t *testing.T) {
	assert.Equal(t, uint32(2), MaxUInt32(1, 2))
}

func TestMaxUInt64(t *testing.T) {
	assert.Equal(t, uint64(2), MaxUInt64(1, 2))
}

func TestMinInt(t *testing.T) {
	assert.Equal(t, 1, MinInt(1, 2))
}

func TestMinInt8(t *testing.T) {
	assert.Equal(t, int8(1), MinInt8(1, 2))
}

func TestMinInt16(t *testing.T) {
	assert.Equal(t, int16(1), MinInt16(1, 2))
}

func TestMinInt32(t *testing.T) {
	assert.Equal(t, int32(1), MinInt32(1, 2))
}

func TestMinInt64(t *testing.T) {
	assert.Equal(t, int64(1), MinInt64(1, 2))
}

func TestMinUInt(t *testing.T) {
	assert.Equal(t, uint(1), MinUInt(1, 2))
}

func TestMinUInt8(t *testing.T) {
	assert.Equal(t, uint8(1), MinUInt8(1, 2))
}

func TestMinUInt16(t *testing.T) {
	assert.Equal(t, uint16(1), MinUInt16(1, 2))
}

func TestMinUInt32(t *testing.T) {
	assert.Equal(t, uint32(1), MinUInt32(1, 2))
}

func TestMinUInt64(t *testing.T) {
	assert.Equal(t, uint64(1), MinUInt64(1, 2))
}
