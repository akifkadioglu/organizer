package config

import (
	"gopkg.in/yaml.v3"
)

// Config struct
type Config struct {
	AppVersion   string `yaml:"app_version"`
	AppName      string `yaml:"app_name"`
	DatabasePath string `yaml:"database_path"`

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

const configYaml = `
app_name: Organizer
app_version: 0.0.1
database_path: "/home/akif/organizer/organizer.db"

audio_folder: "audios"
video_folder: "videos"
image_folder: "images"
application_folder: "applications"
office_folder: "offices"
other_folder: "others"


office_types: ["xls","xlsx","doc","docx","ppt","pptx","pdf","odt","ods","odp","txt","rtf","csv"]
image_types: ["jpeg", "jpg", "png", "gif", "bmp", "svg", "webp", "tiff", "tif"]
audio_types: ["mp3", "wav", "ogg", "flac", "aac", "m4a"]
application_types: ["exe", "apk", "msi", "dmg", "bin"]
video_types: ["mp4", "avi", "mov", "mkv", "wmv", "flv", "webm"]
`

// ReadValue reads the config.yaml file if env variable CONFIG_FILE is not set
func ReadValue() *Config {
	var configs Config
	yaml.Unmarshal([]byte(configYaml), &configs)

	return &configs
}
