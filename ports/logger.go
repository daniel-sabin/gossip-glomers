package ports

type Logger interface {
	Debug(s string)
	Close()
}
