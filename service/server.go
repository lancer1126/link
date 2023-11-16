package service

type server interface {
	start() error
	stop() error
	services() []Service
}

type serverPool[T server] struct {
	servers map[string]T
}

type baseServer struct {
	ss map[string]Service
}

func newServerPool[T server]() *serverPool[T] {
	return &serverPool[T]{
		servers: make(map[string]T),
	}
}

func (p *serverPool[T]) from(addr string, newServer func() T) T {
	s, exist := p.servers[addr]
	if exist {
		return s
	}
	s = newServer()
	p.servers[addr] = s
	return s
}

func (s *baseServer) services() (ss []Service) {
	for _, s := range s.ss {
		ss = append(ss, s)
	}
	return
}
