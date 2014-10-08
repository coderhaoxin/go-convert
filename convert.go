package convert

import (
	"errors"
	"strconv"
)

// args: value, precision(only for float)
func String(args ...interface{}) (string, error) {
	value := args[0]
	var precision int = 12 // default

	switch value.(type) {
	case string:
		v, _ := value.(string)
		return v, nil
	case int:
		v, _ := value.(int)
		return strconv.Itoa(v), nil
	case int32:
		v, _ := value.(int32)
		return strconv.FormatInt(int64(v), 10), nil
	case int64:
		v, _ := value.(int64)
		return strconv.FormatInt(v, 10), nil
	case float32:
		v, _ := value.(float32)
		if len(args) >= 2 {
			precision = args[1].(int)
		}
		return strconv.FormatFloat(float64(v), 'f', precision, 64), nil
	case float64:
		v, _ := value.(float64)
		if len(args) >= 2 {
			precision = args[1].(int)
		}
		return strconv.FormatFloat(v, 'f', precision, 64), nil
	default:
		return "", errors.New("unknown type")
	}
}

func Int(value interface{}) (int, error) {
	switch value.(type) {
	case string:
		v, _ := value.(string)
		return strconv.Atoi(v)
	case int:
		v, _ := value.(int)
		return v, nil
	case int32:
		v, _ := value.(int32)
		return int(v), nil
	case int64:
		v, _ := value.(int64)
		return int(v), nil
	case float32:
		v, _ := value.(float32)
		return int(v), nil
	case float64:
		v, _ := value.(float64)
		return int(v), nil
	default:
		return int(0), errors.New("unknown type")
	}
}

func Int32(value interface{}) (int32, error) {
	switch value.(type) {
	case string:
		v, _ := value.(string)
		result, err := strconv.ParseInt(v, 10, 32)
		return int32(result), err
	case int:
		v, _ := value.(int)
		return int32(v), nil
	case int32:
		v, _ := value.(int32)
		return int32(v), nil
	case int64:
		v, _ := value.(int64)
		return int32(v), nil
	case float32:
		v, _ := value.(float32)
		return int32(v), nil
	case float64:
		v, _ := value.(float64)
		return int32(v), nil
	default:
		return int32(0), errors.New("unknown type")
	}
}

func Int64(value interface{}) (int64, error) {
	switch value.(type) {
	case string:
		v, _ := value.(string)
		return strconv.ParseInt(v, 10, 32)
	case int:
		v, _ := value.(int)
		return int64(v), nil
	case int32:
		v, _ := value.(int32)
		return int64(v), nil
	case int64:
		v, _ := value.(int64)
		return v, nil
	case float32:
		v, _ := value.(float32)
		return int64(v), nil
	case float64:
		v, _ := value.(float64)
		return int64(v), nil
	default:
		return int64(0), errors.New("unknown type")
	}
}

func Float32(value interface{}) (float32, error) {
	switch value.(type) {
	case string:
		v, _ := value.(string)
		result, err := strconv.ParseFloat(v, 32)
		return float32(result), err
	case int:
		v, _ := value.(int)
		return float32(v), nil
	case int32:
		v, _ := value.(int32)
		return float32(v), nil
	case int64:
		v, _ := value.(int64)
		return float32(v), nil
	case float32:
		v, _ := value.(float32)
		return v, nil
	case float64:
		v, _ := value.(float64)
		return float32(v), nil
	default:
		return float32(0), errors.New("unknown type")
	}
}

func Float64(value interface{}) (float64, error) {
	switch value.(type) {
	case string:
		v, _ := value.(string)
		return strconv.ParseFloat(v, 64)
	case int:
		v, _ := value.(int)
		return float64(v), nil
	case int32:
		v, _ := value.(int32)
		return float64(v), nil
	case int64:
		v, _ := value.(int64)
		return float64(v), nil
	case float32:
		v, _ := value.(float32)
		return float64(v), nil
	case float64:
		v, _ := value.(float64)
		return v, nil
	default:
		return float64(0), errors.New("unknown type")
	}
}
