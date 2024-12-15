package actors

import (
	t "handin5/types"
	"handin5/util"
)

type Alice struct {
	Blood t.TYPES
	X     [3]t.K
}

func (a *Alice) Send() (uint8, uint8) {
	return 0, 0
}

func (a *Alice) Receive(round int, m uint8) {}

func (a *Alice) Decode(FED t.FED) uint8 {
	Ks := [5]t.K{}

Cs:
	for i, C_ := range FED.F {
	Csj:
		for j, cj := range C_ {
			Ktau := [32]byte{}
			if i < 3 {
				Ktau = util.XOR32([32]byte(util.G(a.X[i], FED.E[i], i)), [32]byte(cj))
			} else {
				index := int((i-2)/2) * 2

				Ktau = util.XOR32([32]byte(util.G(Ks[index], Ks[index+1], i)), [32]byte(cj))
			}
			res := [16]byte{}
			for v := 0; v < 16; v++ {

				if Ktau[16+v] != 0 {
					if j == 3 {
						panic(0)

					}
					continue Csj
				}
				res[v] = Ktau[v]
			}

			Ks[i] = res
			continue Cs
		}

	}

	if util.EQ(FED.D.K1, Ks[len(Ks)-1]) {
		return 1
	}
	return 0

}
