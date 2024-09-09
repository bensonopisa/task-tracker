package config

import (
	"log"
	"os"
	"path"

	"gopkg.in/yaml.v3"
)

type Config struct {
	AppCfg    AppConfig    `yaml:"app"`
	SrvConfig ServerConfig `yaml:"server"`
}

type AppConfig struct {
	Name     string         `yaml:"name"`
	DbConfig DatabaseConfig `yaml:"database"`
}

type ServerConfig struct {
	Port string `yaml:"port"`
}

type DatabaseConfig struct {
	DbName     string `yaml:"dbname"`
	DbHost     string `yaml:"dbhost"`
	DbPassword string `yaml:"dbpassword"`
	DbUser     string `yaml:"dbuser"`
	DbPort     uint   `yaml:"dbport"`
	SslMode    string `yaml:"sslmode"`
}

var configPath = "./config/config.yml"

func ReadConfig() *Config {

	var config Config

	dir, err := os.Getwd()

	if err != nil {
		log.Fatalf("error reading directory - %v ", err)
	}

	path := path.Clean(path.Join(dir, configPath))

	log.Printf("Resolving path to %v\n", path)

	b, err := os.ReadFile(path)

	if err != nil {
		log.Fatalf("--- Error %v", err)
	}

	yaml.Unmarshal(b, &config)

	return &config
}
