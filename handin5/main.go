package main

import (
	"fmt"
	"handin5/actors"
	"handin5/types"
)

func main() {
	M := [8][8]uint8{}
	for i, receiver := range types.AllTypes {
		for j, donor := range types.AllTypes {

			simulateOT := true
			Alice := actors.Alice{Blood: receiver}
			Bob := actors.MakeBob(donor)

			if simulateOT {
				Alice.X = Bob.SimulateOT(Alice.Blood)
			} else { // Real OT
				for i := 0; i < 3; i++ {
					Alice.Receive(i, Bob.Receive(Alice.Send()))
				}
			}

			FED := Bob.FED()
			z := Alice.Decode(FED)

			M[i][7-j] = z

		}
	}

	for _, m := range M {
		fmt.Println(m)

	}

}
