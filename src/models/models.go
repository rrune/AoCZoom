package models

type Config struct {
	Port      string   `yaml:"port"`
	ImagePath string   `yaml:"imagePath"`
	Ssl       bool     `yaml:"ssl"`
	Certfile  string   `yaml:"certfile"`
	Keyfile   string   `yaml:"keyfile"`
	Keys      []string `yaml:"keys"`
}
