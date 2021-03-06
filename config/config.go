package config

import (
	"fmt"
	"io/ioutil"
	"strconv"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

func (cnf *Config) Address() string {
	return cnf.Host + ":" + strconv.Itoa(cnf.Port)
}

func Load(env string) Config {
	filepath := fmt.Sprintf("./config/config.%s.yaml", env)
	fmt.Printf("[Info]Config file path: %s\n", filepath)

	configBytes, err := ioutil.ReadFile(filepath)
	if err != nil {
		fmt.Printf("[Error]Load config file error: %s\n", err.Error())
		return DefaultConfig()
	}
	fmt.Printf("[Info]Load config file: %s\n", string(configBytes))

	var config Config
	err = yaml.Unmarshal(configBytes, &config)
	if err != nil {
		fmt.Printf("[Error]Unmarshal config error: %s\n", err.Error())
		return DefaultConfig()
	}
	fmt.Printf("[Info]Unmarshal config done\n")

	return config
}

func DefaultConfig() Config {
	return Config{
		Port: 8080,
	}
}
