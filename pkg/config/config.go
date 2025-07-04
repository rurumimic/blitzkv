package config

import (
	"os"
	"path/filepath"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

type Config struct {
	App    AppConfig    `mapstructure:"app"`
	Server ServerConfig `mapstructure:"server"`
	Store  StoreConfig  `mapstructure:"store"`
	Log    LogConfig    `mapstructure:"log"`
}

type AppConfig struct {
	Name    string `mapstructure:"name"`
	Version string `mapstructure:"version"`
	Env     string `mapstructure:"env"`
	Debug   bool   `mapstructure:"debug"`
}

type ServerConfig struct {
	Port int `mapstructure:"port"`
}

type StoreConfig struct {
	MaxKeys      int `mapstructure:"max_keys"`
	MaxValueSize int `mapstructure:"max_value_size"`
}

type LogConfig struct {
	Level string `mapstructure:"level"`
}

func Load() (*Config, error) {
	pflag.String("debug", "false", "Enable debug mode")
	pflag.String("env", "local", "Environment to run the application in (local, dev, prod)")
	pflag.BoolP("help", "h", false, "Display help")
	pflag.CommandLine.SortFlags = false
	pflag.Parse()

	viper.BindPFlags(pflag.CommandLine)
	viper.AutomaticEnv()
	viper.SetConfigType("yaml")

	help := viper.GetBool("help")
	if help {
		pflag.Usage()
		os.Exit(0)
	}

	env := viper.GetString("env")
	wd, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	configPath := filepath.Join(wd, "config", env)
	viper.AddConfigPath(configPath)

	err = viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	var config Config
	err = viper.Unmarshal(&config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
