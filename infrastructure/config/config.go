/*
Package config provides config structures and parse funcs.
*/
package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

// Config is the global configuration.
type Config struct {
	App   ApplicationConfig `json:"app" yaml:"app"`
	Cache CacheConfig       `json:"cache" yaml:"cache"`
	DB    DBConfig          `json:"db" yaml:"db"`
}

// DBConfig is the configuration of databases.
type DBConfig struct {
	FileName    string `json:"file_name" yaml:"file_name"`
	DSN         string `json:"dsn" yaml:"dsn"`
	MongoURI    string `json:"mongo_uri" yaml:"mongo_uri"`
	MongoDBName string `json:"mongo_db_name" yaml:"mongo_db_name"`
}

// ApplicationConfig is the configuration of main app.
type ApplicationConfig struct {
	Port        int    `json:"port" yaml:"port"`
	PageSize    int    `json:"page_size" yaml:"page_size"`
	TokenSecret string `json:"token_secret" yaml:"token_secret"`
	TokenHours  int    `json:"token_hours" yaml:"token_hours"`
}

// CacheConfig is the configuration of cache.
type CacheConfig struct {
	Address  string `json:"address" yaml:"address"`
	Password string `json:"password" yaml:"password"`
	DB       int    `json:"db" yaml:"db"`
	Timeout  int    `json:"timeout" yaml:"timeout"`
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
