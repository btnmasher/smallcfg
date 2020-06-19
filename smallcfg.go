package smallcfg

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

//Load reads a json file from the given file path and loads it into
//memory given the interface.
func Load(path string, config interface{}) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return fmt.Errorf("Config file not found: %s", path)
	}

	file, err := ioutil.ReadFile(path)
	if err != nil {
		return fmt.Errorf("Unable to read config file: %s", err)
	}

	if err := json.Unmarshal(file, config); err != nil {
		return fmt.Errorf("Error parsing config file: %s", err)
	}
	return nil
}

//Save accepts a configuration interface type and writes it to a file
//in json format.
//Caller can specify whether or not to write the json "pretty" printed
//using indentation by specifying true or false to the pretty parameter.
func Save(path string, config interface{}, pretty bool) error {
	var bytes []byte
	var err error

	if pretty {
		bytes, err = json.MarshalIndent(config, "", "    ")
	} else {
		bytes, err = json.Marshal(config)
	}

	if err != nil {
		return fmt.Errorf("Unable to marshal configuration to json: %s", err)
	}

	err = ioutil.WriteFile(path, bytes, 0644)
	if err != nil {
		return fmt.Errorf("Unable to write config file: %s", err)
	}
	return nil
}
