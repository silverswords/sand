package logger

type Option func(*Builder)

func WithLogPath(logpath string) Option {
	return func(b *Builder) {b.Config.LogPath = logpath}
}

func WithLogFileName(logFileName string) Option {
	return func(b *Builder) {b.Config.LogFileName = logFileName}
}

func WithLogLevel(logLevel string) Option {
	return func(b *Builder) {b.Config.LogLevel = logLevel}
}

func WithLogTimeLayOut(logTimeOut string) Option {
	return func(b *Builder) {b.Config.LogTimeLayOut = logTimeOut}
}

func WithLogFileMaxSize(logFileMaxSize int) Option {
	return func(b *Builder) {b.Config.LogFileMaxSize = logFileMaxSize}
}

func WithLogFileMaxBackups(logFileMaxBackups int) Option {
	return func(b *Builder) {b.Config.LogFileMaxBackups = logFileMaxBackups}
}

func WithLogMaxAge(logMaxAge int) Option {
	return func(b *Builder) {b.Config.LogMaxAge = logMaxAge}
}

func WithLogCompress(logCompress bool) Option {
	return func(b *Builder) {b.Config.LogCompress = logCompress}
}

func WithLogStdout(logStdout bool) Option {
	return func(b *Builder) {b.Config.LogStdout = logStdout}
}
