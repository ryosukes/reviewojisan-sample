package main

import (
	"fmt"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/urfave/cli"
)

// Config from config directory
type Config struct {
	Reviewer []ReviewerConfig `toml:"Reviewer"`
	Slack    SlackConfig      `toml:"Slack"`
}

// ReviewerConfig from config.toml
type ReviewerConfig struct {
	Name    string `toml:"name"`
	Account string `toml:"account"`
}

// SlackConfig from config.toml
type SlackConfig struct {
	Channel string
}

var config Config

func main() {
	app := cli.NewApp()
	app.Name = "codereview"
	app.Usage = "please code review!"
	app.Action = func(c *cli.Context) error {
		var file = "./config/config.toml"
		_, err := toml.DecodeFile(file, &config)

		if err != nil {
			panic(err)
		}

		fmt.Println(config)
		return nil
	}

	app.Run(os.Args)
}
