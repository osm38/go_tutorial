package util

import (
	"os"

	"gopkg.in/yaml.v2"
)

type DbConfig struct {
	Host     string
	User     string
	Password string
	DBName   string
	Port     string
	SLLMode  string
	TimeZone string
}

type LogConfig struct {
	Name     string
	LogFile  string
	LogLevel string
}

func LoadConfig(filename string, obj any) error {
	data, err := os.ReadFile("./config/" + filename)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(data, obj)
	return err
}
