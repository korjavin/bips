package main

import (
	"strings"

	"github.com/ethereum/go-ethereum/crypto"
)

// ToChecksumAddress converts an Ethereum address to its checksummed version
// as per EIP-55. The function takes a hexadecimal address string as input,
// converts it to lowercase, and then computes the Keccak-256 hash of the
// address (excluding the '0x' prefix). It then iterates over each character
// of the address and applies the checksum rules:
//   - If the character is a letter (i.e., greater than '9') and the corresponding
//     hash nibble is greater than or equal to 8, the letter is converted to uppercase.
//   - Otherwise, the character remains unchanged.
//
// The resulting checksummed address is returned as a string.
func ToChecksumAddress(address string) string {
	address = strings.ToLower(address)
	hash := crypto.Keccak256([]byte(address[2:]))
	checksummedAddress := "0x"

	for i := 0; i < len(address[2:]); i++ {
		if address[2+i] > '9' && hash[i/2]>>uint(4*(1-i%2))&0xf >= 8 {
			checksummedAddress += string(address[2+i] - 32)
		} else {
			checksummedAddress += string(address[2+i])
		}
	}

	return checksummedAddress
}
