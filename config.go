package main

import (
	"io/ioutil"
	"log"
	"os"
	"time"

	yaml "gopkg.in/yaml.v2"
)

type State int

const (
	open State = iota
	closed
	failure
)

type Check struct {
	host  Host
	state State
}

type Host struct {
	Network string `yaml:"network"`
	Address string `yaml:"address"`
}

type Config struct {
	Timeout  time.Duration `yaml:"timeout"`
	Interval time.Duration `yaml:"interval"`
	Hosts    []Host        `yaml:"hosts"`
}

func readConfig(config []byte) *Config {
	var c Config
	yaml.Unmarshal(config, &c)
	return &c
}

func verifyFile(path string) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		log.Panicf("The file %v doesn't exist.", path)
	}
}

func readConfigFile(path string) ([]byte, error) {
	verifyFile(path)

	data, err := ioutil.ReadFile(path)

	return data, err
}

func getConfigEnv() (*Config, error) {
	configFile := os.Getenv("CONFIG_FILE")
	data, err := readConfigFile(configFile)

	return readConfig(data), err
}
