package pool

import (
	"errors"
	"time"
)

var (
	errSizeTooSmall = errors.New("[error] init: size of the pool is too small")
	errPoolClosed   = errors.New("[error] get: pool is closed")
	errTimeout      = errors.New("[error] get: timeout")
)

// Pool -
type Pool struct {
	active   int64
	max      int64
	res      chan interface{}
	signal   chan bool
	close    chan bool
	interval time.Duration
	creator  func() interface{}
}

// New -
func New(max int64, timer time.Duration, fn func() interface{}) (*Pool, error) {
	if max <= 0 {
		return nil, errSizeTooSmall
	}

	pool := &Pool{
		active:   0,
		max:      max,
		res:      make(chan interface{}, max),
		close:    make(chan bool),
		signal:   make(chan bool),
		interval: timer,
		creator:  fn,
	}

	for i := 0; i < int(pool.max/2); i++ {
		pool.res <- pool.creator()
		pool.active++
	}

	go pool.start()

	return pool, nil
}

func (p *Pool) start() {
	for {
		select {
		case <-p.close:
			close(p.signal)
			for {
				if int64(len(p.res)) == p.active {
					close(p.res)
					return
				}
			}
		case <-p.signal:
			if p.active < p.max {
				p.res <- p.creator()
				p.active++
			}
		}
	}
}

// Get -
func (p *Pool) Get() (interface{}, error) {
	select {
	case r := <-p.res:
		return r, nil
	case p.signal <- true:
	case <-p.close:
		return nil, errPoolClosed
	}

	ticker := time.NewTimer(p.interval)
	select {
	case r := <-p.res:
		ticker.Stop()
		return r, nil
	case <-ticker.C:
		return nil, errTimeout
	case <-p.close:
		return nil, errPoolClosed
	}
}

// Put -
func (p *Pool) Put(obj interface{}) {
	p.res <- obj
}

// Close -
func (p *Pool) Close() {
	close(p.close)
}
