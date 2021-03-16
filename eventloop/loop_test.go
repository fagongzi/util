package eventloop

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNormal(t *testing.T) {
	c := make(chan Event, 1)
	cb := func(evt Event) {
		t.Logf("cb %+v", evt)
		c <- evt
	}
	loop := New(16, cb)

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
	cb := func(evt Event) {
		t.Logf("cb %+v", evt)
		c <- evt
	}

	loop := New(16, cb).(*looper)
	loop.Close()

	event := <-c
	assert.Equal(t, CloseType, event.Type)
	assert.Nil(t, event.Data)
	event = <-c
	assert.Equal(t, EndLoopType, event.Type)
	assert.Nil(t, event.Data)
}
