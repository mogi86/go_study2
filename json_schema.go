package main

import (
	"flag"
	"fmt"
	"github.com/xeipuuv/gojsonschema"
)

func main() {
	flag.Parse()
	args := flag.Args()

	jsonPath := args[0]
	dataPath := args[1]

	schemaLoader := gojsonschema.NewReferenceLoader(fmt.Sprintf("file://%v", jsonPath))
	documentLoader := gojsonschema.NewReferenceLoader(fmt.Sprintf("file://%v", dataPath))

	result, err := gojsonschema.Validate(schemaLoader, documentLoader)
	if err != nil {
		panic(err.Error())
	}

	if result.Valid() {
		fmt.Printf("The document is valid\n")
	} else {
		fmt.Printf("The document is not valid. see errors :\n")
		for _, desc := range result.Errors() {
			fmt.Printf("- %s\n", desc)
		}
	}
}
