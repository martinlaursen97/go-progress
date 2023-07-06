package main

///////////////////////
// Binary Tree in Go //
///////////////////////

import "fmt"

// Comparable interface
type Comparable interface {
	LessThan(other Comparable) bool
	GreaterThan(other Comparable) bool
}

// Comparable type
type Int int

func (i Int) LessThan(other Comparable) bool {
	return i < other.(Int)
}

func (i Int) GreaterThan(other Comparable) bool {
	return i > other.(Int)
}

// Tree code
type TreeNode[T Comparable] struct {
	Value T
	Left  *TreeNode[T]
	Right *TreeNode[T]
}

func (tn *TreeNode[T]) InorderTraversal() {
	if tn.Left != nil {
		tn.Left.InorderTraversal()
	}
	fmt.Println(tn.Value)
	if tn.Right != nil {
		tn.Right.InorderTraversal()
	}
}

func (tn *TreeNode[T]) Find(value T) bool {
	if value.LessThan(tn.Value) {
		if tn.Left == nil {
			return false
		}
		return tn.Left.Find(value)
	}

	if value.GreaterThan(tn.Value) {
		if tn.Right == nil {
			return false
		}
		return tn.Right.Find(value)
	}

	return true
}

func (tn *TreeNode[T]) InsertChild(value T, child **TreeNode[T]) {
	if *child == nil {
		*child = &TreeNode[T]{Value: value}
	} else {
		(*child).Insert(value)
	}
}

func (tn *TreeNode[T]) Insert(value T) {
	if value.LessThan(tn.Value) {
		tn.InsertChild(value, &tn.Left)
	} else {
		tn.InsertChild(value, &tn.Right)
	}
}

func main() {
	rootN := TreeNode[Int]{Value: 2}

	rootN.Insert(2)
	rootN.Insert(5)
	rootN.Insert(22)
	rootN.Insert(1)
	rootN.Insert(0)
	rootN.Insert(5)
	rootN.Insert(11)
	rootN.Insert(8)
	rootN.Insert(72)
	rootN.Insert(5)
	rootN.Insert(9)

	rootN.InorderTraversal()
	fmt.Println()
	exists := rootN.Find(2)

	fmt.Println(exists)

}
