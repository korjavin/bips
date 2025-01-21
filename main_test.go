package main

import (
	"testing"

	"github.com/ethereum/go-ethereum/crypto"
)

func TestPubKey(t *testing.T) {
	seedPhrase := "elbow inmate boy drill divide device noble ecology fog runway potato guilt"
	tests := []struct {
		index           uint32
		expectedAddress string
	}{
		{
			index:           0,
			expectedAddress: "0xC49fb5CF14b357f993bA5FA76BE3Dd438177d5d3",
		},
		{
			index:           1,
			expectedAddress: "0xD1c12fA87B0BFdB481aC7aF52fe056BB4f5A7eA6",
		},
	}

	for _, tt := range tests {
		output := PubKey(seedPhrase, tt.index)
		address := crypto.PubkeyToAddress(output)

		// Convert the address to checksummed format
		checksummedAddress := address.Hex()

		if checksummedAddress != tt.expectedAddress {
			t.Errorf("for index %d, expected %s, got %s", tt.index, tt.expectedAddress, checksummedAddress)
		}
	}
}

func TestPrivKeyHex(t *testing.T) {
	seedPhrase := "elbow inmate boy drill divide device noble ecology fog runway potato guilt"
	tests := []struct {
		index           uint32
		expectedPrivKey string
	}{
		{
			index:           0,
			expectedPrivKey: "39a0ea5dd4f3941b87ddbcd84207e57a1c9d19d8cfef84865b2260d566c9962a",
		},
		{
			index:           1,
			expectedPrivKey: "282e6f994d81975305c845939a95fd75a845ce0a07816d7ad5ab6aa18e404f7b",
		},
	}

	for _, tt := range tests {
		output := PrivKeyHex(seedPhrase, tt.index)

		if output != tt.expectedPrivKey {
			t.Errorf("for index %d, expected %s, got %s", tt.index, tt.expectedPrivKey, output)
		}
	}
}

//Understanding EIP-55 Checksum Encoding
/*
A checksum is a small-sized datum derived from a block of digital data for the purpose of detecting errors that may have been introduced during its transmission or storage. In the context of Ethereum addresses, the checksum helps ensure that the address is correctly copied and pasted.

How EIP-55 Works
Lowercase Address : Start with the lowercase hexadecimal representation of the address.
Keccak-256 Hash : Compute the Keccak-256 hash of the lowercase address.
Checksum Calculation : For each character in the original address:
If the corresponding character in the hash is a 8 or higher, convert the character in the address to uppercase.
Otherwise, keep the character in the address as lowercase.
Example
Given the address c49fb5cf14b357f993ba5fa76be3dd438177d5d3:

Compute the Keccak-256 hash of the lowercase address: 9b2055d370f73ec7d8a03e965129118dc8f5bf83e550736d74d5bf1f52930bd9.
Compare each character of the address with the corresponding character in the hash:
If the hash character is 8 or higher, convert the address character to uppercase.
Otherwise, keep the address character lowercase.
This results in the checksummed address: C49fb5CF14b357f993bA5FA76BE3Dd438177d5d3.
*/
