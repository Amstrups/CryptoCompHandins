package main

import (
	"fmt"
	"handin3/actors"
	"handin3/types"
)

type dictEntry map[types.TYPES]bool
type ExpectationDict map[types.TYPES]dictEntry

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

func main() {
	passed := 0
	failed := 0

	for _, receiver := range allTypes {
		for _, donor := range allTypes {
			d := actors.DealerInit()
			alice := actors.AliceInit(receiver, d) // receiver
			bob := actors.BobInit(donor, d)        // donor

			for !alice.HasOutput {
				bob.Receive(alice.Send())
				alice.Receive(bob.Send())
			}
			result := alice.Output()

			expected := actors.Compatibility(receiver, donor)

			if result == expected {
				passed++
			} else {
				failed++
			}
		}
	}

	fmt.Printf("Passed: %d, failed: %d\n", passed, failed)
}
