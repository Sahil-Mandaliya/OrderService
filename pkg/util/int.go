package util

import (
	"log"
	"strconv"
)

func ParseUint64(numStr string) uint64 {
	num, err := strconv.ParseUint(numStr, 10, 64)
	if err != nil {
		log.Println("Error Parsing numStr :", err.Error(), "due to: ", err.Error())
		return 0
	}
	return num
}
