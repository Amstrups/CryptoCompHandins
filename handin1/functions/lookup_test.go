package functions

import (
	"handin1/types"
	"testing"
)

var allTypes = []types.TYPES{
	types.O_MINUS,
	types.O_PLUS,
	types.B_MINUS,
	types.B_PLUS,
	types.A_MINUS,
	types.A_PLUS,
	types.AB_MINUS,
	types.AB_PLUS,
}

type dictEntry map[types.TYPES]bool
type ExpectationDict map[types.TYPES]dictEntry

var expectations ExpectationDict = ExpectationDict{
	types.O_MINUS: dictEntry{
		types.O_MINUS:  true,
		types.O_PLUS:   false,
		types.B_MINUS:  false,
		types.B_PLUS:   false,
		types.A_MINUS:  false,
		types.A_PLUS:   false,
		types.AB_MINUS: false,
		types.AB_PLUS:  false,
	},
	types.O_PLUS: dictEntry{
		types.O_MINUS:  true,
		types.O_PLUS:   true,
		types.B_MINUS:  false,
		types.B_PLUS:   false,
		types.A_MINUS:  false,
		types.A_PLUS:   false,
		types.AB_MINUS: false,
		types.AB_PLUS:  false,
	},
	types.B_MINUS: dictEntry{
		types.O_MINUS:  true,
		types.O_PLUS:   false,
		types.B_MINUS:  true,
		types.B_PLUS:   false,
		types.A_MINUS:  false,
		types.A_PLUS:   false,
		types.AB_MINUS: false,
		types.AB_PLUS:  false,
	},
	types.B_PLUS: dictEntry{
		types.O_MINUS:  true,
		types.O_PLUS:   true,
		types.B_MINUS:  true,
		types.B_PLUS:   true,
		types.A_MINUS:  false,
		types.A_PLUS:   false,
		types.AB_MINUS: false,
		types.AB_PLUS:  false,
	},
	types.A_MINUS: dictEntry{
		types.O_MINUS:  true,
		types.O_PLUS:   false,
		types.B_MINUS:  false,
		types.B_PLUS:   false,
		types.A_MINUS:  true,
		types.A_PLUS:   false,
		types.AB_MINUS: false,
		types.AB_PLUS:  false,
	},
	types.A_PLUS: dictEntry{
		types.O_MINUS:  true,
		types.O_PLUS:   true,
		types.B_MINUS:  false,
		types.B_PLUS:   false,
		types.A_MINUS:  true,
		types.A_PLUS:   true,
		types.AB_MINUS: false,
		types.AB_PLUS:  false,
	},
	types.AB_MINUS: dictEntry{
		types.O_MINUS:  true,
		types.O_PLUS:   false,
		types.B_MINUS:  true,
		types.B_PLUS:   false,
		types.A_MINUS:  true,
		types.A_PLUS:   false,
		types.AB_MINUS: true,
		types.AB_PLUS:  false,
	},
	types.AB_PLUS: dictEntry{
		types.O_MINUS:  true,
		types.O_PLUS:   true,
		types.B_MINUS:  true,
		types.B_PLUS:   true,
		types.A_MINUS:  true,
		types.A_PLUS:   true,
		types.AB_MINUS: true,
		types.AB_PLUS:  true,
	},
}

func assert(t *testing.T, reciever types.TYPES, giver types.TYPES, expectation bool) {
	result := Lookup(reciever, giver)
	if result != expectation {
		t.FailNow()
	}
}

func TestAllLookup(t *testing.T) {
	for _, receiver := range allTypes {
		exp := expectations[receiver]
		for _, donor := range allTypes {
			assert(t, receiver, donor, exp[donor])
		}
	}
}
