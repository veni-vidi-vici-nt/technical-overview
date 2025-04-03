package configreader

import (
	"encoding/json"
	"io/ioutil"
)

// JSONReader is an implementation of the Reader interface for JSON files
type JSONReader struct {
	Filename string
}

func (r *JSONReader) ReadConfig() (map[string]interface{}, error) {
	data, err := ioutil.ReadFile(r.Filename)
	if err != nil {
		return nil, err
	}
	var config map[string]interface{}
	err = json.Unmarshal(data, &config)
	return config, err
}
