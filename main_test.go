package main

import (
	"bytes"
	"testing"

	"github.com/ethereum/go-ethereum/crypto"
)

func TestPubKey(t *testing.T) {

	input := "elbow inmate boy drill divide device noble ecology fog runway potato guilt"

	expectedPublicKey := "0xc49fb5cf14b357f993ba5fa76be3dd438177d5d3" // Replace with actual expected public key
	//expectedPrivateKey := "39a0ea5dd4f3941b87ddbcd84207e57a1c9d19d8cfef84865b2260d566c9962a"

	output := PubKey(input)
	address := crypto.PubkeyToAddress(output)

	if !bytes.Contains([]byte(address.Hex()), []byte(expectedPublicKey)) {
		t.Errorf("expected %s to be in output, got %s", expectedPublicKey, address.Hex())
	}
}
