package convert

import (
	"testing"
)

func assert(result, should interface{}, errmsg string, t *testing.T) {
	if result != should {
		t.Error(errmsg)
		t.Error("result: ", result, " should: ", should)
		t.Fail()
	}
}

/*
* convert to string
 */

func TestConvertStringToString(t *testing.T) {
	var in string = "123456"
	result, _ := String(in)
	assert(result, "123456", "convert string to string error", t)
}

func TestConvertIntToString(t *testing.T) {
	var in int = 10086
	result, _ := String(in)
	assert(result, "10086", "convert int to string error", t)
}

func TestConvertInt32ToString(t *testing.T) {
	var in int32 = 12345
	result, _ := String(in)
	assert(result, "12345", "convert int32 to string error", t)
}

func TestConvertInt64ToString(t *testing.T) {
	var in int64 = 1258096
	result, _ := String(in)
	assert(result, "1258096", "convert int64 to string error", t)
}

func TestConvertFloat32ToString(t *testing.T) {
	var in float32 = 123.321
	result, _ := String(in, 3)
	assert(result, "123.321", "convert float32 to string error", t)
}

func TestConvertFloat64ToString(t *testing.T) {
	var in float64 = 1234567.7654321
	result, _ := String(in, 7)
	assert(result, "1234567.7654321", "convert float64 to string error", t)
}

/*
* convert to int
 */

func TestConvertStringToInt(t *testing.T) {
	var in string = "123"
	result, _ := Int(in)
	assert(result, int(123), "convert string to int error", t)
}

func TestConvertIntToInt(t *testing.T) {
	var in int = 10086
	result, _ := Int(in)
	assert(result, int(in), "convert int to int error", t)
}

func TestConvertInt32ToInt(t *testing.T) {
	var in int32 = 12345
	result, _ := Int(in)
	assert(result, int(in), "convert int32 to int error", t)
}

func TestConvertInt64ToInt(t *testing.T) {
	var in int64 = 1258096
	result, _ := Int(in)
	assert(result, int(in), "convert int64 to int error", t)
}

func TestConvertFloat32ToInt(t *testing.T) {
	var in float32 = 123.321
	result, _ := Int(in)
	assert(result, int(123), "convert int to int error", t)
}

func TestConvertFloat64ToInt(t *testing.T) {
	var in float64 = 1234567.7654321
	result, _ := Int(in)
	assert(result, int(1234567), "convert float64 to int error", t)
}

/*
* convert to int32
 */

func TestConvertStringToInt32(t *testing.T) {
	var in string = "123"
	result, _ := Int32(in)
	assert(result, int32(123), "convert string to int32 error", t)
}

func TestConvertIntToInt32(t *testing.T) {
	var in int = 10086
	result, _ := Int32(in)
	assert(result, int32(in), "convert int to int32 error", t)
}

func TestConvertInt32ToInt32(t *testing.T) {
	var in int32 = 12345
	result, _ := Int32(in)
	assert(result, int32(in), "convert int32 to int32 error", t)
}

func TestConvertInt64ToInt32(t *testing.T) {
	var in int64 = 1258096
	result, _ := Int32(in)
	assert(result, int32(in), "convert int64 to int32 error", t)
}

func TestConvertFloat32ToInt32(t *testing.T) {
	var in float32 = 123.321
	result, _ := Int32(in)
	assert(result, int32(123), "convert int to int32 error", t)
}

func TestConvertFloat64ToInt32(t *testing.T) {
	var in float64 = 1234567.7654321
	result, _ := Int32(in)
	assert(result, int32(1234567), "convert float64 to int32 error", t)
}

/*
* convert to int64
 */

func TestConvertStringToInt64(t *testing.T) {
	var in string = "123"
	result, _ := Int64(in)
	assert(result, int64(123), "convert string to int64 error", t)
}

func TestConvertIntToInt64(t *testing.T) {
	var in int = 10086
	result, _ := Int64(in)
	assert(result, int64(in), "convert int to int64 error", t)
}

func TestConvertInt32ToInt64(t *testing.T) {
	var in int32 = 12345
	result, _ := Int64(in)
	assert(result, int64(in), "convert int32 to int64 error", t)
}

func TestConvertInt64ToInt64(t *testing.T) {
	var in int64 = 1258096
	result, _ := Int64(in)
	assert(result, int64(in), "convert int64 to int64 error", t)
}

func TestConvertFloat32ToInt64(t *testing.T) {
	var in float32 = 123.321
	result, _ := Int64(in)
	assert(result, int64(123), "convert int to int64 error", t)
}

func TestConvertFloat64ToInt64(t *testing.T) {
	var in float64 = 1234567.7654321
	result, _ := Int64(in)
	assert(result, int64(1234567), "convert float64 to int64 error", t)
}

/*
* convert to float32
 */

func TestConvertStringToFloat32(t *testing.T) {
	var in string = "123.456"
	result, _ := Float32(in)
	assert(result, float32(123.456), "convert string to float32 error", t)
}

func TestConvertIntToFloat32(t *testing.T) {
	var in int = 10086
	result, _ := Float32(in)
	assert(result, float32(in), "convert int to float32 error", t)
}

func TestConvertInt32ToFloat32(t *testing.T) {
	var in int32 = 12345
	result, _ := Float32(in)
	assert(result, float32(in), "convert int32 to float32 error", t)
}

func TestConvertInt64ToFloat32(t *testing.T) {
	var in int64 = 1258096
	result, _ := Float32(in)
	assert(result, float32(in), "convert int64 to float32 error", t)
}

func TestConvertFloat32ToFloat32(t *testing.T) {
	var in float32 = 123.321
	result, _ := Float32(in)
	assert(result, float32(in), "convert float32 to float32 error", t)
}

func TestConvertFloat64ToFloat32(t *testing.T) {
	var in float64 = 1234567.7654321
	result, _ := Float32(in)
	assert(result, float32(in), "convert float64 to float32 error", t)
}

/*
* convert to float64
 */

func TestConvertStringToFloat64(t *testing.T) {
	var in string = "123.456"
	result, _ := Float64(in)
	assert(result, float64(123.456), "convert string to float64 error", t)
}

func TestConvertIntToFloat64(t *testing.T) {
	var in int = 10086
	result, _ := Float64(in)
	assert(result, float64(in), "convert int to float64 error", t)
}

func TestConvertInt32ToFloat64(t *testing.T) {
	var in int32 = 12345
	result, _ := Float64(in)
	assert(result, float64(in), "convert int32 to float64 error", t)
}

func TestConvertInt64ToFloat64(t *testing.T) {
	var in int64 = 1258096
	result, _ := Float64(in)
	assert(result, float64(in), "convert int64 to float64 error", t)
}

func TestConvertFloat32ToFloat64(t *testing.T) {
	var in float32 = 123.321
	result, _ := Float64(in)
	assert(result, float64(in), "convert float32 to float64 error", t)
}

func TestConvertFloat64ToFloat64(t *testing.T) {
	var in float64 = 1234567.7654321
	result, _ := Float64(in)
	assert(result, float64(in), "convert float64 to float64 error", t)
}
