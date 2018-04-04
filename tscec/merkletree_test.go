package tscec

import (
	"crypto/sha256"
	"fmt"
	"testing"
)

var DEBUG bool = true

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
	fmt.Println("======test merkletree begin======")

	var list []Content
	list = append(list, TestContent{x: "1"})
	list = append(list, TestContent{x: "2"})
	list = append(list, TestContent{x: "3"})
	list = append(list, TestContent{x: "4"})
	list = append(list, TestContent{x: "5"})

	tree, _ := NewTree(list)
	mr := tree.MerkleRoot()
	debug(fmt.Sprintf("Merkle Root: %x", mr))

	if len(tree.Leafs) != 6 {
		t.Errorf("The amount of leafs should be 6, %d\n", len(tree.Leafs))
	}

	vt := tree.VerifyTree()
	if !vt {
		t.Errorf("Verify tree failed\n")
	}

	vc := tree.VerifyContent(mr, list[0])
	if !vc {
		t.Errorf("Verify content failed, content: %v\n", list[0])
	}

	debug(fmt.Sprint(tree))
	debug(fmt.Sprintf("All leafs hash with base64: %s", tree.LeafsHash()))

	proof, err := tree.GetProof(list[4])
	if err != nil {
		t.Errorf("Get proof path failed %s\n", err)
	}
	debug(fmt.Sprintf("Proof: %v", proof))

	fmt.Println("======test merkletree end======")
}

func debug(msg string) {
	if DEBUG {
		fmt.Println(msg)
	}
}
