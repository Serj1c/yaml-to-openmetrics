package parsers

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// YamlCurrencies represents the structure of a file to be parsed
type YamlCurrencies struct {
	Currencies []struct {
		Name  string  `yaml:"name"`
		Value float64 `yaml:"value"`
	} `yaml:"currencies"`
}

// ParseYaml takes file name of as an input, parses
func ParseYaml(fileName string) (*YamlCurrencies, error) {
	data := &YamlCurrencies{}
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
