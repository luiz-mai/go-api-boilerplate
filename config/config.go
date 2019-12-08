package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Global GlobalConfig
	HTTP   HTTPConfig
	DB     DBConfig
}

type GlobalConfig struct {
	Name string
}

type HTTPConfig struct {
	Port int
}

type DBConfig struct {
	Host string
	Port int
	User string
	Pass string
	Name string
}

func Read(configPaths ...string) (cfg Config, err error) {
	v := viper.New()
	v.SetConfigName("config")
	v.SetConfigType("toml")
	v.AutomaticEnv()
	v.AddConfigPath("./")
	v.AddConfigPath("./../")
	if err = v.ReadInConfig(); err != nil {
		return cfg, fmt.Errorf("Failed to read the configuration file: %s", err)
	}
	if err = v.Unmarshal(&cfg); err != nil {
		return cfg, err
	}
	return cfg, nil
}
