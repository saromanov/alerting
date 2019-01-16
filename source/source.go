package source

// Source defines interface for reading data from source
// as stream
type Source interface {
	Get(*Config) error
}

// Config defines configuration for specification of source
type Config struct {

}