package convert

import (
	"testing"
)

func TestJSONmap(t *testing.T) {
	b := []byte(`{"name": "hello","count": 2,"users": [{"name": "hao","age": 10086},{"name": "xin","age": 12306}]}`)
	m, e := JSONmap(b)
	if e != nil {
		t.Fail()
	}
	if m["name"] != "hello" {
		t.Fail()
	}
}
