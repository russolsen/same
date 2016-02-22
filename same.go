package same

import (
	"log"
	"math"
	"reflect"
)

func isIntKind(k reflect.Kind) bool {
	switch k {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return true
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return true
	default:
		return false
	}
}

func isFloatKind(k reflect.Kind) bool {
	switch k {
	case reflect.Float32, reflect.Float64:
		return true
	default:
		return false
	}
}

// isCompatibleKinds returns true if the two kinds are compatible -- that
// is that values of two compatible kinds might be equal.
func isCompatibleKinds(k1, k2 reflect.Kind) bool {
	if k1 == k2 {
		return true
	} else if k1 == reflect.Slice && k2 == reflect.Array {
		return true
	} else if k2 == reflect.Slice && k1 == reflect.Array {
		return true
	}

	if isIntKind(k1) && isIntKind(k2) {
		return true
	}

	if isFloatKind(k1) && isFloatKind(k2) {
		return true
	}

	return false
}

// isSameArray returns true if the two array values are the same length
// and have the same contents as defined by IsSame. This function explicitly
// ignores the declared type of the array.
func isSameArray(x, y reflect.Value) bool {
	lx := x.Len()

	if lx != y.Len() {
		return false
	}

	for i := 0; i < lx; i++ {
		xElement := x.Index(i)
		yElement := y.Index(i)
		if !IsSame(xElement.Interface(), yElement.Interface()) {
			return false
		}
	}
	return true
}

type StringConvertable interface {
	String() string
}

// isSameUnknownStruct returns true if the two struct values are the same length
// and have the same contents as defined by IsSame. This function explicitly
// ignores the declared type of the struct.
func isSameStruct(x, y reflect.Value) bool {

	sx, okx := x.Interface().(StringConvertable)
	sy, oky := x.Interface().(StringConvertable)

	if okx && oky {
		//log.Println("Stringable", x, y)
		return sx.String() == sy.String()
	}

	numFields := x.NumField()

	if numFields != y.NumField() {
		return false
	}

	typeX := x.Type()

	for i := 0; i < numFields; i++ {
		path := []int{i}
		vx := x.FieldByIndex(path)
		vy := y.FieldByName(typeX.Field(i).Name)
		if vx.CanInterface() && vy.CanInterface() {
			if !IsSame(vx.Interface(), vy.Interface()) {
				return false
			}
		}
	}

	return true
}

// isSameMap returns true if the two map values are the same length
// and have the same contents as defined by IsSame. This function explicitly
// ignores the declared type of the map.
func isSameMap(x, y reflect.Value) bool {
	//log.Println("\n\nSame map: x:",x, x.Type(),  " y:", y, y.Type())
	lx := x.Len()

	if lx != y.Len() {
		return false
	}

	// Avoid the typing problems with Go maps by turning the
	// two maps into DumbMaps and doing the comparison there.

	dmX := NewDumbMap(x.Interface())
	//log.Println("Len of map:", dmX, dmX.Count())
	dmY := NewDumbMap(y.Interface())

	return dmX.IsSame(dmY)
}

// IsSame defines a fairly lose idea of structual equality. Two values
// are the same if they are integers and their int64 values are equal.
// Or if they are floats and their float64 values are equal.
// Or if they are both true or false.
// Or if they are both the same string.
// Or if they are maps that contain the same keys and values.
// Or if they are arrays that contain the same number and contents.
// Or if they are structs with the same public fields. TBD
// Note that IsSame explicitly ignores the declared types of the values
// except as noted above.
func IsSame(x, y interface{}) bool {
	if x == nil && y == nil {
		return true
	}
	return IsSameValue(reflect.ValueOf(x), reflect.ValueOf(y))
}

func unwindPointer(x reflect.Value) reflect.Value {
	if x.Kind() != reflect.Ptr {
		return x
	}

	for x.Kind() == reflect.Ptr {
		x = x.Elem()
	}

	return reflect.ValueOf(x.Interface())
}

func getInt64(x reflect.Value) int64 {
	switch x.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return x.Int()

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return int64(x.Uint())
	default:
		panic("Not an int")
	}
}

const Epsilon = 1.0E-5

// IsSameThing compares two non-pointer values.
func IsSameValue(x, y reflect.Value) bool {

	x = unwindPointer(x)
	y = unwindPointer(y)

	kindX := x.Kind()
	kindY := y.Kind()

	if !isCompatibleKinds(kindX, kindY) {
		return false
	}

	switch kindX {

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return getInt64(x) == getInt64(y)

	case reflect.Float32, reflect.Float64:
		//log.Println("***Same float!!", x, y)
		xf := x.Float()
		yf := y.Float()
		if xf == yf {
			return true
		} else {
			return math.Abs(xf-yf) < Epsilon
		}

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return getInt64(x) == getInt64(y)

	case reflect.String:
		return x.String() == y.String()

	case reflect.Bool:
		return x.Bool() == y.Bool()

	case reflect.Array, reflect.Slice:
		return isSameArray(x, y)

	case reflect.Map:
		return isSameMap(x, y)

	case reflect.Struct:
		return isSameStruct(x, y)

	default:
		log.Println("IsSame: **** What is:", x, kindX)
		return false

	}
}
