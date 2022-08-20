package sand

import (
	"sync"

	"github.com/silverswords/sand/server"
)

type Sandbox struct {
	Mu sync.Mutex

	Server server.Server
}

func Instance() *Sandbox {
	return &Sandbox{}
}

func (s *Sandbox) Load(srv server.Server) *Sandbox {
	s.Mu.Lock()
	defer s.Mu.Unlock()

	s.Server = srv

	return s
}

func (s *Sandbox) Run() error {
	return s.Server.Start()
}
