package merkledrop

import (
	"bytes"
	"github.com/DenrianWeiss/merkledrop/helpers"
	"log"
	"math/big"
)

// GenericMerkleTree creates a merkle tree from a list of inputs using the provided hash and encode functions
func GenericMerkleTree(encode EncodeFunc, hash HashFunc, inputs []interface{}) (tree []interface{}) {
	tree = make([]interface{}, len(inputs)*2-1)
	for i := 0; i < len(inputs); i++ {
		hashResult := hash(encode(inputs[i]))
		tree[i] = hashResult[:]
	}
	for i := len(inputs); i < len(tree); i++ {
		left := tree[(i-len(inputs))*2].([]byte)
		right := tree[(i-len(inputs))*2+1].([]byte)
		var hashResult [32]byte
		if bytes.Compare(left, right) < 0 {
			hashResult = hash(append(left, right...))
		} else {
			hashResult = hash(append(right, left...))
		}
		tree[i] = hashResult[:]
	}
	return
}

// MerkleTreeProof creates a merkle tree proof for the given index
func MerkleTreeProof(tree []interface{}, index int) (proof []interface{}) {
	leafPow := helpers.Log2(len(tree))
	acc := 0
	proof = append(proof, tree[helpers.FindPairIndex(index)])
	for leafPow > 1 {
		index = index / 2
		acc += 1 << leafPow
		log.Printf("index: %d, acc: %d, leafPow: %d", index, acc, leafPow)
		proof = append(proof, tree[helpers.FindPairIndex(index+acc)])
		leafPow -= 1
	}
	return proof
}

// PaddleTreeTo2n paddles the tree to the nearest 2^n
func PaddleTreeTo2n(tree []helpers.InputAirdrop) (paddedTree []helpers.InputAirdrop) {
	paddedTree = make([]helpers.InputAirdrop, helpers.NextPowerOfTwo(len(tree)))
	for i := 0; i < len(tree); i++ {
		paddedTree[i] = tree[i]
	}
	for i := len(tree); i < len(paddedTree); i++ {
		paddedTree[i] = helpers.InputAirdrop{
			Address: "0x0000000000000000000000000000000000000000",
			Amount:  big.NewInt(0),
		}
	}
	return
}

// AirdropMerkleTree creates a merkle tree for the airdrop special use case.
func AirdropMerkleTree(inputsOrig []helpers.InputAirdrop) (tree []interface{}) {

	inputsP := PaddleTreeTo2n(inputsOrig)
	inputs := ConvertArrayToInterface(inputsP)
	return GenericMerkleTree(helpers.EncoderFunc, helpers.Keccak256HashFunc, inputs)
}

func ConvertArrayToInterface(input []helpers.InputAirdrop) (output []interface{}) {
	output = make([]interface{}, len(input))
	for i := 0; i < len(input); i++ {
		output[i] = input[i]
	}
	return
}

// CreateAirdropTree creates a merkle tree and generate root and proofs.
// Proof for certain input is in the same index of proof
func CreateAirdropTree(inputs []helpers.InputAirdrop) (root []byte, proof [][]string) {
	// Generate merkle tree
	tree := AirdropMerkleTree(inputs)
	// Generate proofs
	for i := 0; i < len(inputs); i++ {
		proof = append(proof, helpers.ProofListToArray(MerkleTreeProof(tree, i)))
	}
	root = tree[len(tree)-1].([]byte)
	return
}
