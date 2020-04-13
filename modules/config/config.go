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
	var configuration Config
	var ConfigPath string

	flag.StringVar(&ConfigPath, "c", os.Getenv("HOME")+"/.config/EtherealBot/config.toml", "determine what config to use")
	flag.Parse()

	file, _ := ioutil.ReadFile(ConfigPath)
	toml.Unmarshal(file, &configuration)

	return &configuration
}
