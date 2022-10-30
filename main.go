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
	BotToken stepconf.Secret `env:"bot_token"`
	ChatID   stepconf.Secret `env:"chat_id"`
	Version  string          `env:"app_version"`

	// Success message
	CommitMessage string `env:"commit_message"`
	CommitAuthor  string `env:"commit_author"`
	FileURL       string `env:"file_url"`
}

// success is true if the build is successful, false otherwise.
var success = os.Getenv("BITRISE_BUILD_STATUS") == "0"

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

	if success {
		if err := sendSuccessMessage(conf); err != nil {
			log.Errorf("Error: %s", err)
			os.Exit(1)
		}
	} else {
		if err := sendErrorMessage(conf); err != nil {
			log.Errorf("Error: %s", err)
			os.Exit(1)
		}
	}

	log.Donef("\nVkTeams message successfully sent! 🚀\n")
	os.Exit(0)
}
