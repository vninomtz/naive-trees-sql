# Tree notes

## AVL
- AVL tree is a self-balancing Binary Search Tree (BST)
- The difference between heights of left and right subtrees cannot be more than once for all nodes
- The height of an AVL tree is always O(Logn) where n is the number of nodes in the tree

### Insertion

- First perform standard BST Insertion of the node and then rebalance the BST by left or right rotation.
- After insertion if the tree is not balanced then following 4 cases might occur:
  - Left Left 
  - Left Right
  - Right Right
  - Right Left
