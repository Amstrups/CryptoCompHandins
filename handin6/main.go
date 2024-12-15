package main

import (
	crand "crypto/rand"
	"fmt"
	"handin6/types"
	"math"
	"math/big"
	"math/rand"
)

// P: 2000 bits
// q_i: 10^7 bits
// r_i: 60 bits

type B []byte

const (
	PLen = 2000
	QLen = 1000000
	RLen = 60
)

type SK struct {
	P *big.Int
}

func NewSK() *SK {
	p := new(big.Int)
	p.SetBytes(Rand(PLen))
	return &SK{P: p}
}

type PK struct {
	Y []*big.Int
}


func NewPK(SK *SK, n int) (*PK, error) {
	TWO := big.NewInt(2)

	Y := make([]*big.Int, n)

	for i := 0; i < n; i++ {
		yn := new(big.Int)

		ri := new(big.Int)
		ri.SetBytes(Rand(RLen))
		ri.Mul(ri, TWO)

		qi := new(big.Int)
		qi.SetBytes(Rand(QLen))

		pq := new(big.Int)
		pq.Mul(SK.P, qi)

		yn.Add(yn, ri)
		yn.Add(yn, pq)
		Y[i] = yn
	}

	return &PK{Y: Y}, nil
}

func NewPK2(SK *SK, n int) (*PK, error) {

	Y := make([]*big.Int, n)

	for i := 0; i < n; i++ {
		yn := new(big.Int)

    ri_bytes := Rand(RLen)
    ri_bytes = append(ri_bytes, 0)

		ri := new(big.Int)
		ri.SetBytes(ri_bytes)

		qi := new(big.Int)
		qi.SetBytes(Rand(QLen))

		pq := new(big.Int)
		pq.Mul(SK.P, qi)

		yn.Add(yn, ri)
		yn.Add(yn, pq)
		Y[i] = yn
	}

	return &PK{Y: Y}, nil
}

func Rand(_k uint) []byte {
	k := math.Ceil(float64(_k / 8))

	by := make(B, int(k)+1)
	_, err := crand.Read(by)
	for err != nil {
		crand.Read(by)
	}

	return by
}

func Enc(m *big.Int, pk *PK, s int) *big.Int {
	perm := rand.Perm(len(pk.Y))
	yn := new(big.Int)
	for i := 0; i < s; i++ {
		index := perm[i]
		yn.Add(yn, pk.Y[index])
	}

	c := new(big.Int)
	c.Add(m, yn)
	return c
}

func Dec(c *big.Int, sk *SK) *big.Int {
	TWO := big.NewInt(2)
	m1 := new(big.Int).Mod(c, sk.P)
	m2 := new(big.Int).Mod(m1, TWO)

	return m2
}

func Eval(X [3]*big.Int, Y [3]*big.Int) *big.Int {
	xy1 := new(big.Int).Mul(X[0], Y[0])
	xy2 := new(big.Int).Mul(X[1], Y[1])
	xy3 := new(big.Int).Mul(X[2], Y[2])

	xy1.Add(xy1, big.NewInt(1))
	xy2.Add(xy2, big.NewInt(1))
	xy3.Add(xy3, big.NewInt(1))

	xy12 := new(big.Int).Mul(xy1, xy2)
	xy123 := new(big.Int).Mul(xy12, xy3)

	return xy123
}

func out(x *big.Int, a bool) {
	s := fmt.Sprintf("%b", x)

	if a {
		fmt.Println(len(s), s)
	} else {
		fmt.Println(len(s))
	}
}

type M [8][8]*big.Int

func main() {
	M := M{}

	n := 10
	s := 5

	for i, receiver := range types.AllTypes {
		sk := NewSK()
		pk, _ := NewPK(sk, n)
		for j, donor := range types.AllTypes {
			fmt.Printf("Round %d:%d\n", i+1, j+1)
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

	for _, m := range M {
		fmt.Println(m)

	}

}
