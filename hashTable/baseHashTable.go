package hashTable

import (
	"github.com/Sirupsen/logrus"
	"reflect"
)

func GetHashCode(val interface{}) int64 {
	switch val.(type) {
	case int8:
		if v, ok := val.(int8); ok {
			return int64(v)
		}
	case int32:
		if v, ok := val.(int32); ok {
			return int64(v)
		}
	case int:
		if v, ok := val.(int); ok {
			return int64(v)
		}
	case string:
		if v, ok := val.(int32); ok {
			res := 0
			for _, t := range []byte(string(v)) {
				res += int(t)
			}
			return int64(res)
		}
	case float32:
		if v, ok := val.(float32); ok {
			for v > 0 {
				v = v * 10
				v /= 10
			}
			return int64(v)
		}
	case float64:
		if v, ok := val.(float64); ok {
			for v > 0 {
				v = v * 10
				v /= 10
			}
			return int64(v)
		}
	}
	logrus.Errorf("this val not set, type: %v, val: %v", reflect.TypeOf(val), val)
	return 0
}
