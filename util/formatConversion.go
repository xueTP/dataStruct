package util

import (
	"strconv"
)

func InterfaceToString(val interface{}) string {
	var s string
	switch val.(type) {
	case int, int8, int16, int32, int64:
		s = strconv.Itoa(val.(int))
	case bool:
		s = strconv.FormatBool(val.(bool))
	case string:
		s = val.(string)
	case float32, float64:
		s = strconv.FormatFloat(val.(float64), 'E', -1, 64)
	case byte:
		s = string(val.(byte))
	default:
		s = InterfaceAssertion(val)
	}

	return s
}

func InterfaceAssertion(val interface{}) string {
	if v, ok := val.(int); ok {
		return strconv.Itoa(v)
	}else if v, ok := val.(float64); ok {
		return strconv.FormatFloat(v, 'E', -1, 64)
	}else if v, ok := val.(bool); ok {
		return strconv.FormatBool(v)
	}else if v, ok := val.(string); ok {
		return v
	}else if v, ok := val.(uint64); ok {
		return strconv.FormatUint(v, 10)
	}else if v, ok := val.(byte); ok {
		return string(v)
	}
	return ""
}