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

	if isJson {
		y, err := yaml.JSONToYAML(data)
		if err != nil {
			log.Fatalf("err: %v\n", err)
		}
		fmt.Print(string(y))
	} else {
		j2, err := yaml.YAMLToJSON(data)
		if err != nil {
			log.Fatalf("err: %v\n", err)
		}
		fmt.Print(string(j2))
	}
}
