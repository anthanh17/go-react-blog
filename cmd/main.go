package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/anthanh17/go-react-blog/pkg/config"
	"github.com/anthanh17/go-react-blog/pkg/db"
	"github.com/anthanh17/go-react-blog/pkg/db/repository"
	"github.com/spf13/viper"
)

var (
	conf = config.Config{}
	file string
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

func parse() bool {
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

func main() {
	if !parse() {
		showHelp()
		os.Exit(-1)
	}

	db.SetupConfig(conf)

	err := repository.SetupDb()
	if err != nil {
		log.Fatal(err)
	}
	// if no error. Ping is successful
	fmt.Println("Ping to database successful, connection is still alive")
}
