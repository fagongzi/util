package format

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUint16ToBytes(t *testing.T) {
	value := uint16(1)
	data := Uint16ToBytes(value)
	expect := []byte{0, 1}
	assert.Equal(t, expect, data)
}

func TestUint32ToBytes(t *testing.T) {
	value := uint32(256)
	data := Uint32ToBytes(value)
	expect := []byte{0, 0, 1, 0}
	assert.Equal(t, expect, data)
}

func TestUint64ToBytes(t *testing.T) {
	value := uint64(256)
	data := Uint64ToBytes(value)
	expect := []byte{0, 0, 0, 0, 0, 0, 1, 0}
	assert.Equal(t, expect, data)
}

func TestBytesToUint64(t *testing.T) {
	expect := uint64(256)
	target := []byte{0, 0, 0, 0, 0, 0, 1, 0}
	value, err := BytesToUint64(target)
	assert.NoError(t, err)
	assert.Equal(t, expect, value)
}

func TestBytesToUint32(t *testing.T) {
	expect := uint32(256)
	target := []byte{0, 0, 1, 0}
	value, err := BytesToUint32(target)
	assert.NoError(t, err)
	assert.Equal(t, expect, value)
	if value != uint32(expect) {
		t.Errorf("failed, got=<%+v>, expect=<%+v>", value, expect)
	}
}

func TestBytesToUint16(t *testing.T) {
	expect := uint16(256)
	target := []byte{1, 0}
	value, err := BytesToUint16(target)
	assert.NoError(t, err)
	assert.Equal(t, expect, value)
}
