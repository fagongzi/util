package eventloop

import (
	"errors"
	"math"
	"sync/atomic"

	dis "github.com/smartystreets-prototypes/go-disruptor"
)

var (
	// CloseType closed event type
	CloseType = 0
	// EndLoopType loop exit type
	EndLoopType    = math.MaxInt32
	preEndLoopType = math.MaxInt32 - 1

	stateRunning    = int32(0)
	stateClosed     = int32(1)
	closeEvent      = Event{}
	endLoopEvent    = Event{Type: EndLoopType}
	preEndLoopEvent = Event{Type: preEndLoopType}
)

var (
	// ErrClosed closed error
	ErrClosed = errors.New("event loop was closed")
)

// EventHandler
type EventHandler interface {
	HandleEvent(Event)
}

// Event event
type Event struct {
	Type int
	Data interface{}
}

// EventLooper event loop
type EventLooper interface {
	// Close the event loop
	Close()
	// Add add events to event loop
	Add(...Event) error
}

type consumer struct {
	h                 EventHandler
	l                 *looper
	main              bool
	closeEventHandled bool
}

func (c *consumer) Consume(lower, upper int64) {
	var idx int64
	addEndLoop := c.closeEventHandled
	for i := lower; i <= upper; i++ {
		idx = i & c.l.mask
		et := c.l.rb[idx].Type
		if et == preEndLoopType {
			continue
		} else if et == EndLoopType {
			c.l.maybeDoClose()
			c.h.HandleEvent(endLoopEvent)
			return
		}

		c.h.HandleEvent(c.l.rb[idx])
		if et == CloseType {
			c.closeEventHandled = true
			addEndLoop = true
		}
	}

	if c.main && addEndLoop {
		// loop was already closed in outside, so doAdd endLoopEvent must be the last event
		if err := c.l.Add(preEndLoopEvent); err != nil {
			c.l.doAdd(endLoopEvent)
		}
	}
}

// New returns a event loopper
func New(capacity int64, handlers ...EventHandler) EventLooper {
	l := &looper{
		rb:          make([]Event, capacity, capacity),
		mask:        capacity - 1,
		closedCount: int32(len(handlers)),
	}

	var consumers []dis.Consumer
	for i, h := range handlers {
		consumers = append(consumers, &consumer{l: l, h: h, main: i == 0})
	}

	l.disruptor = dis.New(dis.WithCapacity(capacity),
		dis.WithConsumerGroup(consumers...))
	go l.disruptor.Read()
	return l
}

type looper struct {
	state       int32
	disruptor   dis.Disruptor
	rb          []Event
	cb          func(Event)
	mask        int64
	closedCount int32
}

func (l *looper) Close() {
	current := l.getState()
	if current == stateClosed {
		return
	}

	if atomic.CompareAndSwapInt32(&l.state, current, stateClosed) {
		l.doAdd(closeEvent)
	}
}

func (l *looper) Add(events ...Event) error {
	current := l.getState()
	if current == stateClosed {
		return ErrClosed
	}

	l.doAdd(events...)
	return nil
}

func (l *looper) maybeDoClose() {
	if atomic.AddInt32(&l.closedCount, -1) == 0 {
		l.disruptor.Close()
	}
}

func (l *looper) doAdd(events ...Event) {
	n := int64(len(events))
	j := 0
	seq := l.disruptor.Reserve(n)
	for i := seq - n + 1; i <= seq; i++ {
		l.rb[i&l.mask] = events[j]
		j++
	}
	l.disruptor.Commit(seq-n+1, seq)
}

func (l *looper) getState() int32 {
	return atomic.LoadInt32(&l.state)
}
