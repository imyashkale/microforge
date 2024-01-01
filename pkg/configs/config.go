package configs

import (
	"github.com/imyashkale/microforge/pkg/log"
	"github.com/spf13/viper"
)

type Config struct {
	AppConfigPath string
	Application   Application
	Dynamodb      DynamoDB
	Services      []Services
	AWS           AWS
}

func New() *Config {
	return &Config{}
}

func (c Config) Validate() error {
	return nil
}

func (c *Config) Load(logger log.Logger) error {
	

	// Viper setup
	viper.AutomaticEnv() // Automatically read environment variables

	// Setting the environment variable keys with prefixes
	viper.SetEnvPrefix("APP")

	// Name of the config file (without extension)
	viper.SetConfigType("yaml")

	c.AppConfigPath = "../configs/local-config.yaml"
	
	// seting configuration file
	viper.SetConfigFile(c.AppConfigPath)

	// Reading the config file
	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	// Unmarshal the configuration into the Config struct
	if err := viper.Unmarshal(&c); err != nil {
		return err
	}
	return nil
}
