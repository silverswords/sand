package mysql

type Option func(*Builder)

func WithHost(h string) Option {
	return func(b *Builder) {
		b.Config.Host = h
	}
}

func WithAddr(port string) Option {
	return func(b *Builder) {
		b.Config.Port = port
	}
}
