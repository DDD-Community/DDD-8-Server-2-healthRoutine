package dbx

import (
	"strconv"
)

func ConvertInterfaceToInt64(t interface{}) int64 {
	switch t := t.(type) {
	case int64:
		return t
	case int:
		return int64(t)
	case string:
		src, _ := strconv.ParseInt(t, 10, 64)
		return src
	case []byte:
		src, _ := strconv.ParseInt(string(t), 10, 64)
		return src
	default:
		return 0
	}
}
