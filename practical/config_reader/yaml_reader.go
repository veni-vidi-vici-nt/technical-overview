package configreader

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// YAMLReader is an implementation of the Reader interface for YAML files
type YAMLReader struct {
	Filename string
}

func (r *YAMLReader) ReadConfig() (map[string]interface{}, error) {
	data, err := ioutil.ReadFile(r.Filename)
	if err != nil {
		return nil, err
	}

	var config map[string]interface{}
	err = yaml.Unmarshal(data, &config)

	return config, err
}
