package main

import (
	"encoding/json"
	"github.com/srevinsaju/gofer/types"
	"io/ioutil"
)

/* ConfigFromFile creates */
func ConfigFromFile(filepath string) (types.GoferConfig, error) {
	rawData, err := ioutil.ReadFile(filepath)
	var cfg types.GoferConfig
	err = json.Unmarshal(rawData, &cfg)
	if err != nil {
		logger.Fatal(err)
		return types.GoferConfig{}, err
	}
	return cfg, nil
}
