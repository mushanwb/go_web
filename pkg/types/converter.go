package types

import "strconv"

func Int64ToString(num int64) string {
	return strconv.FormatInt(num, 10)
}
