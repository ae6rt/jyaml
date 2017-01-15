package main

//
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

	var result []byte

	// try to guess the output format based on the input data
	isJson := json.Unmarshal(data, new(interface{})) == nil

	if isJson {
		result, err = yaml.JSONToYAML(data)
	} else {
		result, err = yaml.YAMLToJSON(data)
	}

	if err != nil {
		log.Fatalf("err: %v\n", err)
	}

	fmt.Print(string(result))
}
