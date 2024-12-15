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

func AsString(in TYPES) string {
	switch in {
	case O_PLUS:
		return "O_PLUS"
	case O_MINUS:
		return "O_MINUS"
	case A_PLUS:
		return "A_PLUS"
	case A_MINUS:
		return "A_MINUS"
	case B_PLUS:
		return "B_PLUS"
	case B_MINUS:
		return "B_MINUS"
	case AB_PLUS:
		return "AB_PLUS"
	case AB_MINUS:
		return "AB_MINUS"
	}
	return "Not a blood type"
}
