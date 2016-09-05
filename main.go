package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/ghodss/yaml"
)

func main() {
	data, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatalf("err: %v\n", err)
	}

	// try to guess the output format based on the input data
	isJson := json.Unmarshal(data, new(interface{})) == nil

	var f func([]byte) ([]byte, error)

	if isJson {
		f = yaml.JSONToYAML
	} else {
		f = yaml.YAMLToJSON
	}

	if result, err := f(data); err != nil {
		log.Fatalf("err: %v\n", err)
	} else {
		fmt.Print(string(result))
	}
}
