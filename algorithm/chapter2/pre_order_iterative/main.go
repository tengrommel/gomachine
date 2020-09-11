package main

import (
	"fmt"
	"math/rand"
)

type BinaryTreeNode struct {
	left  *BinaryTreeNode
	data  int
	right *BinaryTreeNode
}

// NewBinaryTree returns a new, random binary tree holding the values 1k, 2k, ..., nk
func NewBinaryTree(n, k int) *BinaryTreeNode {
	var root *BinaryTreeNode
	for _, v := range rand.Perm(n) {
		root = insert(root, (1+v)*k)
	}
	return root
}

func insert(root *BinaryTreeNode, v int) *BinaryTreeNode {
	if root == nil {
		return &BinaryTreeNode{nil, v, nil}
	}
	if v < root.data {
		root.left = insert(root.left, v)
		return root
	}
	root.right = insert(root.right, v)
	return root
}

type Stack struct {
	top      int
	capacity uint
	array    []interface{}
}

// Returns an initialized stack
func (stack *Stack) Init(capacity uint) *Stack {
	stack.top = -1
	stack.capacity = capacity
	stack.array = make([]interface{}, capacity)
	return stack
}

// Returns an new stack
func NewStack(capacity uint) *Stack {
	return new(Stack).Init(capacity)
}

// Returns the size of the size
func (stack *Stack) Size() uint {
	return uint(stack.top + 1)
}

// Stack is full when top is equal to the last index
func (stack *Stack) IsFull() bool {
	return stack.top == int(stack.capacity)-1
}

// Stack is empty when top is equal to -1
func (stack *Stack) IsEmpty() bool {
	return stack.top == -1
}

func (stack *Stack) Resize() {
	if stack.IsFull() {
		stack.capacity *= 2
	} else {
		stack.capacity /= 2
	}

	target := make([]interface{}, stack.capacity)
	copy(target, stack.array[:stack.top+1])
	stack.array = target
}

func (stack *Stack) Push(data interface{}) error {
	if stack.IsFull() {
		stack.Resize()
	}
	stack.top++
	stack.array[stack.top] = data
	return nil
}

func (stack *Stack) Pop() interface{} {
	var i interface{}
	if stack.IsEmpty() {
		return i
	}
	temp := stack.array[stack.top]
	stack.top--
	if stack.Size() < stack.capacity/2 {
		stack.Resize()
	}
	return temp
}

func (stack *Stack) Peek() interface{} {
	var i interface{}
	if stack.IsEmpty() {
		return i
	}
	temp := stack.array[stack.top]
	return temp
}

// Drain removes all elements that are currently in the stack.
func (stack *Stack) Drain() {
	stack.array = nil
	stack.top = -1
}

func PreOrder(root *BinaryTreeNode) {
	if root == nil {
		return
	}
	current := root
	stack := NewStack(1)
	stack.Push(current)
	for stack.Size() > 0 {
		temp := stack.Pop()
		current = temp.(*BinaryTreeNode)
		fmt.Printf("%d ", current.data)
		if current.right != nil {
			stack.Push(current.right)
		}
		if current.left != nil {
			stack.Push(current.left)
		}
	}
}

func main() {
	t1 := NewBinaryTree(10, 1)

	// PreOrder walk with print statements
	PreOrder(t1)
}
