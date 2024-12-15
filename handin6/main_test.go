package main

import (
	"handin6/types"
	"math/big"
	"testing"
)

const (
	n = 10
	s = 5
)

func BenchmarkPK1(t *testing.B) {
	sk := NewSK()
	NewPK(sk, n)

}

func BenchmarkPK2(t *testing.B) {
	sk := NewSK()
	NewPK2(sk, n)
}

func BenchmarkFHE0(t *testing.B) {
	M := M{}

	n := 10
	s := 5

	for i, receiver := range types.AllTypes {
		sk := NewSK()
		pk, _ := NewPK(sk, n)
		for j, donor := range types.AllTypes {
			X := [3]*big.Int{}
			Y := [3]*big.Int{}
			for k, xi := range types.ToInt(receiver) {
				big_xi := big.NewInt(xi)
				mi := Enc(big_xi, pk, s)
				X[k] = mi
			}
			for k, yi := range types.ToInt(donor) {
        yii := int64(uint8(yi) ^ 1)
				big_yi := big.NewInt(yii)
				mi := Enc(big_yi, pk, 10)
				Y[k] = mi
			}
			c := Eval(X, Y)
			_m := Dec(c, sk)

			M[i][7-j] = _m
		}
	}


}


func BenchmarkFHE1(t *testing.B) {
	M := M{}

	n := 10
	s := 5

	for i, receiver := range types.AllTypes {
		sk := NewSK()
		pk, _ := NewPK(sk, n)
		for j, donor := range types.AllTypes {
			X := [3]*big.Int{}
			Y := [3]*big.Int{}
			for k, xi := range types.ToInt(receiver) {
				big_xi := big.NewInt(xi)
				mi := Enc(big_xi, pk, s)
				X[k] = mi
			}
			for k, yi := range types.ToInt(donor) {
        yii := int64(uint8(yi) ^ 1)
				big_yi := big.NewInt(yii)
				mi := Enc(big_yi, pk, 10)
				Y[k] = mi
			}
			c := Eval(X, Y)
			_m := Dec(c, sk)

			M[i][7-j] = _m
		}
	}


}

func BenchmarkFHE2(t *testing.B) {
	M := M{}

	n := 10
	s := 5

	for i, receiver := range types.AllTypes {
		sk := NewSK()
		pk, _ := NewPK2(sk, n)
		for j, donor := range types.AllTypes {
			X := [3]*big.Int{}
			Y := [3]*big.Int{}
			for k, xi := range types.ToInt(receiver) {
				big_xi := big.NewInt(xi)
				mi := Enc(big_xi, pk, s)
				X[k] = mi
			}
			for k, yi := range types.ToInt(donor) {
        yii := int64(uint8(yi) ^ 1)
				big_yi := big.NewInt(yii)
				mi := Enc(big_yi, pk, 10)
				Y[k] = mi
			}
			c := Eval(X, Y)
			_m := Dec(c, sk)

			M[i][7-j] = _m
		}
	}


}
