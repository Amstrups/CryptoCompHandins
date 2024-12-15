package util

import (
	"crypto/sha256"
	t "handin5/types"
)

func G(K0 t.K, K1 t.K, i int) []byte {
	h := sha256.New()

	k := make([]byte, 33)

	for i, ki := range K0 {
		k[i] = ki
		k[i+16] = K1[i]
	}

	k[32] = byte(i)

	h.Write(k)
	bs := h.Sum(nil)

	return bs
}

func XOR(left []byte, right [16]byte) []byte {
	result := make([]byte, 32)
	for i, b := range left {
		if i >= 16 {
			result[i] = b ^ 0

		} else {
			result[i] = right[i] ^ b
		}
	}

	return result
}

func XOR32(left [32]byte, right [32]byte) [32]byte {
	result := [32]byte{}
	for i, b := range left {
		result[i] = right[i] ^ b
	}
	return result

}

func EQ(left [16]byte, right [16]byte) bool {
	if len(left) != len(right) {
		return false
	}

	for i := 0; i < 16; i++ {
		if left[i] != right[i] {
			return false
		}
	}

	return true
}

var BitCombinations = [4][2]uint8{
	[2]uint8{0, 0},
	[2]uint8{1, 0},
	[2]uint8{0, 1},
	[2]uint8{1, 1},
}
