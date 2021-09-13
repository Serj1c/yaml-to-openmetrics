package util

import "github.com/spf13/viper"

//Config stores all configuration information
type Config struct {
	ServerPort string `mapstructure:"SERVER_PORT"`
}

// LoadConfig ...
func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("api")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	if err = viper.ReadInConfig(); err != nil {
		return
	}
	err = viper.Unmarshal(&config)
	return
}
