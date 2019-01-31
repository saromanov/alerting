package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/saromanov/alerting/alerting"
	"github.com/saromanov/alerting/server"
	"github.com/urfave/cli"
	"gopkg.in/yaml.v2"
)

var commands = []cli.Command{
	{
		Name:    "config",
		Aliases: []string{"c"},
		Usage:   "path to .yml config",
		Action: func(c *cli.Context) error {
			fmt.Println("added task: ", c.Args().First())
			return nil
		},
	},
}

// setupServer provides setup of the server
func setupServer(c *cli.Context) error {
	return server.Create(c.String("github-token"))
}

// parseConfig provides parsing of the config .yml file
func parseConfig(path string) (*alerting.Config, error) {
	yamlFile, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("unable to open config file: %v", err)
	}
	var c *alerting.Config
	err = yaml.Unmarshal(yamlFile, &c)
	if err != nil {
		return nil, fmt.Errorf("unable to parse .born.yml: %v", err)
	}

	return c, nil
}

func main() {
	app := cli.NewApp()
	app.Name = "alerting"
	app.Usage = "app for alert handling"
	app.Action = func(c *cli.Context) error {
		configPath := c.String("config")
		if configPath == "" {
			panic("config path is not defined")
		}
		_, err := parseConfig(configPath)
		if err != nil {
			panic(err)
		}

		return nil

	}
	err := app.Run(os.Args)
	if err != nil {
		panic(err)
	}
}
