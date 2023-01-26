package config

import (
	"fmt"

	"github.com/spf13/viper"
)

var config *configuration

type dbConfig struct {
	Host     string
	Port     string
	Database string
	User     string
	Password string
}

type appConfig struct {
	Port string
}

type configuration struct {
	API appConfig
	DB  dbConfig
}

// Returns the application port
func GetAppPort() string {
	return config.API.Port
}

// TODO: Mover isto para o contexto de bd
// Returns the connection string formated to pgsql
func (d *dbConfig) GetConnectionString() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s database=%s sslmode=disable", d.Host, d.Port, d.User, d.Password, d.Database)
}

func init() {

	viper.SetDefault("api.port", "8000")
	viper.SetDefault("db.host", "localhost")
	viper.SetDefault("db.port", "5432")

}

// Loads the configuration from config file
func Load() error {

	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath("../config/")
	err := viper.ReadInConfig()
	if err != nil { // Handle errors reading the config file
		return fmt.Errorf("fatal error config file: %w", err)
	}

	config = new(configuration)

	config.API = appConfig{
		Port: viper.GetString("api.port"),
	}

	config.DB = dbConfig{
		Host:     viper.GetString("database.host"),
		Port:     viper.GetString("database.port"),
		Database: viper.GetString("database.database"),
		User:     viper.GetString("database.user"),
		Password: viper.GetString("database.password"),
	}

	return nil

}

func GetDB() dbConfig {
	return config.DB
}
