package models

type Config struct {
	Port      string   `yaml:"port"`
	ImagePath string   `yaml:"imagePath"`
	Keys      []string `yaml:"keys"`
}
