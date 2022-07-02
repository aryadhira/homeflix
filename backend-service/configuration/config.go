package configuration

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type Config struct {
	AppServer      string `yaml:"server"`
	AppPort        string `yaml:"port"`
	Database       string `yaml:"db"`
	DatabaseServer string `yaml:"dbserver"`
	DatabasePort   string `yaml:"dbport"`
	DownloadPath   string `yaml:"download"`
	PeerflixPath   string `yaml:"peerflixpath"`
	YtsAPIURL      string `yaml:"ytsurl"`
	Limit          string `yaml:"limit"`
}

func (c *Config) GetConfig() *Config {
	yamlFile, err := ioutil.ReadFile("configuration/app.yaml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return c
}
