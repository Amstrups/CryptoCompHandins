package actors

import t "handin2/types"

type Bob struct {
	Blood t.TYPES
	S     uint8
	M     Matrix
	U     uint8
}

func BobInit(blood t.TYPES, d *Dealer) *Bob {
	s, M := d.ToBob()
	return &Bob{
		Blood: blood,
		S:     s,
		M:     M,
		U:     255,
	}
}
func (b *Bob) Receive(u uint8) {
	b.U = u
}

func (b *Bob) Send() (uint8, uint8) {
	if b.U == 255 {
		panic(1)
	}
	v := (uint8(b.Blood) + b.S) % 8

	return v, b.M[b.U][v]
}
