package convert

import (
	"encoding/json"
)

func JSONmap(in []byte) (map[string]interface{}, error) {
	var i interface{}
	e := json.Unmarshal(in, &i)
	if e != nil {
		return nil, e
	}
	m := i.(map[string]interface{})
	return m, nil
}
