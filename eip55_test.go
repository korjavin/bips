package main

import (
	"testing"
)

func TestToChecksumAddress(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"0xC49fb5CF14b357f993bA5FA76BE3Dd438177d5d3", "0xC49fb5CF14b357f993bA5FA76BE3Dd438177d5d3"},
		{"0x52908400098527886E0F7030069857D2E4169EE7", "0x52908400098527886E0F7030069857D2E4169EE7"},
		{"0x8617E340B3D01FA5F11F306F4090FD50E238070D", "0x8617E340B3D01FA5F11F306F4090FD50E238070D"},
		{"0xde709f2102306220921060314715629080e2fb77", "0xde709f2102306220921060314715629080e2fb77"},
		{"0x27b1fdb04752bbc536007a920d24acb045561c26", "0x27b1fdb04752bbc536007a920d24acb045561c26"},
		{"0x5A1b6eD8eaf8e6a5f6D2e3aD6eD8eaf8e6a5f6D2", "0x5A1b6eD8eaf8e6a5f6D2e3aD6eD8eaf8e6a5f6D2"},
	}

	for _, test := range tests {
		result := ToChecksumAddress(test.input)
		if result != test.expected {
			t.Errorf("For input %s, expected %s but got %s", test.input, test.expected, result)
		}
	}
}
