package parsers

import (
	"fmt"
	"io/ioutil"

	//"github.com/Serj1c/yaml-to-openmetrics/pkg/currencies"
	"gopkg.in/yaml.v2"
)

// YamlFile represents the structure of a file to be parsed
type YamlFile struct {
	Currencies []struct {
		Name  string  `yaml:"name"`
		Value float64 `yaml:"value"`
	} `yaml:"currencies"`
}

// ParseYaml takes file name of as an input, parses
func ParseYaml(fileName string) (*YamlFile, error) {
	data := &YamlFile{}
	yml, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, fmt.Errorf("cannot read yaml file")
	}
	err = yaml.Unmarshal(yml, data)
	if err != nil {
		return nil, fmt.Errorf("cannot parse yaml file")
	}
	return data, nil
}
