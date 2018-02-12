package main

import (
	"fmt"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/urfave/cli"
)

type Config struct {
	Reviewer ReviewerConfig
}

type ReviewerConfig struct {
	Name    string
	account string
}

var config Config

func main() {
	app := cli.NewApp()
	app.Name = "codereview"
	app.Usage = "please code review!"
	app.Action = func(c *cli.Context) error {
		_, err := toml.DecodeFile("./config/config.toml", &config)
		fmt.Println("boom! I say!")
		return nil
	}

	app.Run(os.Args)
}
