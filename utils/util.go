package utils

import (
	"log"
	"strconv"
)

func CheckID(id string) (int, bool) {
	if _, err := strconv.Atoi(id); err != nil {
		log.Printf("this has error for this %s", id)
		return 0, false
	}
	value, _ := strconv.Atoi(id)

	if value <= 0 {
		return 0, false
	}

	return value, true
}
