package config

type Config struct {
	App ApplicationConfig `json:"app" yaml:"app"`
	DB  DBConfig          `json:"db" yaml:"db"`
}

type DBConfig struct {
	FileName string `json:"file_name" yaml:"file_name"`
}

type ApplicationConfig struct {
	Port int `json:"port" yaml:"port"`
}
