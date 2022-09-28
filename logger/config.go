package logger

type Config struct {
	LogPath           string
	LogFileName       string
	LogLevel          string
	LogTimeLayOut     string
	LogFileMaxSize    int
	LogFileMaxBackups int
	LogMaxAge         int
	LogCompress       bool
	LogStdout         bool
	Development       bool
}
