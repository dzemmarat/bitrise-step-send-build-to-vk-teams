package main

import (
	"fmt"
	"github.com/bitrise-io/go-utils/log"
	"github.com/bitrise-tools/go-steputils/stepconf"
	"os"
)

// Config ...
type Config struct {
	// Message
	BotToken      stepconf.Secret `env:"bot_token"`
	ChatID        stepconf.Secret `env:"chat_id"`
	CommitMessage string          `env:"commit_message"`
	CommitAuthor  string          `env:"commit_author"`
	FileURL       string          `env:"BITRISE_APP_URL"`
}

func validate(conf *Config) error {
	if conf.BotToken == "" {
		return fmt.Errorf("if you want to use a bot to send a message provide the bot API token")
	}
	return nil
}

func main() {
	var conf Config
	if err := stepconf.Parse(&conf); err != nil {
		log.Errorf("Error: %s\n", err)
		os.Exit(1)
	}
	stepconf.Print(conf)

	if err := validate(&conf); err != nil {
		log.Errorf("Error: %s\n", err)
		os.Exit(1)
	}

	if err := sendMessage(conf); err != nil {
		log.Errorf("Error: %s", err)
		os.Exit(1)
	}

	log.Donef("\nVkTeams message successfully sent! 🚀\n")
	os.Exit(0)
}
