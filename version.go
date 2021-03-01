package wen

import (
	"fmt"
	"strconv"
	"strings"
)

func VersionParse(version string, width int) int64 {
	strs := strings.Split(version, ".")
	format := fmt.Sprintf("%%s%%0%ds", width)
	v := ""
	for _, value := range strs {
		v = fmt.Sprintf(format, v, value)
	}
	var result int64
	result, _ = strconv.ParseInt(v, 10, 64)
	return result
}
