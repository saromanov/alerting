package alerting

import (
	"fmt"

	"github.com/saromanov/alerting/alerting"
	"github.com/saromanov/alerting/server"
	"github.com/urfave/cli"
)

var flags = []cli.Flag{
	cli.StringFlag{
		EnvVar: "SERVER_ADDRESS",
		Name:   "server-address",
		Usage:  "setup address for alerting app",
	},

	cli.StringFlag{
		EnvVar: "ALERTING_CONFIG",
		Name:   "alerting-config",
		Usage:  "path to alerting config",
	},
}

// setupServer provides setup of the server
func setupServer(c *cli.Context) (*server.Server, error) {
	s, err := server.Create(c.String("github-token"))
	if err != nil {
		return nil, fmt.Errorf("unable to setup server: %v", err)
	}
	return s, nil
}

func parseConfig(path string) (*alerting.Config, error) {
	return
}
func main() {

}
