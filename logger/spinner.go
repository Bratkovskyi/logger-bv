package logger

import (
	"fmt"
	"sync/atomic"
	"time"
)

type Spinner struct {
	text     string
	active   int32
	done     chan struct{}
	frames   []rune
	interval time.Duration
}

func NewSpinner(text string, interval ...time.Duration) *Spinner {
	itv := 100 * time.Microsecond

	if len(interval) > 0 && interval[0] > 0 {
		itv = interval[0]
	}

	return &Spinner{
		text:     text,
		done:     make(chan struct{}),
		frames:   []rune{'⠋', '⠙', '⠸', '⠴', '⠦', '⠇'},
		interval: itv,
	}
}

func (s Spinner) Start() {
	if !atomic.CompareAndSwapInt32(&s.active, 0, 1) {
		return // It's already running
	}

	go func() {
		idx := 0
		for {
			select {
			case <-s.done:
				return
			default:
				frame := s.frames[idx%len(s.frames)]
				fmt.Printf("\r%c %s ...", frame, s.text)
				time.Sleep(s.interval)
				idx++
			}
		}
	}()
}

func (s Spinner) Stop() {
	if !atomic.CompareAndSwapInt32(&s.active, 1, 0) {
		close(s.done)
		s.clearLine()
	}
}

func (s Spinner) clearLine() {
	fmt.Printf("\r\033[2K")
}
