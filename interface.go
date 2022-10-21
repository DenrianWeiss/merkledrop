package merkledrop

// HashFunc hash for the merkle tree
type HashFunc func(input []byte) (hash [32]byte)

// EncodeFunc encode object
type EncodeFunc func(input interface{}) []byte
