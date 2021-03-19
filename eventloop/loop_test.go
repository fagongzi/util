package eventloop

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testHandler struct {
	name string
	t    *testing.T
	c    chan Event
}

func newTestHandler(t *testing.T, c chan Event, name string) EventHandler {
	return &testHandler{
		name: name,
		t:    t,
		c:    c,
	}
}

func (h *testHandler) HandleEvent(evt Event) {
	h.t.Logf("%s handle %+v", h.name, evt)
	h.c <- evt
}

func TestNormal(t *testing.T) {
	c := make(chan Event, 1)
	loop := New(16, newTestHandler(t, c, "c1"))

	assert.NoError(t, loop.Add(Event{Type: 1, Data: 1}))
	event := <-c
	assert.Equal(t, 1, event.Type)
	assert.Equal(t, 1, event.Data)

	loop.Close()
	loop.(*looper).doAdd(Event{Type: 2, Data: 2})

	event = <-c
	assert.Equal(t, CloseType, event.Type)
	assert.Nil(t, event.Data)

	event = <-c
	assert.NotEqual(t, CloseType, event.Type)
	assert.Equal(t, 2, event.Type)
	assert.Equal(t, 2, event.Data)

	event = <-c
	assert.Equal(t, EndLoopType, event.Type)
	assert.Nil(t, event.Data)

	assert.Error(t, loop.Add(Event{Type: 1, Data: 3}))
}

func TestEndLoop(t *testing.T) {
	c := make(chan Event, 1)
	loop := New(16, newTestHandler(t, c, "c1")).(*looper)
	loop.Close()

	event := <-c
	assert.Equal(t, CloseType, event.Type)
	assert.Nil(t, event.Data)
	event = <-c
	assert.Equal(t, EndLoopType, event.Type)
	assert.Nil(t, event.Data)

	assert.Equal(t, int32(0), loop.closedCount)
}

func TestMultiHandler(t *testing.T) {
	var chs []chan Event
	var hs []EventHandler
	n := 10
	for i := 0; i < n; i++ {
		c := make(chan Event, 1)
		chs = append(chs, c)
		hs = append(hs, newTestHandler(t, c, fmt.Sprintf("c-%d", i)))
	}

	loop := New(16, hs...)
	assert.NoError(t, loop.Add(Event{Type: 1, Data: 1}))
	for _, c := range chs {
		event := <-c
		assert.Equal(t, 1, event.Type)
		assert.Equal(t, 1, event.Data)
	}

	loop.Close()
	loop.(*looper).doAdd(Event{Type: 2, Data: 2})

	for _, c := range chs {
		event := <-c
		assert.Equal(t, CloseType, event.Type)
		assert.Nil(t, event.Data)
	}

	for _, c := range chs {
		event := <-c
		assert.NotEqual(t, CloseType, event.Type)
		assert.Equal(t, 2, event.Type)
		assert.Equal(t, 2, event.Data)
	}

	for _, c := range chs {
		event := <-c
		assert.Equal(t, EndLoopType, event.Type)
		assert.Nil(t, event.Data)
	}

	assert.Error(t, loop.Add(Event{Type: 1, Data: 3}))
}

func TestMultiEndLoop(t *testing.T) {
	var chs []chan Event
	var hs []EventHandler
	n := 10
	for i := 0; i < n; i++ {
		c := make(chan Event, 1)
		chs = append(chs, c)
		hs = append(hs, newTestHandler(t, c, fmt.Sprintf("c-%d", i)))
	}
	loop := New(16, hs...).(*looper)
	loop.Close()

	for _, c := range chs {
		event := <-c
		assert.Equal(t, CloseType, event.Type)
		assert.Nil(t, event.Data)
	}

	for _, c := range chs {
		event := <-c
		assert.Equal(t, EndLoopType, event.Type)
		assert.Nil(t, event.Data)
	}

	assert.Equal(t, int32(0), loop.closedCount)
}
