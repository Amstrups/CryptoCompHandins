package types

type Gate struct {
	C1, C2, C3, C4 []byte
}

func (g Gate) Permutated() [][]byte {
	C := [][]byte{g.C1, g.C2, g.C3, g.C4}

	return C
}

