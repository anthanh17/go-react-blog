package config

import (
	"fmt"
	"github.com/spf13/viper"
)

// Configurations exported
type Config struct {
	// Server       ServerConfigurations	
	Database     DatabaseConfigurations	`mapstructure:"mysql"`
	// EXAMPLE_PATH string
	// EXAMPLE_VAR  string
}

// ServerConfigurations exported
type ServerConfigurations struct {
	DbPort int
}

// DatabaseConfigurations exported
type DatabaseConfigurations struct {
	DbSource		string	`mapstructure:"source"`
	DbDriver		string	`mapstructure:"driver"`
	DbName     		string	`mapstructure:"name"`
	DbUsername		string	`mapstructure:"username"`
	DbPassword		string	`mapstructure:"password"`
	DbPort			int		`mapstructure:"port"`
}

func LoadDatabaseConfig(path string) (config Config, err error) {
	viper := viper.New()

	viper.AddConfigPath(path)
	viper.SetConfigName("config")
	viper.SetConfigType("toml")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
	}

	var configuration Config

	if err = viper.Unmarshal(&configuration); err != nil {
		fmt.Printf("Unable to decode into struct, %v", err)
	}

	// Reading variables using the model
	fmt.Println("Reading variables using the model..")
	fmt.Println("Database source is\t", configuration.Database.DbSource)
	fmt.Println("Database name is\t", configuration.Database.DbName)
	fmt.Println("Database username is\t", configuration.Database.DbUsername)
	fmt.Println("Database password is\t\t", configuration.Database.DbPassword)

	return configuration, err
}