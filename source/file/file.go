package file

import (
	"fmt"

	"github.com/hpcloud/tail"
	"github.com/saromanov/alerting/source"
)

type client struct {
	t *tail.Tail
	c *source.Config
}

// New provides initialization of the file watcher method
func New(s *source.Config) (source.Source, error) {
	t, err := tail.TailFile("/var/log/nginx.log", tail.Config{Follow: true})
	if err != nil {
		return nil, err
	}

	return client{
		t: t,
		c: s,
	}, nil
}

// Do provides handling of the log output
func (c client) Do() error {
	for line := range c.t.Lines {
		fmt.Println(line.Text)
	}
	return nil
}
