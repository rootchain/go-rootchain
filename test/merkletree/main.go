package main

import (
	"crypto/sha256"
	"log"

	"github.com/NebulousLabs/merkletree"
)

// All error checking is ignored in the following examples.
func main() {
	log.Printf("fullProof: %v", (1 << 32))

	// Example 3: Using a Tree to build a merkle tree and get a proof for a
	// specific index for non-file objects.
	tree := merkletree.New(sha256.New())
	tree.SetIndex(2)
	tree.Push([]byte("an object - the tree will hash the data after it is pushed"))
	tree.Push([]byte("another object"))
	tree.Push([]byte("another object"))
	tree.Push([]byte("another object"))
	// The merkle root could be obtained by calling tree.Root(), but will also
	// be provided by tree.Prove()
	merkleRoot, _, _, _ := tree.Prove()
	log.Printf("merkleRoot: %x", merkleRoot)

	////////////////////////////////////////////////
	/// Remaining examples deal with cached trees //
	////////////////////////////////////////////////

	// Example 4: Creating a cached set of Merkle roots and then using them in
	// a cached tree. The cached tree is height 1, meaning that all elements of
	// the cached tree will be Merkle roots of data with 2 leaves.
	cachedTree := merkletree.NewCachedTree(sha256.New(), 32)
	subtree1 := merkletree.New(sha256.New())
	subtree1.Push([]byte("first leaf, first subtree"))
	subtree1.Push([]byte("second leaf, first subtree"))
	subtree2 := merkletree.New(sha256.New())
	subtree2.Push([]byte("first leaf, second subtree"))
	subtree2.Push([]byte("second leaf, second subtree"))
	// Using the cached tree, build the merkle root of the 4 leaves.
	cachedTree.Push(subtree1.Root())
	cachedTree.Push(subtree2.Root())
	collectiveRoot := cachedTree.Root()
	log.Printf("collectiveRoot: %x", collectiveRoot)

	// Example 5: Modify the data pushed into subtree 2 and create the Merkle
	// root, without needing to rehash the data in any other subtree.
	revisedSubtree2 := merkletree.New(sha256.New())
	revisedSubtree2.Push([]byte("first leaf, second subtree"))
	revisedSubtree2.Push([]byte("second leaf, second subtree, revised"))
	// Using the cached tree, build the merkle root of the 4 leaves - without
	// needing to rehash any of the data in subtree1.
	cachedTree = merkletree.NewCachedTree(sha256.New(), 32)
	cachedTree.Push(subtree1.Root())
	cachedTree.Push(revisedSubtree2.Root())
	revisedRoot := cachedTree.Root()
	log.Printf("revisedRoot:    %x", revisedRoot)

	// Exapmle 6: Create a proof that leaf 3 (index 2) of the revised root,
	// found in revisedSubtree2 (at index 0 of the revised subtree), is a part of
	// the cached set. This is a two stage process - first we must get a proof
	// that the leaf is a part of revisedSubtree2, and then we must get provide
	// that proof as input to the cached tree prover.
	cachedTree = merkletree.NewCachedTree(sha256.New(), 32)
	cachedTree.SetIndex(2) // leaf at index 2, or the third element which gets inserted.
	revisedSubtree2 = merkletree.New(sha256.New())
	revisedSubtree2.SetIndex(0)
	revisedSubtree2.Push([]byte("first leaf, second subtree"))
	revisedSubtree2.Push([]byte("second leaf, second subtree, revised"))
	_, subtreeProof, _, _ := revisedSubtree2.Prove()
	// Now we can create the full proof for the cached tree, without having to
	// rehash any of the elements from subtree1.
	_, fullProof, _, _ := cachedTree.Prove(subtreeProof)
	log.Printf("fullProof: %v", fullProof)
}
