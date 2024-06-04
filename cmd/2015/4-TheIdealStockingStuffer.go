package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"strconv"
)

func MD5Hash(input string) string {
	hasher := md5.New()
	_, err := io.WriteString(hasher, input)

	if err != nil {
		panic(err)
	}

	return hex.EncodeToString(hasher.Sum(nil))
}

func RunFour(){
	input := "iwrupvqb"

	for i := 0; i < 10000000; i++ {
		hash := MD5Hash(input + strconv.Itoa(i))

		if hash[0:6] == "000000" {
			fmt.Println(i)
		}
	}

}