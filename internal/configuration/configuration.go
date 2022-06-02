package configuration

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Configuration struct {
	Server Server `yaml:"server"`
}
type Server struct {
	Type string `yaml:"type"`
	Port string `yaml:"port"`
}

func GetConfig() *Configuration {
	config := &Configuration{}
	cfgFile, err := ioutil.ReadFile("./config.yaml")

	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(cfgFile, config)

	if err != nil {
		panic(err)
	}

	return config
}
