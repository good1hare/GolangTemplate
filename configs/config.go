package configs

import (
	"github.com/spf13/viper"
	"log"
	"os"
	"path/filepath"
)

type (
	Config struct {
		Http
		Mysql
		Logger
	}

	Http struct {
		Port string
	}

	Mysql struct {
		DriverName   string
		DatabaseName string
		UserName     string
		Password     string
		Host         string
		Port         string
	}

	Logger struct {
		Level string
	}
)

func NewConfig() (*Config, error) {
	v := viper.New()
	v.SetConfigName(".yaml")
	v.AddConfigPath(getProjectRoot())
	v.SetConfigType("yaml")

	//If config file .yaml doesn't exist, use default
	if _, err := os.Stat(".yaml"); os.IsNotExist(err) {
		v.SetConfigName("default.yaml")
	}

	if err := v.ReadInConfig(); err != nil {
		return nil, err
	}

	v.AutomaticEnv()

	var config Config
	if err := v.Unmarshal(&config); err != nil {
		return nil, err
	}

	return &config, nil
}

func getProjectRoot() string {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	for {
		if _, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil {
			return dir
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			break
		}
		dir = parent
	}
	return ""
}
