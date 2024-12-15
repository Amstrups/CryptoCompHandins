package main

import (
	"fmt"
	"math"
	"math/rand"
)

type TYPES int8

const (
	O_MINUS = iota
	O_PLUS
	B_MINUS
	B_PLUS
	A_MINUS
	A_PLUS
	AB_MINUS
	AB_PLUS
)

func toInt(ty TYPES) [3]uint8 {
	arr := [3]uint8{}
	arr[2] = 1 & uint8(ty)
	arr[1] = (2 & uint8(ty)) >> 1
	arr[0] = (4 & uint8(ty)) >> 2
	return arr
}

func Compatibility(reciever TYPES, donor TYPES) uint8 {
	rec := toInt(reciever)
	giv := toInt(donor)

	lt := func(i int) uint8 {
		return (1 ^ ((1 ^ giv[i]) & rec[i]))
	}

	return (lt(0) & lt(1)) & lt(2)
}

func pow(x, y uint) uint {
	return uint(math.Pow(float64(x), float64(y)))
}

var allTypes = []TYPES{
	O_MINUS, O_PLUS, B_MINUS, B_PLUS, A_MINUS, A_PLUS, AB_MINUS, AB_PLUS,
}

type PK struct {
	g, h uint
}

type ENC struct {
	c1, c2 uint
}

const (
	Q = 8
)

type Alice struct {
	ty TYPES
	sk uint

	results []uint
}

func G() uint {
	generators := []uint{3, 5, 7}
	index := rand.Intn(3)
	return generators[index]

}

func (alice *Alice) Generate() [Q]PK {
	pks := [Q]PK{}

	g := G()

	for i := range pks {
		pk := PK{g: g}
		if i == int(alice.ty) { // Run Gen
			alice.sk = uint(rand.Intn(Q))

			pk.h = pow(g, alice.sk) % Q
		} else { // Run OGen
			q := Q
			p := 2*q + 1

			s := rand.Intn(p)

			for s < 2 {
				s = rand.Intn(p)
			}
			h := pow(uint(s), 2) % uint(p)
			pk.h = h
		}
		pks[i] = pk
	}

	return pks
}

func (alice *Alice) Decrypt(msgs [Q]ENC) uint {
	enc := msgs[alice.ty]

	msg := (enc.c2 * (pow(enc.c1, (-alice.sk)) % Q)) % Q


	return (msg >> 2)
}

type Bob struct {
	ty TYPES
}

func (bob *Bob) Encrypt(pks [Q]PK) [Q]ENC {
	cts := [Q]ENC{}
	for i, t := range allTypes {
		comp := uint(Compatibility(t, bob.ty))

		comp = (comp << 2) + uint(rand.Intn(4))

		r := uint(rand.Intn(6)) + 1

		pk := pks[i]

		c1 := pow(pk.g, r) % Q
		c2 := (comp * (pow(pk.h, r) % Q)) % Q

		cts[i] = ENC{c1, c2}
	}

	return cts
}

func main() {
	M := [8][8]uint{}
	for i, ti := range allTypes {
		for j, tj := range allTypes {
			alice := &Alice{ty: ti}
			bob := &Bob{tj}

			M[7-j][i] = alice.Decrypt(bob.Encrypt(alice.Generate()))
		}
	}

	for _, row := range M {
		fmt.Println(row)
	}
}
