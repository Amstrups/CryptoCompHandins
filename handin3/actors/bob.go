package actors

import (
	t "handin3/types"
	"math/rand"
)

type Bob struct {
	Blood           t.TYPES
	sends, received int
	Values
	xs [3]uint8
	ys [3]uint8
	zs [3]uint8
	d  uint8
	e  uint8
}

func BobInit(blood t.TYPES, d *Dealer) *Bob {
	return &Bob{
		Blood:  blood,
		Values: d.ToBob(),
		ys:     toInt(blood),
	}
}

func (bob *Bob) Receive(value uint8) {
	bob.received++

	switch bob.received {
	case 1, 2, 3: // Receive X from Alice
		bob.xs[bob.received-1] = value
		return
	case 4:
		break
	case 5, 7, 9, 11, 13: // Receive Alice's d for AND
		bob.d = value
		break
	case 6, 8, 10, 12, 14: // Receive Alice's e for AND
		bob.e = value
		break
	case 15: // protocol last round
		return
	}
}

func (bob *Bob) Send() uint8 {
	bob.sends++

	switch bob.sends {
	case 1, 2, 3: // Share random bits for input with Alice
		index := bob.sends - 1

		y_r := uint8(rand.Intn(2))
		y_b := y_r ^ bob.ys[index]

		bob.ys[index] = y_b

		return y_r
	case 4: // Alice XOR with 1, Bob does nothing
		return 0
	case 5, 7, 9: // Send Bob's part of d in "AND of Two Wires"
		index := int((bob.sends - 5) / 2)

		d := bob.xs[index] ^ bob.Us[index]

		bob.d ^= d

		return d
	case 6, 8, 10: // Subprotocol 6 in "AND of Two Wires" for most inner ANDS
		index := int((bob.sends - 6) / 2)

		e := bob.ys[index] ^ bob.Vs[index]

		bob.e ^= e

		bob_x := bob.xs[index]
		bob_y := bob.ys[index]

		bob_z := bob.Ws[index] ^ (bob.e & bob_x) ^ ((bob.d & bob_y) & bob.d)

		bob.zs[index] = bob_z

		return e
	case 11, 13: // Send Bob's part of d in "AND of Two Wires"
		index := max(0, bob.sends-12) + 3
		d := bob.zs[0] ^ bob.Us[index]

		bob.d ^= d

		return d
	case 12, 14: // Subprotocol 6 in "AND of Two Wires" for outputs of first 3 AND gates
		index := max(0, bob.sends-13) + 3
		e := bob.zs[1] ^ bob.Vs[index]

		bob.e ^= e

		bob_x := bob.zs[0]
		bob_y := bob.zs[1]

		bob_z := bob.Ws[index] ^ (bob.e & bob_x) ^ ((bob.d & bob_y) & bob.d)

		bob.zs[0] = bob_z

		bob.zs[1] = bob.zs[2]

		return e
	case 15: // Share Bob's z with Alice
		return bob.zs[0]
	}
	return 0
}
