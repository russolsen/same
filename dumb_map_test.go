package same

import (
	"testing"
)


func TestDumbMap(t *testing.T) {
	sm := map[string]string{"fname":"Russ", "lname":"Olsen"}
	im := map[interface{}]interface{}{"fname":"Russ", "lname":"Olsen"}
	addr := map[interface{}]interface{}{"number": 3454, "Street":"Ainslie"}

	VerifySame(t, sm, sm)
	VerifySame(t, im, im)
	VerifySame(t, sm, im)

	VerifyDifferent(t, sm, addr)
}
