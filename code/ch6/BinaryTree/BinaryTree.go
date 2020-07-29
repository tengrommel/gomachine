package main

import "fmt"

type Node struct {
	Data  int   // 数据
	Left  *Node // 指向左边节点
	Right *Node // 指向右边节点
}

type BinaryTree struct {
	Root *Node // 根节点
	Size int   // 数据的数量
}

// 新建一个二叉树
func NewBinaryTree() *BinaryTree {
	bst := &BinaryTree{}
	bst.Size = 0
	bst.Root = nil
	return bst
}

// 获取二叉树的大小
func (bst *BinaryTree) GetSize() int {
	return bst.Size
}

// 判断是否为空
func (bst *BinaryTree) IsEmpty() bool {
	return bst.Size == 0
}

// 根节点插入
func (bst *BinaryTree) Add(data int) {
	bst.Root = bst.add(bst.Root, data)
}

// 插入节点
func (bst *BinaryTree) add(n *Node, data int) *Node {
	if n == nil {
		bst.Size++
		//return &Node{Data: data}
		return &Node{data, nil, nil}
	} else {
		if data < n.Data {
			n.Left = bst.add(n.Left, data) // 比我小，左边
		} else if data > n.Data {
			n.Right = bst.add(n.Right, data) // 比我大，右边
		}
		return n
	}
}

// 判断数据是否存在
func (bst *BinaryTree) IsIn(data int) bool {
	return bst.isIn(bst.Root, data)
}

func (bst *BinaryTree) isIn(node *Node, data int) bool {
	if node == nil {
		return false // 树是空树，找不到
	}
	if data == node.Data {
		return true
	} else if data < node.Data {
		return bst.isIn(node.Left, data)
	} else {
		return bst.isIn(node.Right, data)
	}
}

func (bst *BinaryTree) FindMax() int {
	if bst.Size == 0 {
		panic("二叉树为空")
	}
	return bst.findMax(bst.Root).Data
}

func (bst *BinaryTree) findMax(node *Node) *Node {
	if node.Right == nil {
		return node
	} else {
		return bst.findMax(node.Right)
	}
}

func (bst *BinaryTree) FindMin() int {
	if bst.Size == 0 {
		panic("二叉树为空")
	}
	return bst.findMin(bst.Root).Data
}

func (bst *BinaryTree) findMin(node *Node) *Node {
	if node.Left == nil {
		return node
	} else {
		return bst.findMin(node.Left)
	}
}

// 前序遍历
func (bst *BinaryTree) PreOrder() {
	bst.preOrder(bst.Root)
}

func (bst *BinaryTree) preOrder(node *Node) {
	if node == nil {
		return
	}
	fmt.Println(node.Data)
	bst.preOrder(node.Left)
	bst.preOrder(node.Right)
}

// 中序遍历
func (bst *BinaryTree) InOrder() {
	bst.inOrder(bst.Root)
}

func (bst *BinaryTree) inOrder(node *Node) {
	if node == nil {
		return
	}
	bst.inOrder(node.Left)
	fmt.Println(node.Data)
	bst.inOrder(node.Right)
}

// 后序遍历
func (bst *BinaryTree) PostOrder() {
	bst.postOrder(bst.Root)
}

func (bst *BinaryTree) postOrder(node *Node) {
	if node == nil {
		return
	}
	bst.postOrder(node.Left)
	bst.postOrder(node.Right)
	fmt.Println(node.Data)
}
