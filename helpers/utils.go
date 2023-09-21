package helpers

import (
	"fmt"
	"log"
	"math/bits"
)

// PrintBytesArray prints array of bytes in hex format.
func PrintBytesArray(input []interface{}) {
	for _, v := range input {
		log.Printf("%x", v.([]byte))
	}
}

// NextPowerOfTwo returns the next power of 2 of x.
func NextPowerOfTwo(x int) int {
	if x == 0 {
		return 1
	}
	return 1 << (32 - bits.LeadingZeros32(uint32(x-1)))
}

// Log2 returns the log2 of x in int.
func Log2(x int) int {
	return bits.Len(uint(x)) - 1
}

func BytesToHexString(input []byte) string {
	b32 := [32]byte{}
	copy(b32[:], input)
	return fmt.Sprintf("0x%x", input)
}

// ProofListToArray convert proof to array of bytes.
func ProofListToArray(proof []interface{}) []string {
	output := make([]string, len(proof))
	for i := 0; i < len(proof); i++ {
		output[i] = BytesToHexString(proof[i].([]byte))
	}
	return output
}

// FindPairIndex returns the index of the pair in binary heap of the given index.
func FindPairIndex(index int) int {
	if index%2 == 0 {
		return index + 1
	}
	return index - 1
}

func IsAllZero(b []byte) bool {
	for _, v := range b {
		if v != 0 {
			return false
		}
	}
	return true
}
