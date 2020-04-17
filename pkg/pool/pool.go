package pool

import (
	"errors"
	"log"
	"os"
	"os/signal"
	"sync"
	"sync/atomic"
	"syscall"
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
	signal chan os.Signal
	close  chan os.Signal
	closed bool
}

type obj struct {
	ID   int64
	Name string
	Age  int
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
		signal: make(chan os.Signal),
		close:  make(chan os.Signal),
		closed: false,
	}, nil
}

func (p *pool) create() {
	atomic.AddInt64(&p.active, 1)
	p.res <- &obj{
		ID:   p.active,
		Name: "wbofeng",
		Age:  22,
	}
	log.Printf("create obj ID: %d", p.active)
}

func (p *pool) Get() (interface{}, error) {
	p.m.Lock()
	defer p.m.Unlock()

	if p.closed {
		return nil, errors.New("this pool is closed")
	}
	log.Printf("here")
	if p.active < p.max {
		go p.create()
	}
	select {
	case r := <-p.res:
		log.Println("get a resource")
		return r, nil
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

func (p *pool) Close() error {
	for {
		if p.active == int64(len(p.res)) {
			signal.Notify(p.signal, syscall.SIGINT)
			return nil
		}
	}

}
