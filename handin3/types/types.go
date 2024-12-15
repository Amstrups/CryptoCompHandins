package types

type TYPES int8

const (
	// 000 or 0
	O_MINUS = iota
	// 001 or 1
	O_PLUS
	// 010 or 2
	B_MINUS
	// 011 or 3
	B_PLUS
	// 100 or 4
	A_MINUS
	// 101 or 5
	A_PLUS
	// 110 or 6
	AB_MINUS
	// 111 or 7
	AB_PLUS
)
