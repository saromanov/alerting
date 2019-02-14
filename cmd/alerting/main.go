package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/saromanov/alerting/server"
	"github.com/saromanov/alerting/storage/redis"
	"github.com/saromanov/alerting/structs"
	"github.com/urfave/cli"
	"gopkg.in/yaml.v2"
)

// setupServer provides setup of the server
func setupServer(c *structs.Config) error {
	fmt.Println("setup server...")
	return server.Create(c)
}

// setupStorage at this moment only Redis
func setupStorage(c *structs.Config) error {
	fmt.Println("setup storage...")
	_, err := redis.Setup(c)
	if err != nil {
		return err
	}
	return nil
}

// parseConfig provides parsing of the config .yml file
func parseConfig(path string) (*structs.Config, error) {
	yamlFile, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("unable to open config file: %v", err)
	}
	var c *structs.Config
	err = yaml.Unmarshal(yamlFile, &c)
	if err != nil {
		return nil, fmt.Errorf("unable to parse config file: %v", err)
	}

	return c, nil
}

func setupApp(c *structs.Config) error {
	if err := setupStorage(c); err != nil {
		return err
	}
	setupServer(c)
	return nil
}

func main() {
	app := cli.NewApp()
	app.Name = "alerting"
	app.Usage = "app for alert handling"
	app.Commands = []cli.Command{
		{
			Name:    "config",
			Aliases: []string{"c"},
			Usage:   "path to .yml config",
			Action: func(c *cli.Context) error {
				configPath := c.Args().First()
				config, err := parseConfig(configPath)
				if err != nil {
					panic(err)
				}
				if err := setupApp(config); err != nil {
					panic(err)
				}
				return nil
			},
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		panic(err)
	}
}
