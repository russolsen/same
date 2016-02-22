package same

import (
	"testing"
	"reflect"
)

func VerifySame(t *testing.T, v1, v2 interface{}) {
	if ! IsSame(v1, v2) {
		t.Errorf("Expected %v(%v) and %v(%v) to be the same, but they were not.", 
			v1, reflect.TypeOf(v1), v2, reflect.TypeOf(v2))
	}
}
func VerifyDifferent(t *testing.T, v1, v2 interface{}) {
	if IsSame(v1, v2) {
		t.Errorf("Expected %v(%v) and %v(%v) to be different, but they were not.", 
			v1, reflect.TypeOf(v1), v2, reflect.TypeOf(v2))
	}
}

type XString string

func TestSameValueDifferentType(t *testing.T) {
	var a string
	var b interface{}

	a = "russ"
	b = "russ"

	VerifySame(t, a, b)
}

func TestSimpleValues(t *testing.T){
	VerifySame(t, true, true)
	VerifySame(t, false, false)
	VerifyDifferent(t, true, false)

	VerifySame(t, 1, 1)
	VerifySame(t, 1, int8(1))
	VerifySame(t, 1, int16(1))
	VerifySame(t, 1, int32(1))
	VerifySame(t, 1, int64(1))
	
	VerifySame(t, int8(1), int8(1))
	VerifySame(t, int16(1), int16(1))
	VerifySame(t, int32(1), int32(1))
	VerifySame(t, int64(1), int64(1))
	
	VerifySame(t, int8(1), 1)
	VerifySame(t, int16(1), 1)
	VerifySame(t, int32(1), 1)
	VerifySame(t, int64(1), 1)

	VerifySame(t, uint64(1), int8(1))
	VerifySame(t, uint8(1), int64(1))

	VerifySame(t, uint64(1), int16(1))
	VerifySame(t, uint16(16), int64(16))

	VerifySame(t, int64(1), int32(1))
	VerifySame(t, int32(1), int64(1))

	VerifySame(t, 1.23, 1.23)

	VerifySame(t, float32(1.23), 1.23)
	VerifySame(t, 1.23, float32(1.23))

	VerifySame(t, float64(1.23), 1.23)
	VerifySame(t, 1.23, float64(1.23))

	VerifySame(t, float32(1.23), float64(1.23))

	VerifySame(t, "", "")
	VerifySame(t, "a", "a")
	VerifySame(t, "abc", "abc")

	VerifyDifferent(t, "x", "")
	VerifyDifferent(t, "b", "a")
	VerifyDifferent(t, "xbc", "abc")

	VerifySame(t, true, true)
	VerifySame(t, false, false)
}

func TestArrays(t *testing.T) {
	intArray := []int{1,2,3}
	int8Array := []int8{1,2,3}
	int16Array := []int16{1,2,3}
	int32Array := []int32{1,2,3}
	int64Array := []int64{1,2,3}

	VerifySame(t, intArray, int8Array)
	VerifySame(t, intArray, int16Array)
	VerifySame(t, intArray, int32Array)
	VerifySame(t, intArray, int64Array)

	VerifySame(t, int16Array, int8Array)
	VerifySame(t, int16Array, int16Array)
	VerifySame(t, int16Array, int32Array)
	VerifySame(t, int16Array, int64Array)

	VerifySame(t, int32Array, int8Array)
	VerifySame(t, int32Array, int16Array)
	VerifySame(t, int32Array, int32Array)
	VerifySame(t, int32Array, int64Array)

	VerifySame(t, int64Array, int8Array)
	VerifySame(t, int64Array, int16Array)
	VerifySame(t, int64Array, int32Array)
	VerifySame(t, int64Array, int64Array)
}

func TestMaps(t *testing.T) {
	ssName := map[string]string{"fname": "Russ", "lname": "Olsen"}
	iiName := map[interface{}]interface{}{"fname": "Russ", "lname": "Olsen"}

	ssBook := map[string]string{"title": "Jaws", "Author": "Benchley"}

	VerifySame(t, ssName, ssName)
	VerifySame(t, iiName, iiName)

	VerifyDifferent(t, ssName, ssBook)
	VerifyDifferent(t, iiName, ssBook)
}
