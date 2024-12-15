package actors

import (
	t "handin5/types"
	u "handin5/util"
)

type Bob struct {
	Blood t.TYPES
	Ks    [t.T]t.KPair
	Y     [3]t.K
}

func MakeBob(y t.TYPES) *Bob {
	b := &Bob{Blood: y}
	b.Ks = [t.T]t.KPair{}

	for i := 0; i < t.T; i++ { // For i \in [1,T]
		b.Ks[i] = t.KPair{
			K0: t.MakeK(),
			K1: t.MakeK(),
		}
	}

	y1 := uint8(b.Blood & 1)
	y2 := uint8((b.Blood >> 1) & 1)
	y3 := uint8((b.Blood >> 2) & 1)

	b.Y = [3]t.K{
		b.KSXOR(3, y1, 1),
		b.KSXOR(4, y2, 1),
		b.KSXOR(5, y3, 1),
	}

	return b
}

// Simulate OT transfer
func (b *Bob) SimulateOT(x t.TYPES) [3]t.K {

	x1 := uint8(x & 1)
	x2 := uint8((x >> 1) & 1)
	x3 := uint8((x >> 2) & 1)

	X := [3]t.K{
		b.Ks[0].Choose(x1),
		b.Ks[1].Choose(x2),
		b.Ks[2].Choose(x3),
	}

	return X
}

// Real OT transfer
func (b *Bob) Receive(m0 uint8, m1 uint8) uint8 {
	return 0
}

func (bob *Bob) KSAND(index int, a uint8, b uint8) t.K {
	x := a & b
	return bob.Ks[index].Choose(x)
}

func (bob *Bob) KSXOR(index int, a uint8, b uint8) t.K {
	x := (a ^ b) & 1
	return bob.Ks[index].Choose(x)
}

func (bob *Bob) KSNAND(index int, a uint8, b uint8) t.K {
	x := (^(a & b)) & 1
	return bob.Ks[index].Choose(x)
}

func (b *Bob) FED() t.FED {
	Gb := b.Create()
	f := [][][]byte{}
	for _, gb := range Gb {
		f = append(f, gb.Permutated())
	}

	return t.FED{
		F: f,
		E: b.Y,
		D: b.Ks[t.T-1],
	}
}

func (bob *Bob) Create() []t.Gate {
	// z1_1: (y1 ^ 1)
	// z1_2: (y2 ^ 1)
	// z1_3: (y3 ^ 1)
	// z2_1: (x1 & z1_1)
	// z2_2: (x2 & z1_2)
	// z2_3: (x3 & z1_3)
	// z3_1: (1 ^ z2_1)
	// z3_2: (1 ^ z2_2)
	// z3_3: (1 ^ z2_3)
	// z4_1: (z3_1 & z3_2)
	// z5_1: (z4_1 & z3_3)

	gates := make([]t.Gate, 5)

	for i := 0; i < 3; i++ {
		Gi := t.Gate{}
		Gi.C1 = u.XOR(u.G(bob.Ks[i].K0, bob.Ks[i+3].K0, i), bob.KSNAND(i+6, 0, 0))
		Gi.C2 = u.XOR(u.G(bob.Ks[i].K1, bob.Ks[i+3].K0, i), bob.KSNAND(i+6, 1, 0))
		Gi.C3 = u.XOR(u.G(bob.Ks[i].K0, bob.Ks[i+3].K1, i), bob.KSNAND(i+6, 0, 1))
		Gi.C4 = u.XOR(u.G(bob.Ks[i].K1, bob.Ks[i+3].K1, i), bob.KSNAND(i+6, 1, 1))

		gates[i] = Gi
	}

	// (z3_1 & z3_2)
	G3 := t.Gate{}
	G3.C1 = u.XOR(u.G(bob.Ks[6].K0, bob.Ks[7].K0, 3), bob.KSAND(9, 0, 0))
	G3.C2 = u.XOR(u.G(bob.Ks[6].K1, bob.Ks[7].K0, 3), bob.KSAND(9, 1, 0))
	G3.C3 = u.XOR(u.G(bob.Ks[6].K0, bob.Ks[7].K1, 3), bob.KSAND(9, 0, 1))
	G3.C4 = u.XOR(u.G(bob.Ks[6].K1, bob.Ks[7].K1, 3), bob.KSAND(9, 1, 1))
	gates[3] = G3

	// (z4_1 & z3_3)
	G4 := t.Gate{}
	G4.C1 = u.XOR(u.G(bob.Ks[8].K0, bob.Ks[9].K0, 4), bob.KSAND(10, 0, 0))
	G4.C2 = u.XOR(u.G(bob.Ks[8].K1, bob.Ks[9].K0, 4), bob.KSAND(10, 1, 0))
	G4.C3 = u.XOR(u.G(bob.Ks[8].K0, bob.Ks[9].K1, 4), bob.KSAND(10, 0, 1))
	G4.C4 = u.XOR(u.G(bob.Ks[8].K1, bob.Ks[9].K1, 4), bob.KSAND(10, 1, 1))
	gates[4] = G4

	return gates
}
