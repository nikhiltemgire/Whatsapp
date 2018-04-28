package utils

import (
	"encoding/json"
	"log"
	"strconv"
)

// StructToJSONString : converts structures to json strings
func StructToJSONString(t interface{}) string {
	jsonStr, err := json.Marshal(t)

	if err != nil {
		log.Fatal(err)
	}
	return string(jsonStr)
}

// StringToInteger : Converts string to integer
func StringToInteger(str string) int {
	integer, err := strconv.Atoi(str)
	if err != nil {
		log.Fatal(err)
	}
	return integer
}
