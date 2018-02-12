package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/BurntSushi/toml"
	"github.com/urfave/cli"
)

// Config from config directory
type Config struct {
	Reviewers []ReviewerConfig `toml:"Reviewer"`
	Slack     SlackConfig      `toml:"Slack"`
}

// ReviewerConfig from config.toml
type ReviewerConfig struct {
	Name         string `toml:"name"`
	SlackAccount string `toml:"slack_account"`
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
		loadConfig()

		reviewer := selectReviewer()
		message := generateMessage(reviewer, c)

		// TODO: send message to slack
		fmt.Println(message)

		return nil
	}

	app.Run(os.Args)
}

func loadConfig() {
	var file = "./config/config.toml"
	_, err := toml.DecodeFile(file, &config)

	if err != nil {
		panic(err)
	}
}

func selectReviewer() ReviewerConfig {
	reviewerCount := len(config.Reviewers)
	rand.Seed(time.Now().UnixNano())
	reviewerNum := rand.Intn(reviewerCount)

	return config.Reviewers[reviewerNum]
}

func generateMessage(reviewer ReviewerConfig, c *cli.Context) string {
	return reviewer.SlackAccount + " " + reviewer.Name + "さん、コードレビューをお願いします！ " + c.Args().First()
}
