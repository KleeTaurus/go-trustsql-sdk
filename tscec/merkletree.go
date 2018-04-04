package tscec

import (
	"bytes"
	"crypto/sha256"
	"errors"
	"fmt"
	"strings"
)

// Content MerkleTree 中叶子节点对应的存储内容
type Content interface {
	CalculateHash() []byte
	Equals(other Content) bool
}

// MerkleTree 用于存储哈希树的根节点、merkleRoot 值及所有叶子节点列表
type MerkleTree struct {
	Root       *Node
	merkleRoot []byte
	Leafs      []*Node
}

// Node MerkleTree 中的节点，包括：叶子节点、中间节点或根节点
type Node struct {
	Parent *Node
	Left   *Node
	Right  *Node
	leaf   bool
	dup    bool
	Hash   []byte
	C      Content
}

// String 对 Node 节点进行格式化输出
func (n Node) String() string {
	return fmt.Sprintf("Hash: %x, Leaf: %t, Value: %v", n.Hash, n.leaf, n.C)
}

// verifyNode 验证当前节点下所有子节点的哈希值
func (n *Node) verifyNode() []byte {
	if n.leaf {
		return n.C.CalculateHash()
	}
	h := sha256.New()
	h.Write(append(n.Left.verifyNode(), n.Right.verifyNode()...))
	return h.Sum(nil)
}

// calculateNodeHash 计算当前节点的哈希值
func (n *Node) calculateNodeHash() []byte {
	if n.leaf {
		return n.C.CalculateHash()
	}
	h := sha256.New()
	h.Write(append(n.Left.Hash, n.Right.Hash...))
	return h.Sum(nil)
}

// NewTree 使用 cs 构建一个 MerkleTree
func NewTree(cs []Content) (*MerkleTree, error) {
	root, leafs, err := buildWithContent(cs)
	if err != nil {
		return nil, err
	}
	t := &MerkleTree{
		Root:       root,
		merkleRoot: root.Hash,
		Leafs:      leafs,
	}
	return t, nil
}

// buildWithContent 根据传入的 cs 构建完整的 MerkleTree（cs 为空时返回错误）
func buildWithContent(cs []Content) (*Node, []*Node, error) {
	if len(cs) == 0 {
		return nil, nil, errors.New("Error: cannot construct tree with no content.")
	}
	var leafs []*Node
	for _, c := range cs {
		leafs = append(leafs, &Node{
			Hash: c.CalculateHash(),
			C:    c,
			leaf: true,
		})
	}
	if len(leafs)%2 == 1 {
		duplicate := &Node{
			Hash: leafs[len(leafs)-1].Hash,
			C:    leafs[len(leafs)-1].C,
			leaf: true,
			dup:  true,
		}
		leafs = append(leafs, duplicate)
	}
	root := buildIntermediate(leafs)
	return root, leafs, nil
}

// buildIntermediate 根据传入的叶子节点列表构建完整的 MerkleTree
func buildIntermediate(nl []*Node) *Node {
	var nodes []*Node
	for i := 0; i < len(nl); i += 2 {
		var left, right int = i, i + 1
		if i+1 == len(nl) {
			right = i
		}
		chash := append(nl[left].Hash, nl[right].Hash...)
		h := sha256.New()
		h.Write(chash)
		n := &Node{
			Left:  nl[left],
			Right: nl[right],
			Hash:  h.Sum(nil),
		}
		nodes = append(nodes, n)
		nl[left].Parent = n
		nl[right].Parent = n
		if len(nl) == 2 {
			return n
		}
	}
	return buildIntermediate(nodes)
}

// MerkleRoot 获取 MerkleTree 根节点哈希
func (m *MerkleTree) MerkleRoot() []byte {
	return m.merkleRoot
}

// RebuildTree 使用现有叶子节点重新构建 MerkleTree
func (m *MerkleTree) RebuildTree() error {
	var cs []Content
	for _, c := range m.Leafs {
		cs = append(cs, c.C)
	}
	root, leafs, err := buildWithContent(cs)
	if err != nil {
		return err
	}
	m.Root = root
	m.Leafs = leafs
	m.merkleRoot = root.Hash
	return nil
}

