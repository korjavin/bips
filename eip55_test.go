package main

import (
	"errors"
	"fmt"
	"strings"
	"testing"

	"crypto/rand"
	"encoding/hex"

	"github.com/ethereum/go-ethereum/common"
)

// GenerateRandomAddress generates a random Ethereum address
func GenerateRandomAddress() string {
	bytes := make([]byte, 20)
	_, err := rand.Read(bytes)
	if err != nil {
		panic(err)
	}
	return "0x" + hex.EncodeToString(bytes)
}

// compareCase compares the case of characters in two strings and returns an error if they differ
func compareCase(str1, str2 string) error {
	if len(str1) != len(str2) {
		return errors.New("strings are of different lengths")
	}

	for i := 0; i < len(str1); i++ {
		if str1[i] != str2[i] && strings.ToLower(string(str1[i])) == strings.ToLower(string(str2[i])) {
			return fmt.Errorf("%s*%s* is different", string(str1[:i]), string(str2[i]))
		}
	}

	return nil
}

func TestToChecksumAddress(t *testing.T) {
	for i := 0; i < 100; i++ {
		randomAddress := GenerateRandomAddress()
		byteSlice, _ := hex.DecodeString(randomAddress[2:])
		expected := common.BytesToAddress(byteSlice).Hex()
		result := ToChecksumAddress(randomAddress)
		if result != expected {
			t.Errorf("For input %s and result %s, %v", expected, result, compareCase(expected, result))
		}
	}
}
