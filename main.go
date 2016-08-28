package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/ghodss/yaml"
)

func main() {
	var output = flag.String("o", "", "output format (j=json, y=yaml)")
	var input = flag.String("file", "", "input file")
	flag.Parse()

	data, err := ioutil.ReadFile(*input)
	if err != nil {
		log.Fatalf("err: %v\n", err)
	}

	switch *output {
	case "j":
		j2, err := yaml.YAMLToJSON(data)
		if err != nil {
			log.Fatalf("err: %v\n", err)
		}
		fmt.Println(string(j2))
	case "y":
		y, err := yaml.JSONToYAML(data)
		if err != nil {
			log.Fatalf("err: %v\n", err)
		}
		fmt.Println(string(y))
	default:
		log.Fatalf("Unsupported output format %s\n", *output)
	}

}
