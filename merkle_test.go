package merkledrop

import (
	"github.com/DenrianWeiss/merkledrop/helpers"
	"math/big"
	"testing"
)

var airdropList = []helpers.InputAirdrop{
	{
		Address: "0x1000000000000000000000000000000000000001",
		Amount:  big.NewInt(0x101),
	},
	{
		Address: "0x1000000000000000000000000000000000000002",
		Amount:  big.NewInt(0x102),
	},
	{
		Address: "0x1000000000000000000000000000000000000003",
		Amount:  big.NewInt(0x103),
	},
}

func TestPaddleTreeTo2n(t *testing.T) {
	tree := PaddleTreeTo2n(airdropList)
	for i, _ := range tree {
		t.Logf("Account: %s", tree[i])
	}
}

func TestHashFunc(t *testing.T) {
	hash := helpers.Keccak256HashFunc([]byte("123"))
	t.Logf("Hash: %v", hash)
}

func TestEncode(t *testing.T) {
	encoded := helpers.EncoderFunc(airdropList[0])
	t.Logf("Encoded: %x", encoded)
}

func TestGeneratedAirdrop(t *testing.T) {

	tree := AirdropMerkleTree(airdropList)
	for i, _ := range tree {
		t.Logf("Account: %x", tree[i])
	}
	for i, _ := range airdropList {
		t.Logf("Account: %s, Hash: %x, Proof: %v", airdropList[i], helpers.Keccak256HashFunc(helpers.EncoderFunc(airdropList[i])), helpers.ProofListToArray(MerkleTreeProof(tree, i)))
	}
}

func TestTreeResult(t *testing.T) {
	root, proof := CreateAirdropTree(airdropList)
	t.Logf("Root: %x", root)
	for i, airdrop := range airdropList {
		t.Logf("Account: %s, Amount %s, Hash: %x, Proof: %v", airdrop.Address, airdrop.Amount.String(), helpers.Keccak256HashFunc(helpers.EncoderFunc(airdrop)), proof[i])
	}
}
