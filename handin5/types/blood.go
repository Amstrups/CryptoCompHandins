package types

type TYPES uint8

const (
	O_MINUS = iota
	O_PLUS
	B_MINUS
	B_PLUS
	A_MINUS
	A_PLUS
	AB_MINUS
	AB_PLUS
)

var AllTypes = []TYPES{
	O_MINUS, O_PLUS, B_MINUS, B_PLUS, A_MINUS, A_PLUS, AB_MINUS, AB_PLUS,
}
