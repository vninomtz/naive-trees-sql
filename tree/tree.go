package tree

import (
	"fmt"
	"math"
)

type Node struct {
	Key   int
	Left  *Node
	Right *Node
	Value string
}

type Tree struct {
	root *Node
}

func (t *Tree) Preorder() {
	preOrder(t.root)
}
func height(n *Node) int {
	if n == nil {
		return -1
	}

	left := height(n.Left)
	right := height(n.Right)

	ans := math.Max(float64(left), float64(right))

	return int(ans)
}

func insert(n *Node, key int) *Node {
	if n == nil {
		return &Node{
			Key: key,
		}
	}
	if key < n.Key {
		n.Left = insert(n.Left, key)
	} else if key > n.Key {
		n.Right = insert(n.Right, key)
	}
	// Equal keys
	return n
}

func preOrder(n *Node) {
	// print: root-left-right
	if n == nil {
		return
	}
	fmt.Printf("n -> %s", n.Value)
	preOrder(n.Left)
	preOrder(n.Right)
}

func inOrder(n *Node) {
	// print: left-root-right
	if n == nil {
		return
	}
	inOrder(n.Left)
	fmt.Printf("n -> %s", n.Value)
	inOrder(n.Right)
}

func postOrder(n *Node) {
	// print: left-root-right
	if n == nil {
		return
	}
	postOrder(n.Left)
	postOrder(n.Right)
	fmt.Printf("n -> %s", n.Value)
}

func isFullTree(n *Node) bool {
	if n == nil {
		return true
	}

	if n.Left == nil && n.Right == nil {
		return true
	}

	if n.Left != nil && n.Right != nil {
		return isFullTree(n.Left) && isFullTree(n.Right)
	}

	return false
}
