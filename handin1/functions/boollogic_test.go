package functions

import (
	"handin1/types"
	"testing"
)

func boolAssert(t *testing.T, reciever types.TYPES, giver types.TYPES, expectation bool) {
	result := Compatibility(reciever, giver)
	bool_result := result == 1
	if bool_result != expectation {
		t.FailNow()
	}
}

func TestAllBool(t *testing.T) {
	for _, receiver := range allTypes {
		exp := expectations[receiver]
		for _, donor := range allTypes {
			boolAssert(t, receiver, donor, exp[donor])
		}
	}
}
