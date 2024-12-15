package actors

import t "handin3/types"

type Values struct {
	Us [5]uint8
	Vs [5]uint8
	Ws [5]uint8
}

// [3]uint8{ is plus, contains B-gene, contains A-gene }
func toInt(ty t.TYPES) [3]uint8 {
	arr := [3]uint8{}
	arr[2] = 1 & uint8(ty)
	arr[1] = (2 & uint8(ty)) >> 1
	arr[0] = (4 & uint8(ty)) >> 2
	return arr
}

// a = 1 XOR (( 1 XOR x_0 ) AND y_0 )
// b = 1 XOR (( 1 XOR x_1 ) AND y_1 )
// c = 1 XOR (( 1 XOR x_2 ) AND y_2 )
// return ( a AND b ) AND c
func Compatibility(reciever t.TYPES, donor t.TYPES) uint8 {
	rec := toInt(reciever)
	giv := toInt(donor)

	lt := func(i int) uint8 {
		return (1 ^ ((1 ^ giv[i]) & rec[i]))
	}

	return (lt(0) & lt(1)) & lt(2)
}
