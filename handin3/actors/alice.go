package actors

import (
	t "handin3/types"
	"math/rand"
)

type Alice struct {
	Blood           t.TYPES
	sends, received int
	Values
	xs        [3]uint8
	ys        [3]uint8
	zs        [3]uint8
	d         uint8
	e         uint8
	HasOutput bool
}

func AliceInit(blood t.TYPES, d *Dealer) *Alice {
	return &Alice{
		Blood:  blood,
		Values: d.ToAlice(),
		xs:     toInt(blood),
	}
}

func (alice *Alice) Send() uint8 {
	alice.sends++

	switch alice.sends {
	case 1, 2, 3: // Share random bits for input with Bob
		index := alice.sends - 1

		x_r := uint8(rand.Intn(2))
		x_b := x_r ^ alice.xs[index]

		alice.xs[index] = x_b

		return x_r
	case 4:
		return 0
	case 5, 7, 9: // Compute Alice's (d,e) in AND
		index := int((alice.sends - 5) / 2)

		alice.d = alice.xs[index] ^ alice.Values.Us[index]
		alice.e = alice.ys[index] ^ alice.Values.Vs[index]

		return alice.d
	case 6, 8, 10, 12, 14:
		return alice.e
	case 11, 13: // Compute Alice's (d,e) in AND
		index := int((alice.sends-11)/2) + 3

		alice.d = alice.zs[0] ^ alice.Values.Us[index]
		alice.e = alice.zs[1] ^ alice.Values.Vs[index]

		return alice.d
	case 15:
		// protocol last round, bob needs no more information
		return 0
	}
	return 0
}

func (alice *Alice) Receive(value uint8) {
	alice.received++

	switch alice.received {
	case 1, 2, 3: // Receive Y from Bob
		alice.ys[alice.received-1] = value
		break
	case 4: // XOR with constant, only for Alice
		for i := 0; i < len(alice.ys); i++ {
			alice.ys[i] ^= 1
		}
		break
	case 5, 7, 9, 11, 13: // Receive Bob's part of d in "AND of Two Wires"
		alice.d ^= value
		break
	case 6, 8, 10: // Subprotocol 6 in "AND of Two Wires" for most inner ANDS
		index := int((alice.received - 6) / 2)

		alice.e ^= value

		x := alice.xs[index]
		y := alice.ys[index]

		alice_z := alice.Ws[index] ^ (alice.e & x) ^ (alice.d & y) ^ alice.e&alice.d

		alice.zs[index] = alice_z
		alice.zs[index] ^= 1 // XOR with constant, only for alice

		break
	case 12, 14: // Subprotocol 6 in "AND of Two Wires" for outputs of first 3 AND gates
		index := max(0, alice.received-13) + 3

		alice.e ^= value

		alice_x := alice.zs[0]
		alice_y := alice.zs[1]

		alice_z := alice.Ws[index] ^ (alice.e & alice_x) ^ (alice.d & alice_y) ^ alice.e&alice.d
		alice.zs[0] = alice_z

		alice.zs[1] = alice.zs[2]

		break
	case 15: // (z,_) <- OpenTo(A,[z])
		alice.zs[0] ^= value
		alice.HasOutput = true
		break
	}
}

func (alice *Alice) Output() uint8 {
	return alice.zs[0]
}
