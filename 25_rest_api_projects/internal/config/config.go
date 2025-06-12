package config

import (
	"flag"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type HTTPServer struct {
	Addr string
}

type Config struct {
	Env         string `yaml:"env" env:"ENV" env-required:"true" env-default:"production"`
	StoragePath string `yaml:"storage_path" env:"STORAGE_PATH" env-required:"true"`
	HTTPServer  `yaml:"http_server" env:"HTTP_SERVER" env-required:"true"`
	CORS        struct {
		AllowOrigin      []string `yaml:"allow_origin" env:"ALLOW_ORIGIN" env-required:"true"`
		AllowMethods     []string `yaml:"allow_methods" env:"ALLOW_METHODS" env-required:"true"`
		AllowHeaders     []string `yaml:"allow_headers" env:"ALLOW_HEADERS" env-required:"true"`
		ExposeHeaders    []string `yaml:"expose_headers" env:"EXPOSE_HEADERS"`
		MaxAge           int      `yaml:"max_age" env:"MAX_AGE" env-required:"true"`
		AllowCredentials bool     `yaml:"allow_credentials" env:"ALLOW_CREDENTIALS" env-required:"true"`
	} `yaml:"cors" env:"CORS" env-required:"true"`
}

func MustLoad() *Config {

	var configPath string

	configPath = os.Getenv("CONFIG_PATH")

	if configPath == "" {
		flags := flag.String("config", "", "Path to the configuration file")
		flag.Parse()
		configPath = *flags

		if configPath == "" {
			log.Fatal("CONFIG_PATH environment variable or --config flag must be set")
		}

	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("Configuration file does not exist: %s", configPath)
	}

	var cfg Config

	err := cleanenv.ReadConfig(configPath, &cfg)

	if err != nil {
		log.Fatalf("Error reading environment variables: %s", err.Error())
	}

	return &cfg
}
