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

// Derive the path m/44'/60'/0'/0/addressIndex
//BIP 44 defines a hierarchical deterministic (HD) wallet path structure for generating multiple addresses from a single seed. The path is a sequence of indices that specify the hierarchy levels. Each level in the path is represented by a 32-bit unsigned integer.
// Hardened Keys (Prime Notation, e.g., 44') :
// These keys are derived using a different method that involves the private key.
// They are denoted with a prime symbol (').
// Non-Hardened Keys (e.g., 0) :
// These keys are derived using the parent public key.
// They are less secure but can be useful for certain applications where you need to derive public keys without exposing the private key.

func GeneratePath(addressIndex uint32) []uint32 {
	return []uint32{
		bip32.FirstHardenedChild + 44,
		bip32.FirstHardenedChild + 60,
		bip32.FirstHardenedChild + 0,
		0,
		addressIndex,
	}
}

func DeriveChildKey(masterKey *bip32.Key, path []uint32) (*bip32.Key, error) {
	var err error
	childKey := masterKey
	for _, index := range path {
		childKey, err = childKey.NewChildKey(index)
		if err != nil {
			return nil, err
		}
	}
	return childKey, nil
}
func PubKey(mnemonic string, addressIndex uint32) ecdsa.PublicKey {
	seed := bip39.NewSeed(mnemonic, "")
	masterKey, err := bip32.NewMasterKey(seed)
	if err != nil {
		log.Fatal(err)
	}
	path := GeneratePath(addressIndex)
	childKey, err := DeriveChildKey(masterKey, path)
	if err != nil {
		log.Fatal(err)
	}

	privKey, err := crypto.ToECDSA(childKey.Key)
	if err != nil {
		log.Fatal(err)
	}

	return privKey.PublicKey
}

func PrivKeyHex(mnemonic string, addressIndex uint32) string {
	seed := bip39.NewSeed(mnemonic, "")
	masterKey, err := bip32.NewMasterKey(seed)
	if err != nil {
		log.Fatal(err)
	}

	path := GeneratePath(addressIndex)
	childKey, err := DeriveChildKey(masterKey, path)
	if err != nil {
		log.Fatal(err)
	}

	return fmt.Sprintf("%x", childKey.Key)
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter mnemonic (only if you have no crypto on it): ")
	mnemonic, _ := reader.ReadString('\n')
	mnemonic = mnemonic[:len(mnemonic)-1] // Remove newline character

	pubKey := PubKey(mnemonic, 0)
	address := crypto.PubkeyToAddress(pubKey)
	fmt.Printf("Ethereum Address: %s\n", address.Hex())
}
