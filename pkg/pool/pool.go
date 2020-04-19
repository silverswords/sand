package pool

import (
	"errors"
	"log"
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
	m      sync.Mutex
	active int64
	max    int64
	res    chan interface{}
	close  chan bool
	closed bool
}

// New -
func New(max int64) (*pool, error) {
	if max <= 0 {
		return nil, errors.New("size of the pool is too small")
	}

	log.Printf("create pool")
	return &pool{
		active: 0,
		max:    max,
		res:    make(chan interface{}, max),
		close:  make(chan bool),
		closed: false,
	}, nil
}

func (p *pool) create() (interface{}, error) {
	atomic.AddInt64(&p.active, 1)
	obj := p.active
	return obj, nil
}

func (p *pool) Get() (interface{}, error) {
	p.m.Lock()
	defer p.m.Unlock()

	if p.closed {
		return nil, errors.New("this pool is closed")
	}

	select {
	case r := <-p.res:
		log.Println("get a resource from pool")
		return r, nil
	default:
		if p.active < p.max {
			log.Println("get a new one")
			return p.create()
		}
		return nil, errors.New("pool is empty")
	}
}

func (p *pool) Put(obj interface{}) error {
	p.m.Lock()
	defer p.m.Unlock()

	select {
	case p.res <- obj:
		log.Printf("put conneciton ID: %d back", obj)
		return nil
	default:
		return errors.New("put back error")
	}
}

func (p *pool) Close() error {
	for {
		if int64(len(p.res)) == p.active {
			p.closed = true
			close(p.res)
			log.Println("pool close")
			return nil
		}
	}
}
