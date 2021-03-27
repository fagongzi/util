package optional

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNone(t *testing.T) {
	assert.True(t, None().IsNone())
}

func TestSome(t *testing.T) {
	assert.True(t, Some(1).IsSome())
}

func TestUnwrap(t *testing.T) {
	assert.Equal(t, 1, Some(1).Unwrap())
	assert.Nil(t, None().Unwrap())
}

func TestUnwrapOrDefault(t *testing.T) {
	assert.Equal(t, 1, Some(1).UnwrapOrDefault(2))
	assert.Equal(t, 2, None().UnwrapOrDefault(2))
}

func TestUnwrapOrElse(t *testing.T) {
	assert.Equal(t, 1, Some(1).UnwrapOrElse(func() interface{} { return 2 }))
	assert.Equal(t, 2, None().UnwrapOrElse(func() interface{} { return 2 }))
}

func TestOr(t *testing.T) {
	assert.Equal(t, 1, Some(1).Or(Some(2)).Unwrap())
	assert.Equal(t, 2, None().Or(Some(2)).Unwrap())
}

func TestOrElse(t *testing.T) {
	assert.Equal(t, 1, Some(1).OrElse(func() Option { return Some(2) }).Unwrap())
	assert.Equal(t, 2, None().OrElse(func() Option { return Some(2) }).Unwrap())
}
