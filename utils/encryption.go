package utils

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

var hasher = md5.New()

//GetMD5Hash : return the md5 hash of given string
func GetMD5Hash(text string) string {

	fmt.Println(">> ", text, "<<")
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}
