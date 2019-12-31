package json_schema

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
	dataLoader := gojsonschema.NewReferenceLoader(fmt.Sprintf("file://%v", dataPath))

	result, err := gojsonschema.Validate(schemaLoader, dataLoader)
	if err != nil {
		panic(err.Error())
	}

	if result.Valid() {
		fmt.Printf("The date is valid\n")
	} else {
		fmt.Printf("The date is not valid. see errors :\n")
		for _, desc := range result.Errors() {
			fmt.Printf("- %s\n", desc)
		}
	}
}
