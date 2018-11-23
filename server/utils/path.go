package utils

import (
	"strings"
	"strconv"
)

func PathTail(path string) int {
	pathBits := strings.Split(path, "/")
	number, _ := strconv.Atoi(pathBits[len(pathBits)-1])
	return number
}
