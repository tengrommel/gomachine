package main

// B树的结点
type BTreeNode struct {
	leaf     bool         // 是否叶子
	N        int          // 分支的数量
	keys     []int        // 存储数据
	Children []*BTreeNode // 指向自己的多个分支节点
}

// 新建一个结点
func NewBTreeNode(n int, branch int, leaf bool) *BTreeNode {
	return &BTreeNode{leaf,
		n,
		make([]int, branch*2-1),
		make([]*BTreeNode, branch*2),
	}
}

// B树
type BTree struct {
	Root   *BTreeNode // 根节点
	branch int        // 分支的数量
}

// 搜索
func (bTreeNode *BTreeNode) Search(key int) (myNode *BTreeNode, idx int) {
	i := 0
	// 找到合适的位置，找到最后一个小于key的
	for i < bTreeNode.N && bTreeNode.keys[i] < key {
		i += 1
	}
	if i < bTreeNode.N && bTreeNode.keys[i] == key {
		myNode, idx = bTreeNode, i // 找到
	} else if bTreeNode.leaf == false {
		// 进入孩子叶子继续搜索
		myNode, idx = bTreeNode.Children[i].Search(key)
	}
	return
}
