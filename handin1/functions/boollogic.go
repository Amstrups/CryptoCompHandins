package functions

import (
	t "handin1/types"
)

// [3]uint8{ is plus, contains B-gene, contains A-gene }
func toInt(ty t.TYPES) [3]uint8 {
	arr := [3]uint8{}
	arr[2] = 1 & uint8(ty)
	arr[1] = (2 & uint8(ty)) >> 1
	arr[0] = (4 & uint8(ty)) >> 2
	return arr
}

func OriginalAnswer(rec [3]uint8, giv [3]uint8) uint8 {
	sign := (^(rec[0] & giv[0]) ^ giv[0]) & 1

	trail_eq := ^(rec[2] ^ giv[2]) & 1
	lead_eq := ^(rec[1] ^ giv[1]) & 1
	eq := lead_eq & trail_eq

	is_ab := rec[2] & rec[1]
	is_o := (^giv[2] & ^giv[1]) & 1

	return sign & (is_ab | is_o | eq)

}

// [3]uint8{ contains A-gene, contains B-gene, is plus }
func Compatibility(reciever t.TYPES, donor t.TYPES) uint8 {
	rec := toInt(reciever)
	giv := toInt(donor)

	lt := func(i int) uint8 {
		return (1 ^ ((1 ^ rec[i]) & giv[i]))
	}

	return (lt(0) & lt(1)) & lt(2)
}

