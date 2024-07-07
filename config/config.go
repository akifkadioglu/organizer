package config

import (
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

// Config struct
type Config struct {
	AppVersion string `yaml:"app_version"`
	AppName    string `yaml:"app_name"`

	AudioTypes       []string `yaml:"audio_types"`
	VideoTypes       []string `yaml:"video_types"`
	ImageTypes       []string `yaml:"image_types"`
	ApplicationTypes []string `yaml:"application_types"`
	OfficeTypes      []string `yaml:"office_types"`

	AudioFolder       string `yaml:"audio_folder"`
	VideoFolder       string `yaml:"video_folder"`
	ImageFolder       string `yaml:"image_folder"`
	ApplicationFolder string `yaml:"application_folder"`
	OfficeFolder      string `yaml:"office_folder"`
	OthersFolder      string `yaml:"other_folder"`
}

// ReadValue reads the config.yaml file if env variable CONFIG_FILE is not set
func ReadValue() *Config {
	var configs Config
	configFile := os.Getenv("CONFIG_FILE")
	if configFile == "" {
		configFile = "config.yaml"
	}
	filename, _ := filepath.Abs("./" + configFile)
	yamlFile, _ := os.ReadFile(filename)
	yaml.Unmarshal(yamlFile, &configs)

	return &configs
}
