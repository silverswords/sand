package web

type Option func(*Builder)

func WithHost(h string) Option {
	return func(b *Builder) {
		b.Config.Host = h
	}
}

func WithAddr(addr string) Option {
	return func(b *Builder) {
		b.Config.Addr = addr
	}
}
