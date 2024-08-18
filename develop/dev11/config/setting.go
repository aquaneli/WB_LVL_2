package config

import (
	"os"

	"gopkg.in/yaml.v2"
)

// Config структура в которую будут десериализоваться данные из config .yaml
type Config struct {
	Server struct {
		IP   string `yaml:"ip"`
		Port string `yaml:"port"`
	} `yaml:"server"`
}

// ReadConfig cчитывает данные из конфига .yaml
func ReadConfig(filename string) (*Config, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
