package alerting

import (
	"fmt"
	"github.com/urfave/cli"
	"github.com/saromanov/alerting/server"
)

var flags = []cli.Flag{
	cli.StringFlag{
		EnvVar: "SERVER_ADDRESS",
		Name:   "github-token",
		Usage:  "using of github token for test",
	},
}

// setupServer provides setup of the server
func setupServer(c *cli.Context) (server*Server, error) {
	s, err := server.Create(c.String("github-token"))
	if err != nil {
		return nil, fmt.Errorf("unable to setup server: %v", err)
	}
	return s, nil
}
func main(){

}