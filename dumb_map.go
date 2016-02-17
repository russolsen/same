package same

import (
	"reflect"
	//"log"
)

type DumbMap struct {
	Keys []interface{}
	Values []interface{}
}

func NewDumbMap(sourceMap interface{})*DumbMap {
	vSourceMap := reflect.ValueOf(sourceMap)
	sourceKeys := vSourceMap.MapKeys()

	keys := []interface{}{}
	values := []interface{}{}

	for _, key := range sourceKeys {
		value := vSourceMap.MapIndex(key)
		keys = append(keys, key.Interface())
		values = append(values, value.Interface())
	}

	result := DumbMap{keys, values}
	return &result
}

func (m DumbMap)Count()int {
	return len(m.Keys)
}

func (m DumbMap)Index(key interface{})(interface{}) {
	for i, thisKey := range m.Keys {
		if IsSame(key, thisKey) {
			return m.Values[i]
		}
	}
	return nil
}

func (m DumbMap)IsSame(other *DumbMap) bool {

	if m.Count() != other.Count() {
		return false
	}

	for _, key := range m.Keys {
		v1 := m.Index(key)
		v2 := other.Index(key)

		//log.Println("***DM Is same:", key, v1, v2)

		if ! IsSame(v1, v2) {
			return false
		}
	}
	return true
}
