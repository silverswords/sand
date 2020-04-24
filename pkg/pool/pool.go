package pool

import (
	"errors"
	"sync"
	"time"
)

var (
	errSizeTooSmall = errors.New("[error] size of the pool is too small")
	errPoolClosed   = errors.New("[error] pool is closed")
	errPoolEmpty    = errors.New("[error] pool is empty")
	errTimeout      = errors.New("[error] timeout")

	timerCh      = make(chan bool)
	isRequsetEnd = make(chan bool)
)

// Pool -
type Pool struct {
	lock    sync.Mutex
	active  int64
	max     int64
	res     chan interface{}
	isClose bool
	close   chan bool
	timer   time.Duration
}

// New -
func New(max int64, timer time.Duration) (*Pool, error) {
	if max <= 0 {
		return nil, errSizeTooSmall
	}

	pool := &Pool{
		active:  0,
		max:     max,
		res:     make(chan interface{}, max),
		close:   make(chan bool),
		isClose: false,
		timer:   timer,
	}

	go pool.start()

	return pool, nil
}

func (p *Pool) start() {
	select {
	case <-p.close:
		for {
			p.lock.Lock()
			if int64(len(p.res)) == p.active {
				p.isClose = true
				close(p.res)
				close(p.close)
				p.lock.Unlock()
				return
			}
			p.lock.Unlock()
		}

	}
}

func (p *Pool) create() {
	p.active++
	obj := p.active
	p.res <- obj
}

func (p *Pool) timeout() {
	timer := time.NewTimer(p.timer)
	select {
	case <-timer.C:
		timerCh <- true
	case <-isRequsetEnd:
		return
	}
}

// Get -
func (p *Pool) Get() (interface{}, error) {
	p.lock.Lock()
	defer p.lock.Unlock()

	if p.isClose {
		return nil, errPoolClosed
	}

	if p.active < p.max {
		p.create()
	}

	go p.timeout()
	select {
	case r := <-p.res:
		if p.active < p.max {
			p.create()
		}
		isRequsetEnd <- true
		return r, nil
	case <-timerCh:
		return nil, errTimeout
	}
}

// Put -
func (p *Pool) Put(obj interface{}) error {
	p.lock.Lock()
	defer p.lock.Unlock()

	select {
	case p.res <- obj:
		return nil
	}
}

// Close -
func (p *Pool) Close() bool {
	p.close <- true

	for {
		if p.isClose {
			return true
		}
	}
}
