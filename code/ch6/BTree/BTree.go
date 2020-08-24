package main

import "fmt"

// B树的结点
type BTreeNode struct {
	Leaf     bool         // 是否叶子
	N        int          // 分支的数量
	keys     []int        // 存储数据
	Children []*BTreeNode // 指向自己的多个分支节点
}

// 新建一个结点
func NewBTreeNode(n int, branch int, leaf bool) *BTreeNode {
	return &BTreeNode{
		leaf,
		n,
		make([]int, branch*2-1), // n个branch对应2n 去掉root 2n-1
		make([]*BTreeNode, branch*2),
	}
}

// 切割
func (parent *BTreeNode) Split(branch int, idx int) {
	full := parent.Children[idx]
	// 孩子节点
	newNode := NewBTreeNode(branch-1, branch, full.Leaf)
	// 新建一个节点，备份
	for i := 0; i < branch-1; i++ {
		newNode.keys[i] = full.keys[i+branch]
		// 数据移动，跳过一个分支
		newNode.Children[i] = full.Children[i+branch]
	}
	newNode.Children[branch-1] = full.Children[2*branch-1]
	// 处理最后
	full.N = branch - 1
	// 新增一个key到children
	for i := parent.N; i > idx; i-- {
		parent.Children[i] = parent.Children[i-1]
		parent.keys[i+1] = parent.keys[i]
	}
	parent.keys[idx] = full.keys[branch-1]
	parent.Children[idx+1] = newNode
	parent.N++
}

// 节点插入数据
func (bTreeNode BTreeNode) InsertNonFull(branch int, key int) {
	i := bTreeNode.N    // 记录叶子节点的总量
	if bTreeNode.Leaf { // 是叶子或者不是叶子
		for i > 0 && key < bTreeNode.keys[i-1] {
			bTreeNode.keys[i] = bTreeNode.keys[i-1] // 从后往前移动
			i--                                     // i从后往前移动
		}
		bTreeNode.keys[i] = key // 插入数据
		bTreeNode.N++           // 总量加一
	} else {
		for i > 0 && key < bTreeNode.keys[i-1] {
			i-- // i从后往前移动
		}
		c := bTreeNode.Children[i] // 找到下标
		if c.N == 2*branch-1 {
			bTreeNode.Split(branch, i) // 切割
			if key > bTreeNode.keys[i] {
				i++
			}
		}
		bTreeNode.Children[i].InsertNonFull(branch, key) // 递归插入到子叶子
	}
}

// 查找一个数据是否在一个节点上存在
func (bTreeNode *BTreeNode) Search(key int) (myNode *BTreeNode, idx int) {
	i := 0
	// 找到合适的位置，找到最后一个小于key的，那么i之后的就是大于等于
	for i < bTreeNode.N && bTreeNode.keys[i] < key {
		i += 1
	}
	if i < bTreeNode.N && bTreeNode.keys[i] == key {
		myNode, idx = bTreeNode, i // 找到
	} else if bTreeNode.Leaf == false {
		// 进入孩子叶子继续搜索
		myNode, idx = bTreeNode.Children[i].Search(key) // 递归搜索
	}
	return
}

// 将节点显示为字符串
func (bTreeNode *BTreeNode) String() string {
	return fmt.Sprintf("{n=%d, leaf=%v, Children=%v}\n", bTreeNode.N, bTreeNode.keys, bTreeNode.Children)
}

// B树
type BTree struct {
	Root   *BTreeNode // 根节点
	branch int        // 分支的数量
}

func (tree *BTree) Insert(key int) {
	root := tree.Root // 根节点
	if root.N == 2*tree.branch-1 {
		s := NewBTreeNode(0, tree.branch, false)
		tree.Root = s // 新建一个节点备份根节点
		s.Children[0] = root
		s.Split(tree.branch, 0) // 数据拆分并整合
		root.InsertNonFull(tree.branch, key)
	} else {
		root.InsertNonFull(tree.branch, key)
	}
}

func (tree *BTree) Search(key int) (n *BTreeNode, idx int) {
	return tree.Root.Search(key)
}

func (tree *BTree) String() string {
	return tree.Root.String()
}

func NewBTree(branch int) *BTree {
	return &BTree{NewBTreeNode(0, branch, true), branch}
}
