package tscec

import (
	"crypto/sha256"
	"fmt"
	"testing"
)

type TestContent struct {
	x string
}

func (t TestContent) CalculateHash() []byte {
	h := sha256.New()
	h.Write([]byte(t.x))
	return h.Sum(nil)
}

func (t TestContent) Equals(other Content) bool {
	return t.x == other.(TestContent).x
}

func TestMerkleTree(t *testing.T) {
	var list []Content
	list = append(list, TestContent{x: "1"})
	list = append(list, TestContent{x: "2"})
	list = append(list, TestContent{x: "3"})
	list = append(list, TestContent{x: "4"})
	list = append(list, TestContent{x: "5"})

	tree, _ := NewTree(list)
	mr := tree.MerkleRoot()
	fmt.Printf("Merkle Root: %x\n", mr)

	vt := tree.VerifyTree()
	fmt.Println("Verify Tree:", vt)

	vc := tree.VerifyContent(mr, list[0])
	fmt.Println("Verify Content:", vc)

	fmt.Println(tree)
	fmt.Println(tree.LeafsHash())
}
