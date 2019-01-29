package alerting

import (
	"github.com/urfave/cli"
)

var flags = []cli.Flag{
	cli.StringFlag{
		EnvVar: "SERVER_ADDRESS",
		Name:   "github-token",
		Usage:  "using of github token for test",
	},
}

func setupServer(c *cli.Context)
func main(){

}