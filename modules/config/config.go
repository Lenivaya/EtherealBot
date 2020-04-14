package config

import (
	"flag"
	"io/ioutil"
	"os"

	"github.com/pelletier/go-toml"
)

type Telegram struct {
	TelegramBotToken string
}

type Config struct {
	Telegram Telegram
}

func GetConfig() *Config {
	var useEnv bool
	var ConfigPath string

	flag.BoolVar(&useEnv, "env", false, "set to true if you want to use environment variables")
	flag.StringVar(&ConfigPath, "c", os.Getenv("HOME")+"/.config/EtherealBot/config.toml", "determine what config to use")
	flag.Parse()

	switch {
	case useEnv:
		return getFromEnv()
	default:
		return getFromFile(ConfigPath)
	}
}

func getFromFile(path string) *Config {
	var configuration Config

	file, _ := ioutil.ReadFile(path)
	toml.Unmarshal(file, &configuration)

	return &configuration
}

func getFromEnv() *Config {
	var configuration Config

	api_token := os.Getenv("TG_API_TOKEN")

	if api_token != "" {
		configuration.Telegram.TelegramBotToken = api_token
	} else {
		os.Exit(1)
	}

	return &configuration
}
