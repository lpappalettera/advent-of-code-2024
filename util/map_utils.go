package util

import "strconv"

func StrToInt(s string) int {
	val, err := strconv.Atoi(s)
	HandleError(err)
	return val
}
