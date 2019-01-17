package source

// Source defines interface for reading data from source
// as stream
type Source interface {
	Do() error
}

// Config defines configuration for specification of source
type Config struct {
	Path string
}
