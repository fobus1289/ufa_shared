package utils

import "strconv"

func Int64ToString(value int64) string {
	return strconv.FormatInt(value, 10)
}

func StringToInt64(value string) (int64, error) {
	return strconv.ParseInt(value, 10, 64)
}
