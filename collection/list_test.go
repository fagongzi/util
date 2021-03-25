package collection

import (
	"container/list"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTruncate(t *testing.T) {
	l := list.New()
	for i := 0; i < 10; i++ {
		l.PushBack(i)
	}
	l = Truncate(l, 9)
	assert.Equal(t, 9, l.Len())
	assert.Equal(t, 8, l.Back().Value)

	l = list.New()
	for i := 0; i < 10; i++ {
		l.PushBack(i)
	}
	l = Truncate(l, 10)
	assert.Equal(t, 10, l.Len())
}

func TestIterate(t *testing.T) {
	l := list.New()
	for i := 0; i < 10; i++ {
		l.PushBack(i)
	}

	c := 0
	sum := 0
	fn := func(v interface{}) bool {
		c++
		sum += v.(int)
		return true
	}

	Iterate(l, 0, fn)
	assert.Equal(t, 10, c)
	assert.Equal(t, 45, sum)

	c = 0
	sum = 0
	Iterate(l, 2, fn)
	assert.Equal(t, 8, c)
	assert.Equal(t, 44, sum)

	c = 0
	sum = 0
	Iterate(l, 10, fn)
	assert.Equal(t, 0, c)
	assert.Equal(t, 0, sum)

	c = 0
	sum = 0
	Iterate(l, 9, fn)
	assert.Equal(t, 1, c)
	assert.Equal(t, 9, sum)
}
