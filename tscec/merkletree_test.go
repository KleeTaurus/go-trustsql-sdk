package tscec

import (
	"testing"
)

func TestCalcMerkleRoot(t *testing.T) {
	leafHashes := []string{}
	leaves := []string{
		"1",
		"2",
		"3",
	}
	// test for calc merkle root
	for _, l := range leaves {
		leafHashes = append(leafHashes, dsha256(l))
	}
	var root string
	root = ComputeMerkleRoot(leafHashes)

	h1 := combineHash(leafHashes[0], leafHashes[1])
	h2 := combineHash(leafHashes[2], leafHashes[2])
	h := combineHash(h1, h2)
	if root != h {
		t.Errorf("calc merkle root failed, got: %x, want: %x", root, h)
	}

	// test for calc merkle path && validate through merkle path
	for pos := uint32(0); pos <= 2; pos++ {
		var merklePath []string
		merklePath = ComputeMerklePath(leafHashes, pos)
		t.Logf("merkle path: \n")
		for i, p := range merklePath {
			t.Logf("%d\t%x\n", i, p)
		}

		merkleRoot := ComputeMerkleRootFromPath(leafHashes[pos], merklePath, pos)

		if merkleRoot != root {
			t.Errorf("validate failed, got: %x, want: %x", merkleRoot, root)
		}
	}
}
