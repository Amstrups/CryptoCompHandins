package functions

import (
	"handin2/actors"
	"handin2/types"
)

func FullOTTTRun(receiver types.TYPES, donor types.TYPES) bool {
	d := actors.DealerInit(LOOKUP_TABLE)

	alice := actors.AliceInit(receiver, d)

	bob := actors.BobInit(donor, d)

	bob.Receive(alice.Send())
	alice.Receive(bob.Send())

	output := alice.Output()

	return output == 1
}
