package main

import (
	"bufio"
	"crypto/ecdsa"
	"fmt"
	"log"
	"os"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/tyler-smith/go-bip32"
	"github.com/tyler-smith/go-bip39"
)

func PubKey(mnemonic string, addressIndex uint32) ecdsa.PublicKey {
	seed := bip39.NewSeed(mnemonic, "")
	masterKey, err := bip32.NewMasterKey(seed)
	if err != nil {
		log.Fatal(err)
	}

	// Derive the path m/44'/60'/0'/0/addressIndex
	//BIP 44 defines a hierarchical deterministic (HD) wallet path structure for generating multiple addresses from a single seed. The path is a sequence of indices that specify the hierarchy levels. Each level in the path is represented by a 32-bit unsigned integer.
	// Hardened Keys (Prime Notation, e.g., 44') :
	// These keys are derived using a different method that involves the private key.
	// They are denoted with a prime symbol (').
	// Non-Hardened Keys (e.g., 0) :
	// These keys are derived using the parent public key.
	// They are less secure but can be useful for certain applications where you need to derive public keys without exposing the private key.
	path := []uint32{
		bip32.FirstHardenedChild + 44, // 44' Use the BIP 44 standard.
		bip32.FirstHardenedChild + 60, // 60' Use the Ethereum coin type.
		bip32.FirstHardenedChild + 0,  // 0' Use the default account.
		0,                             // 0 Use external addresses.
		addressIndex,                  // Use the specified address index.
	}

	var childKey *bip32.Key = masterKey
	for _, index := range path {
		childKey, err = childKey.NewChildKey(index)
		if err != nil {
			log.Fatal(err)
		}
	}

	privKey, err := crypto.ToECDSA(childKey.Key)
	if err != nil {
		log.Fatal(err)
	}

	return privKey.PublicKey
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter mnemonic: ")
	mnemonic, _ := reader.ReadString('\n')
	mnemonic = mnemonic[:len(mnemonic)-1] // Remove newline character

	pubKey := PubKey(mnemonic, 0)
	address := crypto.PubkeyToAddress(pubKey)
	fmt.Printf("Ethereum Address: %s\n", address.Hex())
}
