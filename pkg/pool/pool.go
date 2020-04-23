package pool

import (
	"errors"
	"sync"
)

// Pool -
type Pool struct {
	lock    sync.Mutex
	active  int64
	max     int64
	res     chan interface{}
	isClose bool
	close   chan bool
}

// New -
func New(max int64) (*Pool, error) {
	if max <= 0 {
		return nil, errors.New("size of the pool is too small")
	}

	pool := &Pool{
		active:  0,
		max:     max,
		res:     make(chan interface{}, max),
		close:   make(chan bool),
		isClose: false,
	}

	go pool.start()

	return pool, nil
}

func (p *Pool) start() {
	select {
	case <-p.close:
		p.isClose = true
		for {
			p.lock.Lock()
			if int64(len(p.res)) == p.active {
				close(p.res)
				close(p.close)
				p.lock.Unlock()
				return
			}
			p.lock.Unlock()
		}

	}
}

func (p *Pool) create() interface{} {
	p.active++
	obj := p.active
	return obj
}

// Get -
func (p *Pool) Get() (interface{}, error) {
	p.lock.Lock()
	defer p.lock.Unlock()

	if p.isClose {
		return nil, errors.New("this pool is closed")
	}

	select {
	case r := <-p.res:
		return r, nil
	default:
		if p.active < p.max {
			obj := p.create()
			return obj, nil
		}
	}

	return nil, errors.New("pool is empty")
}

// Put -
func (p *Pool) Put(obj interface{}) error {
	p.lock.Lock()
	defer p.lock.Unlock()

	select {
	case p.res <- obj:
		return nil
	default:
		return errors.New("put back error")
	}
}

// Close -
func (p *Pool) Close() error {
	p.close <- true
	return nil
}
