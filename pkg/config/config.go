package config

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	MySQLConfig    *MySQLConfig    `mapstructure:"mysql"`
	TelegramConfig *TelegramConfig `mapstructure:"telegram"`
}

type MySQLConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	DbName   string `mapstructure:"db-name"`
}

type TelegramConfig struct {
	Token  string `mapstructure:"token"`
	ChatId int64  `mapstructure:"chat-id"`
	BotId  int64  `mapstructure:"bot-id"`
}

type Settings struct {
	conf        Config
	file        string
	MySqlCfg    *MySQLConfig
	TelegramCfg *TelegramConfig
}

func showHelp() {
	fmt.Printf("Usage:%s {params}\n", os.Args[0])
	fmt.Println("	  -c {config file}")
	fmt.Println("     -h (show help info)")
}

func LoadConfig() (settings Settings, err error) {
	// parse flag
	var file string
	flag.StringVar(&file, "c", "conf/conf.toml", "config file")
	help := flag.Bool("h", false, "help info")
	flag.Parse()
	if *help {
		showHelp()
	}

	// load file
	if _, error := os.Stat(file); error != nil {
		return
	}
	viper.SetConfigFile(file)
	viper.SetConfigType("toml")

	if err = viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
		return
	}

	if err = viper.Unmarshal(&settings.conf); err != nil {
		fmt.Printf("Unable to decode into struct, %v", err)
		return
	}

	// Setup Config
	if settings.conf.MySQLConfig == nil {
		log.Fatal(errors.New(fmt.Sprintf("Can not load MySQL config!")))
		return
	}
	settings.MySqlCfg = settings.conf.MySQLConfig

	if settings.conf.TelegramConfig == nil {
		log.Fatal(errors.New(fmt.Sprintf("Can not load telegram config !")))
		return
	}
	settings.TelegramCfg = settings.conf.TelegramConfig
	return
}
