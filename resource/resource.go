package resource

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"net/http"
	"time"
)

func loadResource(resource *Resource) {
	Commands = resource.Commands
	Tags = map[string][]Example{}
	for _, command := range Commands {
		if command.Tag != nil {
			for _, tag := range command.Tag {
				for _, exam := range command.Examples {
					if Tags[tag] == nil {
						Tags[tag] = []Example{exam}
					} else {
						_ = append(Tags[tag], exam)
					}
				}
			}
		}
	}
}

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
	loadResource(resource)
	go sync()
}

// 每分钟从 github 同步一次 最新数据
func sync() {
	url := "https://raw.githubusercontent.com/nailcui/exman/master/config/resource.yml"
	ticker := time.NewTicker(10 * time.Minute)
	for {
		<- ticker.C
		fmt.Println("sync start")
		resp, err := http.Get(url)
		if err != nil {
			fmt.Printf("get error: %e url: %s\n", err, url)
			continue
		}
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Printf("read from response body error: %e status: %d\n", err, resp.StatusCode)
			continue
		}
		resource := new(Resource)
		err = yaml.Unmarshal(body, resource)
		if err != nil {
			fmt.Printf("yaml.Unmarshal error: %e\n", err)
			continue
		}
		loadResource(resource)
	}
}

var (
	Commands map[string]command
	Tags     map[string][]Example
)

type Resource struct {
	Commands map[string]command `yaml:"commands"`
}

type command struct {
	Tag []string	`yaml:"tag"`
	Examples []Example	`yaml:"examples"`
}

type Example struct {
	Cmd string	`yaml:"cmd"`
	Desc string	`yaml:"desc"`
}
