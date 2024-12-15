package actors

import (
	"math/rand"
)

type Row [8]uint8
type Matrix [8]Row

type Dealer struct {
	T    Matrix
	R, S uint8
	Mb   Matrix
}

func RandM() Matrix {

	M := Matrix{}

	for i := 0; i < 8; i++ {
		next := [8]uint8{}
		for j := 0; j < 8; j++ {
			x := rand.Intn(2)
			if x > 1 {
				panic("random integer out of bound")
			}
			next[j] = uint8(x)
		}

		M[i] = next
	}
	return M
}

func (d *Dealer) Ma() Matrix {
	Ma := Matrix{}

	for i := 0; i < 8; i++ {
		next := Row{}

		for j := 0; j < 8; j++ {
			a := (uint8(i) - d.R) % 8
			b := (uint8(j) - d.S) % 8

			next[j] = d.Mb[i][j] ^ d.T[a][b]
		}

		Ma[i] = next
	}
	return Ma
}

func DealerInit(T Matrix) *Dealer {
	r := uint8(rand.Intn(8))
	s := uint8(rand.Intn(8))

	Mb := RandM()

	return &Dealer{R: r, S: s, Mb: Mb, T: T}
}

func (d *Dealer) ToAlice() (uint8, Matrix) {
	return d.R, d.Ma()
}

func (d *Dealer) ToBob() (uint8, Matrix) {
	return d.S, d.Mb
}
