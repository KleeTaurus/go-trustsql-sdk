package tscec

import (
	"crypto/sha256"
	"fmt"
)

func ComputeMerkleRootFromPath(leafHash string, merklePath []string, leafIndex uint32) string {
	var hash string = leafHash
	for _, p := range merklePath {
		if (leafIndex & uint32(1)) != 0 {
			hash = combineHash(p, hash)
		} else {
			hash = combineHash(hash, p)
		}
		leafIndex >>= 1
	}
	return hash
}

func ComputeMerkleRoot(leaves []string) string {
	var merkleRoot string
	merkleComputation(leaves, &merkleRoot, 0, nil)
	return merkleRoot
}

func ComputeMerklePath(leaves []string, leafIndex uint32) []string {
	merklePath := []string{}
	merkleComputation(leaves, nil, leafIndex, &merklePath)
	return merklePath
}

// Ported from consensus/merkle::MerkleComputation in Bitcoin Core
func merkleComputation(leaves []string, pRoot *string, leafIndex uint32, pPath *[]string) {
	if len(leaves) == 0 {
		return
	}

	// the number of leaves processed so far.
	var count uint32 = 0
	// inner is an array of eagerly computed subtree hashes, indexed by tree
	// level (0 being the leaves).
	// For example, when count is 25 (11001 in binary), inner[4] is the hash of
	// the first 16 leaves, inner[3] of the next 8 leaves, and inner[0] equal to
	// the last leaf. The other inner entries are undefined.
	var inner [32]string
	var level uint32
	var h string
	var matchLevel int = -1

	// First process all leaves into 'inner' values.
	for count < uint32(len(leaves)) {
		h = leaves[count]
		var matchh bool = (count == leafIndex)
		count++
		// For each of the lower bits in count that are 0, do 1 step. Each
		// corresponds to an inner value that existed before processing the
		// current leaf, and each needs a hash to combine it.
		for level = 0; (count & ((uint32(1)) << level)) == 0; level++ {
			condition := (count & ((uint32(1)) << level)) == 0
			fmt.Printf("count: %d, level: %d, condition: %t\n", count, level, condition)
			if pPath != nil {
				if matchh {
					*pPath = append(*pPath, inner[level])
				} else if matchLevel == int(level) {
					*pPath = append(*pPath, h)
					matchh = true
				}
			}
			h = combineHash(inner[level], h)
		}
		inner[level] = h
		if matchh {
			matchLevel = int(level)
		}
	}

	// Do a final 'sweep' over the rightmost branch of the tree to process
	// odd levels, and reduce everything to a single top value.
	// Level is the level (counted from the bottom) up to which we've sweeped.
	level = 0
	// As long as bit number level in count is zero, skip it. It means there
	// is nothing left at this level.
	for (count & (uint32(1) << level)) == 0 {
		level++
	}
	h = inner[level]
	var matchh bool = (matchLevel == int(level))
	for count != (uint32(1) << level) {
		// If we reach this point, h is an inner value that is not the top.
		// We combine it with itself (Bitcoin's special rule for odd levels in
		// the tree) to produce a higher level one.
		if pPath != nil && matchh {
			*pPath = append(*pPath, h)
		}
		h = combineHash(h, h)
		// Increment count to the value it would have if two entries at this
		// level had existed
		count += (uint32(1) << level)
		level++
		// And propagate the result upwards accordingly.
		for (count & (uint32(1) << level)) == 0 {
			if pPath != nil {
				if matchh {
					*pPath = append(*pPath, inner[level])
				} else if matchLevel == int(level) {
					*pPath = append(*pPath, h)
					matchh = true
				}
			}
			h = combineHash(inner[level], h)
			level++
		}
	}

	// Return result.
	if pRoot != nil {
		*pRoot = h
	}
}

func dsha256(s string) string {
	firstHash := sha256.Sum256([]byte(s))
	// Convert array to slice: firstHash[:]
	secondHash := sha256.Sum256(firstHash[:])
	// convert byte array to str
	ret := fmt.Sprintf("%s", secondHash)
	return ret
}

func combineHash(x, y string) string {
	cat := x + y
	return dsha256(cat)
}
