package config

type Config struct {
	MySQLConfig *MySQLConfig `mapstructure:"mysql"`
}

type MySQLConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	DbName   string `mapstructure:"db-name"`
	DbSource string `mapstructure:"source"`
}
