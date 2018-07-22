package history

import (
	"io/ioutil"
	"log"
	"gopkg.in/yaml.v2"
)

type TypeConfig struct {
	Baseurl string
	Apiurl  string
	Session string
	Token   string
	Host    string
	Port    string
}

func LoadConfig() *TypeConfig {
	conf := &TypeConfig{}

	data, errReadFile := ioutil.ReadFile("config/approved-history-service.yaml")
	if errReadFile != nil {
		log.Fatalf("error: %v", errReadFile)
	}

	errYaml := yaml.Unmarshal(data, &conf)
	if errYaml != nil {
		log.Fatalf("error: %v", errYaml)
	}

	return conf

}
