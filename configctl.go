package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os/user"
)

func main() {
	var c Config
	c.getConfigs()

	// split configs up and read each into their own Config

	fmt.Println(c)
}

type Config struct {
	Name    string `yaml:"name"`
	Kind    string `yaml:"kind"`
	Configs []struct {
		Variable string   `yaml:"variable"`
		Values   []string `yaml:"values"`
	} `yaml:"configs"`
}

func (c *Config) getConfigs() *Config {
	usr, err := user.Current()
	check(err, "Current User")
	yamlFile, err := ioutil.ReadFile(usr.HomeDir + "/.configctl.yaml")
	check(err, "Read File")
	err = yaml.Unmarshal(yamlFile, c)
	check(err, "Unmarshal")

	return c
}

func check(e error, m string) {
	if e != nil {
		log.Fatalf(m + ": %v", e)
	}
}
