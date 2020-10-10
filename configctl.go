package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os/user"
	"bytes"
)

func main() {

	configs := getConfigs()

	selectedConfig := selectConfig(*configs)
	option := selectConfigSetting(selectedConfig)
	fmt.Printf(option)
}

type config struct {
	Name    string `yaml:"name"`
	Kind    string `yaml:"kind"`
	Configs []struct {
		Key string   `yaml:"key"`
		Values   []string `yaml:"values"`
	} `yaml:"configs"`
}

func getConfigs() *[]config {
	// get path of current users home directory for reading config file
	usr, err := user.Current()
	check(err, "Current User")

	// read in configs from yaml
	yamlFile, err := ioutil.ReadFile(usr.HomeDir + "/.configctl/config.yaml")
	check(err, "Read File")

	// split up multiple yamls 
	yamls := bytes.Split(yamlFile, []byte("---"))

	var configs []config
	for _, y := range yamls {
		var c config
		err = yaml.Unmarshal(y, &c)
		check(err, "Unmarshal")

		configs = append(configs, c)
	}

	return &configs
}

func selectConfig(configs []config) *config{
	fmt.Println("Select a config type:")
	for i, c := range configs {
		fmt.Printf("%d.) %s\n", i+1, c.Name)
	}
	var selected int
	fmt.Printf("> ")
	fmt.Scanln(&selected)
	fmt.Printf("\n")

	if selected < 0 || selected > len(configs) {
		fmt.Println("ERROR: Invalid selection made.")
		return selectConfig(configs)
	}
	return &configs[selected-1]
}

func selectConfigSetting(c *config) string {
	fmt.Printf("Select %s config:\n", c.Name)
	for i, s := range c.Configs {
		fmt.Printf("%d.) %s\n", i+1, s.Key)
	}
	var selected int
	fmt.Printf("> ")
	fmt.Scanln(&selected)
	fmt.Printf("\n")
	if selected < 0 || selected > len(c.Configs) {
		fmt.Println("ERROR: Invalid selection made.")
		return selectConfigSetting(c)
	}
	return c.Configs[selected-1].Key

}

func check(e error, m string) {
	if e != nil {
		log.Printf(m + ": %v", e)
	}
}
