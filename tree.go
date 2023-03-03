package main

import "log"


type Node struct {
  Key int
	Left  *Node
	Right *Node
  Value string
}


type OrganizationTree struct {
	root  *Node
  repo NodeRepository
}

func NewOrg(repo NodeRepository) Organization {
  return &OrganizationTree{
    repo: repo,
  }
}

type Organization interface {
  AddUser(parent *Node, user, side string) (*Node, error)
}

func insert(t *Node, v int) *Node {
	if t == nil {
		return &Node{0,nil, nil, ""}
	}
	if v == t.Key{
		return t
	}
	if v < t.Key {
		t.Left = insert(t.Left, v)
		return t
	}
	t.Right = insert(t.Right, v)
	return t
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
  node := &Node{
    Value: user,
  }
  var parent *Node
  if subtree != nil {
    parent = findParent(subtree, side)
  }else if o.hasRoot() {
    parent = findParent(o.root, side)
  } 
  

  id, err := o.repo.Save(parent, node)
  if err != nil {
    log.Println(err)
    return nil, err
  }

  node.Key = id

  return node, nil
}
 func (o *OrganizationTree) hasRoot() bool  {
  return o.root != nil
 }
