package utils

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Server struct {
		Port int `yaml:"port"`
	} `yaml:"server"`
	Grpc struct {
		Port int `yaml:"port"`
	} `yaml:"grpc"`
	Games    []string `yaml:"games"`
	Docker   struct{ Enable bool } `yaml:"docker"`
	Kubernetes struct{ Enable bool } `yaml:"kubernetes"`
}

var AppConfig Config

// LoadConfig reads and decodes the YAML configuration file into AppConfig.
// It terminates the program with a fatal log if the config cannot be loaded or is invalid.
func LoadConfig(configPath string) {
	if configPath == "" {
		configPath = "configs/config.yaml" // Default path if not specified
	}

	file, err := os.Open(configPath)
	if err != nil {
		log.Fatalf("Failed to open config file at '%s': %v", configPath, err)
	}
	defer file.Close()

	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(&AppConfig); err != nil {
		log.Fatalf("Failed to decode config file at '%s': %v", configPath, err)
	}

	validateConfig()
	log.Println("Configuration loaded successfully")
}

// validateConfig checks the essential fields in AppConfig for basic validation.
func validateConfig() {
	if AppConfig.Server.Port == 0 {
		log.Fatalf("Server port is not specified in the configuration file")
	}
	if AppConfig.Grpc.Port == 0 {
		log.Fatalf("gRPC port is not specified in the configuration file")
	}
	if len(AppConfig.Games) == 0 {
		log.Println("Warning: No games are defined in the configuration file")
	}
}
