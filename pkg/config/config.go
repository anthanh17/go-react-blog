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

var (
	conf        Config
	file        string
	MySqlCfg    *MySQLConfig
	TelegramCfg *TelegramConfig
)

func showHelp() {
	fmt.Printf("Usage:%s {params}\n", os.Args[0])
	fmt.Println("	  -c {config file}")
	fmt.Println("     -h (show help info)")
}

func load() bool {
	_, err := os.Stat(file)
	if err != nil {
		return false
	}

	viper.SetConfigFile(file)
	viper.SetConfigType("toml")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
		return false
	}

	if err = viper.Unmarshal(&conf); err != nil {
		fmt.Printf("Unable to decode into struct, %v", err)
	}

	fmt.Printf("config %s load ok!\n", file)
	// Reading variables using the model
	fmt.Println("Reading variables using the model..")
	fmt.Println("Database Host:\t", conf.MySQLConfig.Host)
	fmt.Println("Database Port:\t", conf.MySQLConfig.Port)
	fmt.Println("Database Db-name:", conf.MySQLConfig.DbName)
	fmt.Println("Database username:", conf.MySQLConfig.User)
	fmt.Println("Database password:", conf.MySQLConfig.Password)
	return true
}

func parseFlag() bool {
	flag.StringVar(&file, "c", "conf/conf.toml", "config file")
	help := flag.Bool("h", false, "help info")
	flag.Parse()
	if !load() {
		return false
	}

	if *help {
		showHelp()
		return false
	}
	return true
}

func loadConfig() Config {
	if !parseFlag() {
		showHelp()
		os.Exit(-1)
	}
	return conf
}

func SetupConfig() {
	cfg := loadConfig()
	if cfg.MySQLConfig == nil {
		log.Fatal(errors.New(fmt.Sprintf("Can not load MySQL config!")))
		return
	}
	MySqlCfg = cfg.MySQLConfig

	if cfg.TelegramConfig == nil {
		log.Fatal(errors.New(fmt.Sprintf("Can not load telegram config !")))
		return
	}
	TelegramCfg = cfg.TelegramConfig
}
