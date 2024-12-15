package functions

import (
	"handin2/actors"
	"handin2/types"
	"testing"
)

func TestAliceSend1(t *testing.T) {
	d := actors.DealerInit(LOOKUP_TABLE)

	alice := actors.AliceInit(types.O_MINUS, d)

	alice_R := alice.Send()

	expected := d.R

	if alice_R != expected {
		t.FailNow()
	}
}

func TestAliceSend2(t *testing.T) {
	var chosen types.TYPES = types.AB_MINUS // = 110 = 6
	d := &actors.Dealer{R: 2}

	alice := actors.AliceInit(chosen, d)

	alice_R := alice.Send()

	var expected uint8 = 0 // 6 + 2 % 8 = 0

	if alice_R != expected {
		t.FailNow()
	}
}

func TestBobSend1(t *testing.T) {
	var bob_chosen types.TYPES = types.O_MINUS // = 000 = 0

	d := &actors.Dealer{R: 2, S: 4}

	alice := actors.AliceInit(types.AB_PLUS, d)
	bob := actors.BobInit(bob_chosen, d)

	bob.Receive(alice.Send())

	bob_v, bob_zb := bob.Send()

	var expected_v uint8 = 4  //
	var expected_zb uint8 = 0 // default value

	if bob_v != expected_v {
		t.Fatal("v is not equal")
	}

	if bob_zb != expected_zb {
		t.Fatal("z_b is not equal")
	}
}

func TestBobSend2(t *testing.T) {
	var bob_chosen types.TYPES = types.A_MINUS // = 010 = 4

	d := &actors.Dealer{R: 2, S: 4}
	d.Mb[0][0] = 1

	bob := actors.BobInit(bob_chosen, d)
	bob.U = 0

	bob_v, bob_zb := bob.Send()

	var expected_v uint8 = 0 // 4 + 4 % 8 = 0
	var expected_zb uint8 = 1

	if bob_v != expected_v {
		t.Fatal("v is not equal")
	}

	if bob_zb != expected_zb {
		t.Fatal("z_b is not equal")
	}
}

func TestOTTTLookup(t *testing.T) {
	for _, receiver := range allTypes {
		for _, donor := range allTypes {

			isMatch := FullOTTTRun(receiver, donor)

			if isMatch != expectations[receiver][donor] {
				t.FailNow()
			}
		}
	}
}
