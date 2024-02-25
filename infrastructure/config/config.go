package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	App ApplicationConfig `json:"app" yaml:"app"`
	DB  DBConfig          `json:"db" yaml:"db"`
}

type DBConfig struct {
	FileName    string `json:"file_name" yaml:"file_name"`
	DSN         string `json:"dsn" yaml:"dsn"`
	MongoURI    string `json:"mongo_uri" yaml:"mongo_uri"`
	MongoDBName string `json:"mongo_db_name" yaml:"mongo_db_name"`
}

type ApplicationConfig struct {
	Port int `json:"port" yaml:"port"`
}

// Parse parses config file and returns a Config.
func Parse(filename string) (*Config, error) {
	buf, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	c := &Config{}
	err = yaml.Unmarshal(buf, c)
	if err != nil {
		return nil, fmt.Errorf("failed to parse file %s: %v", filename, err)
	}
	return c, nil
}
