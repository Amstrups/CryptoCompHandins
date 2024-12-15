package util

import (
	"math"
	"math/big"
)

type PK struct {
	G, P, H *big.Int
}

type SK struct {
	PK
	X *big.Int
}

func pow(x, y uint) uint {
	return uint(math.Pow(float64(x), float64(y)))
}

func Encrypt(pub *PK, m []byte) (c1, c2 *big.Int) {


	return 0, 0
}

func Decrypt(sk *SK, c1, c2 *big.Int) (m []byte) {

	return []byte{}
}
