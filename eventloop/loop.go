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
	EndLoopType = math.MaxInt32

	stateRunning = int32(0)
	stateClosed  = int32(1)
	closeEvent   = Event{}
	endLoopEvent = Event{Type: EndLoopType}
)

var (
	// ErrClosed closed error
	ErrClosed = errors.New("event loop was closed")
)

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

// New returns a event loopper
func New(capacity int64, cb func(Event)) EventLooper {
	l := &looper{
		rb:   make([]Event, capacity, capacity),
		mask: capacity - 1,
		cb:   cb,
	}

	l.disruptor = dis.New(dis.WithCapacity(capacity),
		dis.WithConsumerGroup(l))
	go l.disruptor.Read()
	return l
}

type looper struct {
	state             int32
	disruptor         dis.Disruptor
	rb                []Event
	cb                func(Event)
	mask              int64
	closeEventHandled bool
	doEndLoop         bool
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

func (l *looper) Consume(lower, upper int64) {
	var idx int64
	addEndLoop := l.closeEventHandled
	for i := lower; i <= upper; i++ {
		idx = i & l.mask

		if l.rb[idx].Type == EndLoopType {
			// not last event
			if !l.doEndLoop {
				continue
			}

			l.disruptor.Close()
			l.cb(endLoopEvent)
			return
		}

		l.cb(l.rb[idx])
		if l.rb[idx].Type == CloseType {
			l.closeEventHandled = true
			addEndLoop = true
		}
	}

	if addEndLoop {
		// loop was already closed in outside, so doAdd endLoopEvent must be the last event
		if err := l.Add(endLoopEvent); err != nil {
			l.doEndLoop = true
			l.doAdd(endLoopEvent)
		}
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
