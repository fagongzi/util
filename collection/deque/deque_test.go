package deque

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPushBack(t *testing.T) {
	q := New()
	q.PushBack(1)
	assert.Equal(t, 1, q.Len())
}

func TestTruncate(t *testing.T) {
	q := New()
	for i := 0; i < 10; i++ {
		q.PushBack(i)
	}

	q.Truncate(q.Len())
	assert.Equal(t, 10, q.Len())

	q.Truncate(q.Len() + 1)
	assert.Equal(t, 10, q.Len())

	q.Truncate(1)
	assert.Equal(t, 1, q.Len())
	v, ok := q.Front()
	assert.True(t, ok)
	i := 0
	for e := v; e != nil; e = e.Next() {
		assert.Equal(t, i, e.Value)
		i++
	}
	v, ok = q.Back()
	for e := v; e != nil; e = e.Prev() {
		i--
		assert.Equal(t, i, e.Value)
	}

	q.Truncate(0)
	assert.Equal(t, 0, q.Len())
	_, ok = q.Front()
	assert.False(t, ok)
	_, ok = q.Back()
	assert.False(t, ok)
}

func TestDrain(t *testing.T) {
	// case 1: drain 5,10 from 0-9
	q := New()
	for i := 0; i < 10; i++ {
		q.PushBack(i)
	}
	drained := q.Drain(5, 10)

	assert.Equal(t, 5, q.Len())
	v, ok := q.Front()
	assert.True(t, ok)
	i := 0
	for e := v; e != nil; e = e.Next() {
		assert.Equal(t, i, e.Value)
		i++
	}
	v, ok = q.Back()
	for e := v; e != nil; e = e.Prev() {
		i--
		assert.Equal(t, i, e.Value)
	}

	v, ok = drained.Front()
	assert.Equal(t, 5, drained.Len())
	assert.True(t, ok)
	i = 5
	for e := v; e != nil; e = e.Next() {
		assert.Equal(t, i, e.Value)
		i++
	}
	v, ok = drained.Back()
	for e := v; e != nil; e = e.Prev() {
		i--
		assert.Equal(t, i, e.Value)
	}

	// case 2: drain 5,9 from 0-9
	q.Clear()
	for i := 0; i < 10; i++ {
		q.PushBack(i)
	}
	drained = q.Drain(5, 9)

	assert.Equal(t, 6, q.Len())

	i = 0
	values := []int{0, 1, 2, 3, 4, 9}
	v, ok = q.Front()
	assert.True(t, ok)
	for e := v; e != nil; e = e.Next() {
		assert.Equal(t, values[i], e.Value)
		i++
	}
	v, ok = q.Back()
	for e := v; e != nil; e = e.Prev() {
		i--
		assert.Equal(t, values[i], e.Value)
	}

	v, ok = drained.Front()
	assert.Equal(t, 4, drained.Len())
	assert.True(t, ok)
	i = 5
	for e := v; e != nil; e = e.Next() {
		assert.Equal(t, i, e.Value)
		i++
	}
	v, ok = drained.Back()
	for e := v; e != nil; e = e.Prev() {
		i--
		assert.Equal(t, i, e.Value)
	}

	// case 3: drain 0,5 from 0-9
	q.Clear()
	for i := 0; i < 10; i++ {
		q.PushBack(i)
	}
	drained = q.Drain(0, 5)

	assert.Equal(t, 5, q.Len())
	i = 5
	v, ok = q.Front()
	assert.True(t, ok)
	for e := v; e != nil; e = e.Next() {
		assert.Equal(t, i, e.Value)
		i++
	}
	v, ok = q.Back()
	for e := v; e != nil; e = e.Prev() {
		i--
		assert.Equal(t, i, e.Value)
	}

	v, ok = drained.Front()
	assert.Equal(t, 5, drained.Len())
	assert.True(t, ok)
	i = 0
	for e := v; e != nil; e = e.Next() {
		assert.Equal(t, i, e.Value)
		i++
	}
	v, ok = drained.Back()
	for e := v; e != nil; e = e.Prev() {
		i--
		assert.Equal(t, i, e.Value)
	}

	// case 3: drain 0,9 from 0-9
	q.Clear()
	for i := 0; i < 10; i++ {
		q.PushBack(i)
	}
	drained = q.Drain(0, 10)
	assert.Equal(t, 0, q.Len())
	_, ok = q.Front()
	assert.False(t, ok)
	_, ok = q.Back()
	assert.False(t, ok)

	assert.Equal(t, 10, drained.Len())
	v, ok = drained.Front()
	assert.True(t, ok)
	i = 0
	for e := v; e != nil; e = e.Next() {
		assert.Equal(t, i, e.Value)
		i++
	}
	v, ok = drained.Back()
	for e := v; e != nil; e = e.Prev() {
		i--
		assert.Equal(t, i, e.Value)
	}
}
