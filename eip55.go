package main

import (
	"github.com/ethereum/go-ethereum/crypto"
)

// ToChecksumAddress generates the checksummed Ethereum address
func ToChecksumAddress(address string) string {
	hash := crypto.Keccak256([]byte(address[1:]))
	checksummedAddress := "0x"

	for i := 2; i < len(address); i++ {
		char := address[i]
		hashByte := hash[i/2]

		if char >= '0' && char <= '9' {
			checksummedAddress += string(char)
		} else {
			nibble := hashByte >> (4 * (1 - (i % 2)))
			if nibble > 7 {
				checksummedAddress += string(char &^ 0x20) // Uppercase
			} else {
				checksummedAddress += string(char | 0x20) // Lowercase
			}
		}
	}

	return checksummedAddress
}
