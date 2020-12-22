package types

import (
	"go_web/pkg/logger"
	"strconv"
)

func Int64ToString(num int64) string {
	return strconv.FormatInt(num, 10)
}

func StringToInt64(str string) int64 {
	i, err := strconv.ParseInt(str, 10, 64)

	if err != nil {
		logger.LogError(err)
	}
	return i
}