// RebuildTreeWith 使用传入的 cs 重新构建 MerkleTree
func (m *MerkleTree) RebuildTreeWith(cs []Content) error {
	root, leafs, err := buildWithContent(cs)
	if err != nil {
		return err
	}
	m.Root = root
	m.Leafs = leafs
	m.merkleRoot = root.Hash
	return nil
}

// VerifyTree 验证 MerkleTree 的根哈希是否与全部叶子节点计算出的 Merkle Root 一致
func (m *MerkleTree) VerifyTree() bool {
	calculatedMerkleRoot := m.Root.verifyNode()
	if bytes.Compare(m.merkleRoot, calculatedMerkleRoot) == 0 {
		return true
	}
	return false
}

//VerifyContent indicates whether a given content is in the tree and the hashes are valid for that content.
//Returns true if the expected Merkle Root is equivalent to the Merkle root calculated on the critical path
//for a given content. Returns true if valid and false otherwise.
func (m *MerkleTree) VerifyContent(expectedMerkleRoot []byte, content Content) bool {
	for _, l := range m.Leafs {
		if l.C.Equals(content) {
			currentParent := l.Parent
			for currentParent != nil {
				h := sha256.New()
				if currentParent.Left.leaf && currentParent.Right.leaf {
					h.Write(append(currentParent.Left.calculateNodeHash(), currentParent.Right.calculateNodeHash()...))
					if bytes.Compare(h.Sum(nil), currentParent.Hash) != 0 {
						return false
					}
					currentParent = currentParent.Parent
				} else {
					h.Write(append(currentParent.Left.calculateNodeHash(), currentParent.Right.calculateNodeHash()...))
					if bytes.Compare(h.Sum(nil), currentParent.Hash) != 0 {
						return false
					}
					currentParent = currentParent.Parent
				}
			}
			return true
		}
	}
	return false
}

// Proof 用于验证的结构体
type Proof struct {
	Position string
	Hash     []byte
}

// String 获取 Proof 的字符串表达式
func (p *Proof) String() string {
	return fmt.Sprintf("%s: %x", p.Position, p.Hash)
}

// getProof 获取当前节点对应层级之上的所有 Proof 路径
func getProof(t *Node, currentNodes []*Node) []*Proof {
	// fmt.Printf("\t%v\n", currentNodes)
	if len(currentNodes) < 2 {
		return []*Proof{}
	}

	parentNodes := []*Node{}
	for i, l := range currentNodes {
		if i%2 == 0 {
			parentNodes = append(parentNodes, l.Parent)
		}
	}

	var proof *Proof
	path := []*Proof{}
	for i, l := range currentNodes {
		if bytes.Equal(t.Hash, l.Hash) {
			if i%2 == 0 {
				if i == len(currentNodes)-1 {
					proof = &Proof{"Right", currentNodes[i].Hash[:]}
				} else {
					proof = &Proof{"Right", currentNodes[i+1].Hash[:]}
				}
			} else {
				proof = &Proof{"Left", currentNodes[i-1].Hash[:]}
			}

			path = append(path, proof)
			path = append(path, getProof(l.Parent, parentNodes)...)

			break
		}
	}

	return path
}

// GetProof 获取指定节点对应的 Proof 路径
func (m *MerkleTree) GetProof(content Content) ([]*Proof, error) {
	for _, l := range m.Leafs {
		if l.C.Equals(content) {
			return getProof(l, m.Leafs), nil
		}
	}
	return nil, errors.New("Error: cannot find the content in merkle tree.")
}

// String 获取 MerkleTree 的字符串表达式
func (m *MerkleTree) String() string {
	s := ""
	for _, l := range m.Leafs {
		s += fmt.Sprint(l)
		s += "\n"
	}
	return s
}

// LeafsHash 获取 MerkleTree 全部叶子节点哈希值拼接的字符串（哈希值采用 Base64 编码）
func (m *MerkleTree) LeafsHash() string {
	var h []string
	for _, l := range m.Leafs {
		if l.leaf == true && l.dup == false {
			h = append(h, string(Base58Encode(l.Hash)))
		}
	}
	return strings.Join(h, ",")
}
