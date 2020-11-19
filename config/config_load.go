package config

import (
	"io"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

// LoadConfiguration takes a reader to parse and load the configuration
// for the ACME process to run with.
func LoadConfiguration(configReader io.Reader) (*Configuration, error) {

	b, err := ioutil.ReadAll(configReader)
	if err != nil {
		return nil, err
	}

	t := Configuration{}

	err = yaml.Unmarshal([]byte(b), &t)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	return &t, nil
}
