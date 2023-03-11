package tree

import (
	"fmt"
)

type Node struct {
	Key      int     `json:"key"`
	Parent   int     `json:"parent"`
	Left     *Node   `json:"left"`
	Right    *Node   `json:"right"`
	Value    string  `json:"value"`
	Children []*Node `json:"children"`
}

type Tree struct {
	Root *Node
}

func NewLevelOrderTree(nds []*Node) *Tree {
	root := insertLevelOrder(nds, 0, nil)

	return &Tree{
		Root: root,
	}
}

func (t *Tree) Height() int {
	return height(t.Root)
}

func (t *Tree) PrintLevelOrder() {
	h := height(t.Root)

	for i := 1; i <= h; i++ {
		printCurrentLevel(t.Root, i)
		fmt.Println()
	}
}

func printCurrentLevel(n *Node, level int) {
	if n == nil {
		return
	}
	if level == 1 {
		fmt.Printf("(P=%d, V=%s)  ",n.Parent, n.Value)
	} else if level > 1 {
		printCurrentLevel(n.Left, level-1)
		printCurrentLevel(n.Right, level-1)
	}

}
func Inorder(r *Node) {
	inOrder(r)
}
func height(n *Node) int {
	if n == nil {
		return 0
	}

	left := height(n.Left)
	right := height(n.Right)

	if left > right {
		return (left + 1)
	} else {
		return (right + 1)
	}
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
	fmt.Printf("%s ->", n.Value)
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

func insertLevelOrder(nds []*Node, i int, parent *Node) *Node {
	var root *Node
	if i < len(nds) {
		root = &Node{
			Key:   nds[i].Key,
			Value: nds[i].Value,
      Children: []*Node{},
		}
    if parent != nil {
      root.Parent = parent.Key
      nds[i].Parent = parent.Key
    }
		root.Left = insertLevelOrder(nds, 2*i+1, root)
		if root.Left != nil {
			root.Children = append(root.Children, root.Left)
		}
		root.Right = insertLevelOrder(nds, 2*i+2, root)
		if root.Right != nil {
			root.Children = append(root.Children, root.Right)
		}
	}

	return root
}
