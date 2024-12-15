package functions

import "handin1/types"

var LOOKUP_TABLE = [8][8]uint{
	//O- O+ B- B+ A- A+ AB- AB+
	{1, 0, 0, 0, 0, 0, 0, 0}, // O_MINUS
	{1, 1, 0, 0, 0, 0, 0, 0}, // O_PLUS
	{1, 0, 1, 0, 0, 0, 0, 0}, // B_MINUS
	{1, 1, 1, 1, 0, 0, 0, 0}, // B_PLUS
	{1, 0, 0, 0, 1, 0, 0, 0}, // A_MINUS
	{1, 1, 0, 0, 1, 1, 0, 0}, // A_PLUS
	{1, 0, 1, 0, 1, 0, 1, 0}, // AB_MINUS
	{1, 1, 1, 1, 1, 1, 1, 1}, // AB_PLUS
}

func Lookup(reciever types.TYPES, giver types.TYPES) bool {
	return LOOKUP_TABLE[reciever][giver] == 1
}
