package pool

import (
	"errors"
	"sync"
	"time"
)

var (
	errSizeTooSmall = errors.New("[error] init: size of the pool is too small")
	errPoolClosed   = errors.New("[error] get: pool is closed")
	errTimeout      = errors.New("[error] get: timeout")
)

// Pool -
type Pool struct {
	lock     sync.Mutex
	active   int64
	max      int64
	res      chan interface{}
	signal   chan bool
	isClose  bool
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
		isClose:  false,
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
			for {
				p.lock.Lock()
				if int64(len(p.res)) == p.active {
					p.isClose = true
					close(p.res)
					close(p.signal)
					p.lock.Unlock()
					return
				}
				p.lock.Unlock()
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
	p.lock.Lock()
	defer p.lock.Unlock()

	if p.isClose {
		return nil, errPoolClosed
	}

	ticker := time.NewTimer(p.interval)
	select {
	case r := <-p.res:
		ticker.Stop()
		p.signal <- true
		return r, nil
	case <-ticker.C:
		return nil, errTimeout
	}
}

// Put -
func (p *Pool) Put(obj interface{}) error {
	p.lock.Lock()
	defer p.lock.Unlock()

	p.res <- obj
	return nil
}

// Close -
func (p *Pool) Close() {
	p.close <- true
	close(p.close)
}
