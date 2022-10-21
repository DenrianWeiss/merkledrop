package helpers

import (
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"math/big"
)

type InputAirdrop struct {
	Address string
	Amount  *big.Int
}

// Keccak256HashFunc return keccak256 of input.
func Keccak256HashFunc(input []byte) (hash [32]byte) {
	val := crypto.Keccak256(input)
	copy(hash[:], val)
	return
}

// EncodeForAirdrop encode address and amount to bytes for hashing.
func EncodeForAirdrop(address string, amount *big.Int) ([]byte, error) {
	Uint256, _ := abi.NewType("uint256", "", nil)
	Address, _ := abi.NewType("address", "", nil)
	args := abi.Arguments{
		{
			Type: Address,
		},
		{
			Type: Uint256,
		},
	}
	bytes, err := args.Pack(common.HexToAddress(address), amount)
	return bytes[12:], err
}

// EncoderFunc encode address and amount to bytes for hashing, provide interface.
func EncoderFunc(input interface{}) []byte {
	v, _ := EncodeForAirdrop(input.(InputAirdrop).Address, input.(InputAirdrop).Amount)
	return v
}
