package actors

import (
	"math/rand"
)

type Dealer struct {
	vAlice Values
	vBob   Values
}

func setupValues() (Values, Values) {
	V_a := Values{[5]uint8{}, [5]uint8{}, [5]uint8{}}
	V_b := Values{[5]uint8{}, [5]uint8{}, [5]uint8{}}

	i := 0

	for i < 5 {
		ua := uint8(rand.Intn(2))
		va := uint8(rand.Intn(2))
		ub := uint8(rand.Intn(2))
		vb := uint8(rand.Intn(2))
		wb := uint8(rand.Intn(2))

		wa := ((ua ^ ub) & (va ^ vb)) ^ wb

		V_a.Vs[i] = va
		V_a.Us[i] = ua
		V_a.Ws[i] = wa

		V_b.Vs[i] = vb
		V_b.Us[i] = ub
		V_b.Ws[i] = wb

		i++
	}

	return V_a, V_b
}

func DealerInit() *Dealer {
	va, vb := setupValues()
	return &Dealer{vAlice: va, vBob: vb}
}

func (d *Dealer) ToAlice() Values {
	return d.vAlice
}

func (d *Dealer) ToBob() Values {
	return d.vBob
}
