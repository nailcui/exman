package resource

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

func init() {
	resource := new(Resource)
	file, err := ioutil.ReadFile("config/resource.yml")
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(file, resource)
	if err != nil {
		panic(err)
	}
	Commands = resource.Commands
}

var (
	Commands map[string][]command
)

type Resource struct {
	Commands map[string][]command `yaml:"commands"`
}

type command struct {
	Cmd string	`yaml:"cmd"`
	Desc string	`yaml:"desc"`
}
