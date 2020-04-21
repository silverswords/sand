package pool

import (
	"errors"
	"sync"
	"sync/atomic"
)

// Pool -
type Pool interface {
	Get() (interface{}, error)
	Put(interface{}) error
	Close() error
}

type pool struct {
	m       sync.Mutex
	wg      sync.WaitGroup
	active  int64
	max     int64
	res     chan interface{}
	isClose bool
	closed  chan bool
}

// New -
func New(max int64) (*pool, error) {
	if max <= 0 {
		return nil, errors.New("size of the pool is too small")
	}

	pool := &pool{
		active:  0,
		max:     max,
		res:     make(chan interface{}, max),
		closed:  make(chan bool),
		isClose: false,
	}
	return pool, nil
}

func (p *pool) create() {
	atomic.AddInt64(&p.active, 1)
	obj := p.active
	p.res <- obj
	p.wg.Done()
}

func (p *pool) Get() (interface{}, error) {
	p.m.Lock()
	defer p.m.Unlock()

	if p.isClose {
		return nil, errors.New("this pool is closed")
	}

	select {
	case r := <-p.res:
		return r, nil
	default:
		if p.active < p.max {
			p.wg.Add(1)
			go p.create()

			select {
			case r := <-p.res:
				return r, nil
			}
		}
		return nil, errors.New("pool is empty")
	}
}

func (p *pool) Put(obj interface{}) error {
	p.m.Lock()
	defer p.m.Unlock()

	select {
	case p.res <- obj:
		return nil
	default:
		return errors.New("put back error")
	}
}

func (p *pool) wait() {
	for {
		if int64(len(p.res)) == p.active {
			p.closed <- true
			return
		}
	}
}

func (p *pool) Close() error {
	p.isClose = true
	p.wg.Wait()

	go p.wait()

	<-p.closed

	close(p.res)
	close(p.closed)
	return nil
}
