//ref: https://qiita.com/ryskiwt/items/17617d4f3e8dde7c2b8e

package main

import (
	"crypto/sha256"
	"encoding/hex"
	//"crypto/sha512"
	"fmt"
	//"golang.org/x/crypto/ripemd160"
	//"golang.org/x/crypto/sha3"
	//"io"
	"strconv"
)

func main() {
	//hoge := []int{1, 2, 3, 4, 5, 6}

	var arr []string

	/*for i := 0; i < 10; i++ {
		str := strconv.Itoa(i)
		strSha256 := hash(i, str)

		checkDup(arr, strSha256)
		arr = append(arr, strSha256)


	}*/

	i := 9223372036854775807
	str := strconv.Itoa(i)
	strSha256 := hash(i, str)

	/*if i == 3 {
		arr = append(arr, "6b86b273ff34fce19d6b804eff5a3f5747ada4eaa22f1d49c01e52ddb7875b4b")
	} else {
		arr = append(arr, strSha256)
	}*/
	checkDup(arr, strSha256)
	arr = append(arr, strSha256)

	/*for i, _ := range hoge {
		hash(i)
	}*/
}

func hash(id int, str string) string {

	// sha256
	strSha256 := sha256.Sum256([]byte(str))
	fmt.Printf("Id: %d\n", id)
	fmt.Printf("SHA-256 : %x\n", strSha256)

	// sha512
	/*strSha512 := sha512.Sum512([]byte(str))
	fmt.Printf("Id: %d\n", id)
	fmt.Printf("SHA-512 : %x\n", strSha512)*/

	// sha3-256
	/*fmt.Printf("SHA-3 256 : %x\n", sha3.Sum256([]byte(str)))

	// RIPEMD-160
	rip := ripemd160.New()
	io.WriteString(rip, str)
	fmt.Printf("RIPEMD160 : %x\n", rip.Sum(nil))*/

	fmt.Println("")

	return hex.EncodeToString(strSha256[:])
	//return hex.EncodeToString(strSha512[:])
}

func checkDup(arr []string, str string) {
	for _, s := range arr {
		if str == s {
			fmt.Printf("Duplicated!!! %s", str)
			fmt.Println("")
		}
	}
}
