package main

import (
	"fmt"
	"log"
)

type Node struct {
	Key    int
	Parent int
	Left   *Node
	Right  *Node
	Value  string
}

type OrganizationTree struct {
	root *Node
	repo NodeRepository
}

func NewOrg(repo NodeRepository) Organization {
	return &OrganizationTree{
		repo: repo,
	}
}

type Organization interface {
	AddUser(parent *Node, user, side string) (*Node, error)
	Print()
}

func findParent(root *Node, side string) *Node {
	if side == "left" {
		return findLeft(root)
	}

	return findRight(root)
}
func findRight(root *Node) *Node {
	if root.Right == nil {
		return root
	}

	return findRight(root.Right)
}

func findLeft(root *Node) *Node {
	if root.Left == nil {
		return root
	}

	return findLeft(root.Left)
}

func (o *OrganizationTree) AddUser(subtree *Node, user, side string) (*Node, error) {
	o.root = insert(o.root, user)

	n := get(o.root, user)
	_, err := o.repo.Save(o.root, n)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	//n.Key = id

	return n, nil
}
func (o *OrganizationTree) hasRoot() bool {
	return o.root != nil
}

func (o *OrganizationTree) Print() {
	printTree(o.root)
}

func printTree(n *Node) {
	if n == nil {
		return
	}
	printTree(n.Left)
	fmt.Printf("%s  ", n.Value)
	printTree(n.Right)
}

func insert(t *Node, user string) *Node {
	if t == nil {
		return &Node{0, 0, nil, nil, user}
	}
	if user == t.Value {
		return t
	}

	if t.Left == nil {
		t.Left = insert(t.Left, user)
	}
	if t.Right == nil {
		t.Right = insert(t.Right, user)
	}

	t.Left = insert(t.Left, user)
	return t
}

func get(x *Node, user string) *Node {
	if x == nil {
		return nil
	}
	if x.Value == user {
		return x
	}

	n := get(x.Left, user)

	if x.Value == user {
		return n
	}

	return get(x.Right, user)
}
