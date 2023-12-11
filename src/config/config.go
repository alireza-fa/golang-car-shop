package config

import (
	"fmt"
	"github.com/dvln/yaml"
	"log"
	"os"
	"time"
)

type Config struct {
	Server struct {
		Port    int    `yaml:"port"`
		RunMode string `yaml:"runMode"`
	}
	Logger struct {
		FilePath string `yaml:"filePath"`
		Encoding string `yaml:"Encoding"`
		Level    string `yaml:"level"`
	}
	Cors struct {
		AllowOrigins string `yaml:"allowOrigins"`
	}
	Postgres struct {
		Host            string        `yaml:"host"`
		Port            int           `yaml:"port"`
		User            string        `yaml:"user"`
		Password        string        `yaml:"password"`
		DbName          string        `yaml:"dbName"`
		SslMode         string        `yaml:"sslMode"`
		MaxIdleConns    int           `yaml:"maxIdleConns"`
		MaxOpenConns    int           `yaml:"maxOpenConns"`
		ConnMaxLifetime time.Duration `yaml:"connMaxLifetime"`
	}
	Redis struct {
		Host               string        `yaml:"host"`
		Port               int           `yaml:"port"`
		Password           string        `yaml:"password"`
		Db                 int           `yaml:"db"`
		DialTimeout        time.Duration `json:"dialTimeout"`
		ReadTimeout        time.Duration `json:"readTimeout"`
		WriteTimeout       time.Duration `json:"writeTimeout"`
		IdleCheckFrequency time.Duration `json:"idleCheckFrequency"`
		PoolSize           int           `json:"poolSize"`
		PoolTimeout        time.Duration `json:"poolTimeout"`
	}
	Password struct {
		IncludeChars     bool `yaml:"includeChars"`
		IncludeDigits    bool `yaml:"includeDigits"`
		MinLength        int  `yaml:"minLength"`
		MaxLength        int  `yaml:"maxLength"`
		IncludeUppercase bool `yaml:"includeUppercase"`
		IncludeLowercase bool `yaml:"includeLowercase"`
	}
}

func GetConfig() *Config {
	cfgPath := getConfigPath(os.Getenv("APP_ENV"))
	b, err := LoadConfig(cfgPath, "yml")
	if err != nil {
		log.Fatalf("Error in load config %v", err)
	}

	cfg, err := ParseConfig(b)
	if err != nil {
		log.Fatalf("Error in parse config %v", err)
	}

	return cfg
}

func ParseConfig(b []byte) (*Config, error) {
	var cnf Config
	err := yaml.Unmarshal(b, &cnf)
	if err != nil {
		fmt.Printf("Erro in parse Config: %v", err)
	}
	return &cnf, nil
}

func LoadConfig(filename string, fileType string) ([]byte, error) {
	yamlFile, err := os.ReadFile(filename + "." + fileType)
	if err != nil {
		return nil, err
	}
	return yamlFile, nil
}

func getConfigPath(env string) string {
	if env == "docker" {
		return "/app/config/config-docker"
	} else if env == "production" {
		return "config/config-production"
	} else {
		return "config/config-development"
	}
}
