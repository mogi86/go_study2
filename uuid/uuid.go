package main

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"github.com/google/uuid"
)

func main() {
	//uuidStr := uuid.New().String()
	oldId := "9999999998"
	generatedUUID := newUUID2(oldId)
	fmt.Println("generated uuid: " + generatedUUID)

	oldId2 := "9999999999"
	generatedUUID2 := newUUID2(oldId2)
	fmt.Println("generated uuid: " + generatedUUID2)
}

func newUUID(oldId string) string {
	/*b := []byte(oldId)
	hashed := md5.Sum(b)
	fmt.Println("hashed value by md5: " + hex.EncodeToString(hashed[:]))
	return fmt.Sprintf("%x-%x-%x-%x-%x", hashed[0:4], hashed[4:6], hashed[6:8], hashed[8:10], hashed[10:])*/

	b := []byte(oldId)
	hashed := md5.Sum(b)
	fmt.Println("hashed value by md5: " + hex.EncodeToString(hashed[:]))
	fmt.Println(uuid.MustParse(hex.EncodeToString(hashed[:])))
	/*
		// variant bits; see section 4.1.1 0xc0=192, 0x80=128
		hashed[8] = hashed[8]&^0xc0 | 0x80
		// version 4 (pseudo-random); see section 4.1.3 0xf0=240, 64
		hashed[6] = hashed[6]&^0xf0 | 0x40
	*/
	//ref: https://play.golang.org/p/4FkNSiUDMg
	//ref: https://github.com/google/uuid/blob/master/version4.go#L39-L40
	hashed[6] = (hashed[6] & 0x0f) | 0x40 // Version 4
	hashed[8] = (hashed[8] & 0x3f) | 0x80 // Variant is 10
	return fmt.Sprintf("%x-%x-%x-%x-%x", hashed[0:4], hashed[4:6], hashed[6:8], hashed[8:10], hashed[10:])
}

func newUUID2(oldId string) string {
	data := []byte(oldId)
	NamespaceUUID := uuid.MustParse("3409f30e-6415-4b90-8dd1-5b813343f8f9")
	return uuid.NewHash(sha1.New(), NamespaceUUID, data, 5).String()
}
