package actors

import (
	t "handin2/types"
)

type Alice struct {
	Blood t.TYPES
	R     uint8
	M     Matrix
	u     uint8
	v     uint8
	z_b   uint8
}

func AliceInit(blood t.TYPES, d *Dealer) *Alice {
	r, M := d.ToAlice()

	return &Alice{
		Blood: blood,
		R:     r,
		M:     M,
		u:     255,
		v:     255,
		z_b:   255,
	}
}
func (a *Alice) Receive(v uint8, z_b uint8) {
	a.v = v
	a.z_b = z_b
}

func (a *Alice) Send() uint8 {
	a.u = (uint8(a.Blood) + a.R) % 8

	return a.u
}

func (a *Alice) Output() uint8 {
	return a.M[a.u][a.v] ^ a.z_b
}
