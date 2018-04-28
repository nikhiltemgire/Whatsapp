package utils

import (
	"crypto/md5"
	"encoding/hex"
)

//GetMD5Hash : return the md5 hash of given string
func GetMD5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}
