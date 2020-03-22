package main

import (
	"fmt"
	gouuid "github.com/satori/go.uuid"
)

func main() {
	// Creating UUID Version 4
	// panic on error
	u1 := gouuid.Must(gouuid.NewV4(), nil)
	fmt.Printf("UUIDv4: %s\n", u1)

	// or error handling
	u2 := gouuid.NewV4()
	fmt.Printf("UUIDv4: %s\n", u2)

	// Parsing UUID from string input
	u2, err := gouuid.FromString("6ba7b810-9dad-11d1-80b4-00c04fd430c8")
	if err != nil {
		fmt.Printf("Something went wrong: %s", err)
		return
	}
	fmt.Printf("Successfully parsed: %s\n", u2)

	idId := "1"
	u3 := gouuid.NewV5(gouuid.Nil, idId)
	u4 := gouuid.NewV5(gouuid.NamespaceOID, idId)
	fmt.Printf("UUIDv5 NIL: %s\n", u3)
	fmt.Printf("UUIDv5 NIL: %s\n", u4)
}
