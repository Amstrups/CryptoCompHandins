package types

import "crypto/rand"

type _K = []byte
type K = [16]byte
type K2 = [32]byte


type KPair struct {
	K0, K1 K
}

func (k *KPair) Choose(i uint8) K {
	if i == 0 {
		return k.K0
	}
	return k.K1
}

const (
	T = 11
)

func MakeK() K {
	tok := make(_K, 16)
	_, err := rand.Read(tok)
	for err != nil {
		rand.Read(tok)
	}

	return K(tok)
}
