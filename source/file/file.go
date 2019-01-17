package file

import (
	"github.com/saromanov/alerting/source"
	"github.com/hpcloud/tail"
)
type client struct {
	t *tail.Tail
}

// New provides initialization of the file watcher method
func New(s *source.Config) (client, error){
	t, err := tail.TailFile("/var/log/nginx.log", tail.Config{Follow: true})
	if err != nil {
		return nil, err
	}

	return client{
		t:t,
	}, nil
}