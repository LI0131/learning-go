/*
 * This file defines a binary tree structure including functions for
 * updating and deleting elements in the tree.
 */

package main

import (
	"fmt"
)

type TreeNode struct {
	left, right *TreeNode
	value       int64
}

func (t *TreeNode) insert(item int64) {
	if t == nil {
		return
	} else if item <= t.value {
		if t.left == nil {
			t.left = &TreeNode{value: item}
		} else {
			t.left.insert(item)
		}
	} else {
		if t.right == nil {
			t.right = &TreeNode{value: item}
		} else {
			t.right.insert(item)
		}
	}
}

func delete(t *TreeNode, item int64) *TreeNode {
	if t == nil {
		return t
	}

	if item < t.value {
		t.left = delete(t.left, item)
		return t
	} else if item > t.value {
		t.right = delete(t.right, item)
		return t
	}

	if t.left == nil {
		temp := t.right
		t = nil
		return temp
	} else if t.right == nil {
		temp := t.left
		t = nil
		return temp
	} else {
		parent := t
		successor := t.right

		for successor.left != nil {
			parent = successor
			successor = successor.left
		}

		if parent != t {
			parent.left = successor.right
		} else {
			parent.right = successor.right
		}

		t.value = successor.value
		return t
	}
}

func (t *TreeNode) inorder() {
	if t.left != nil {
		t.left.inorder()
	}
	if t.value != 0 {
		fmt.Println(t.value)
	}
	if t.right != nil {
		t.right.inorder()
	}
}

func (t *TreeNode) postorder() {
	if t.left != nil {
		t.left.postorder()
	}
	if t.right != nil {
		t.right.postorder()
	}
	if t.value != 0 {
		fmt.Println(t.value)
	}
}

func (t *TreeNode) preorder() {
	if t.value != 0 {
		fmt.Println(t.value)
	}
	if t.left != nil {
		t.left.preorder()
	}
	if t.right != nil {
		t.right.preorder()
	}
}

func main() {
	root := TreeNode{}
	root.insert(1)
	root.insert(14)
	root.insert(2)
	root.insert(7)
	root.insert(28)
	root.insert(5)
	root.insert(9)
	root.inorder()
	root = *delete(&root, 7)
	fmt.Println()
	root.inorder()
}
