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

func ToInt(ty TYPES) [3]int64 {
	arr := [3]int64{}
	arr[2] = int64(1 & uint8(ty))
	arr[1] = int64((2 & uint8(ty)) >> 1)
	arr[0] = int64((4 & uint8(ty)) >> 2)
	return arr
}

