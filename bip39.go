package main

import (
	"log"

	"github.com/tyler-smith/go-bip39"
)

func generateSeedPhrase() string {
	entropy, err := bip39.NewEntropy(128) // 128 bits for a 12-word mnemonic
	if err != nil {
		log.Fatalf("Failed to generate entropy: %v", err)
	}

	mnemonic, err := bip39.NewMnemonic(entropy)
	if err != nil {
		log.Fatalf("Failed to generate mnemonic: %v", err)
	}

	return mnemonic
}
