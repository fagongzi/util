package format

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseStringBool(t *testing.T) {
	got, err := ParseStringBool("true")
	assert.NoError(t, err)
	assert.True(t, got)

	got, err = ParseStringBool("True")
	assert.NoError(t, err)
	assert.True(t, got)

	got, err = ParseStringBool("True1")
	assert.NoError(t, err)
	assert.False(t, got)
}

func TestParseStringInt64(t *testing.T) {
	value := "10"
	got, err := ParseStringInt64(value)
	assert.NoError(t, err)
	assert.Equal(t, int64(10), got)
}

func TestParseStringFloat64(t *testing.T) {
	value := "10.10"
	got, err := ParseStringFloat64(value)
	assert.NoError(t, err)
	assert.Equal(t, float64(10.10), got)
}

func TestParseStringInt(t *testing.T) {
	value := "10"
	got, err := ParseStringInt(value)
	assert.NoError(t, err)
	assert.Equal(t, int(10), got)
}

func TestParseStringIntSlice(t *testing.T) {
	value := []string{"10", "11", "12"}
	expects := []int{10, 11, 12}
	got, err := ParseStringIntSlice(value)
	assert.NoError(t, err)
	assert.Equal(t, len(value), len(got))
	for idx, expect := range expects {
		assert.Equal(t, expect, got[idx])
	}
}

func TestParseStringInt64Slice(t *testing.T) {
	value := []string{"10", "11", "12"}
	expects := []int64{10, 11, 12}
	got, err := ParseStringInt64Slice(value)
	assert.NoError(t, err)
	assert.Equal(t, len(value), len(got))
	for idx, expect := range expects {
		assert.Equal(t, expect, got[idx])
	}
}

func TestParseStringUInt64Slice(t *testing.T) {
	value := []string{"10", "11", "12"}
	expects := []uint64{10, 11, 12}
	got, err := ParseStringUint64Slice(value)
	assert.NoError(t, err)
	assert.Equal(t, len(value), len(got))
	for idx, expect := range expects {
		assert.Equal(t, expect, got[idx])
	}
}
