package configs

import "github.com/spf13/viper"

var CONFIGS *config

type config struct {
	API APIConfig
	DB DBConfig
} 

type APIConfig struct {
	Port string
}

type DBConfig struct {
	Host string
	Port string
	User string
	Pass string
	Database string
}

func init() {
	viper.SetDefault("api.port", "9000")
	viper.SetDefault("database.host", "localhost")
	viper.SetDefault("database.port", "5432")
}

func LoadConfigs() error {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()

	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return err
		}
	}

	CONFIGS = new(config)

	CONFIGS.API = APIConfig{
		Port: viper.GetString("api.port"),
	}
	CONFIGS.DB = DBConfig{
		Host: viper.GetString("database.host"),
		Port: viper.GetString("database.port"),
		User: viper.GetString("database.user"),
		Pass: viper.GetString("database.pass"),
		Database: viper.GetString("database.name"),
	}

	return nil
}

func GetDB() DBConfig {
	return CONFIGS.DB
}

func GetApiServerPort() string {
	return CONFIGS.API.Port
}