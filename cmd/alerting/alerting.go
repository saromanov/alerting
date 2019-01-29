package alerting

import (
	"fmt"
	"os"

	"github.com/saromanov/alerting/alerting"
	"github.com/saromanov/alerting/server"
	"github.com/urfave/cli"
	"gopkg.in/yaml.v2"
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

// parseConfig provides parsing of the config .yml file
func parseConfig(data []byte) (*alerting.Config, error) {
	var c *alerting.Config
	err := yaml.Unmarshal(data, &c)
	if err != nil {
		return nil, fmt.Errorf("unable to parse .born.yml: %v", err)
	}

	return c, nil
}

func main() {
	err := cli.NewApp().Run(os.Args)
	if err != nil {
		panic(err)
	}
}
