package main

import "errors"

// 适用于没有删除的情况
// 红黑树的增删查改是最优的
type AVLNode struct {
	Data   interface{} // 数据
	Left   *AVLNode    // 指向左边节点
	Right  *AVLNode    // 指向右边节点
	height int         // 高度
}

// comparator函数指针
type comparator func(a, b interface{}) int

// compare函数指针
var compare comparator

func Max(a, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}
}

func NewNode(data interface{}) *AVLNode {
	node := new(AVLNode) // 新建节点
	node.Data = data
	node.Left = nil
	node.Right = nil
	node.height = 0
	return node
}

// 新建AVLTree
func NewAVLTree(data interface{}, myFunc comparator) (*AVLNode, error) {
	if data == nil && myFunc == nil {
		return nil, errors.New("不能为空")
	}
	compare = myFunc
	return NewNode(data), nil
}

// 左旋，逆时针
func (avlNode *AVLNode) LeftRotate() *AVLNode {
	headNode := avlNode.Right
	avlNode.Right = headNode.Left
	avlNode.Left = avlNode
	// 更新高度
	avlNode.height = Max(avlNode.Left.GetHeight(), avlNode.Right.GetHeight()) + 1
	headNode.height = Max(headNode.Left.GetHeight(), headNode.Right.GetHeight()) + 1
	return headNode
}

// 右旋，顺时针
func (avlNode *AVLNode) RightRotate() *AVLNode {
	headNode := avlNode.Left      // 保存左边节点
	avlNode.Left = headNode.Right // 设为nil
	headNode.Right = avlNode
	// 更新高度
	avlNode.height = Max(avlNode.Left.GetHeight(), avlNode.Right.GetHeight()) + 1
	headNode.height = Max(headNode.Left.GetHeight(), headNode.Right.GetHeight()) + 1
	return headNode
}

// 两次左旋
// 两次右旋
// 先左旋再右旋
func (avlNode *AVLNode) LeftThenRightRotate() *AVLNode {
	sonHeadNode := avlNode.Left.LeftRotate()
	avlNode.Left = sonHeadNode
	return avlNode.RightRotate()
}

// 先右旋再左旋
func (avlNode *AVLNode) RightThenLeftRotate() *AVLNode {
	sonHeadNode := avlNode.Right.RightRotate()
	avlNode.Right = sonHeadNode
	return avlNode.LeftRotate()
}

// 自动处理不平衡，差距为1平衡，差距为2不平衡
func (avlNode *AVLNode) adjust() *AVLNode {
	if avlNode.Right.GetHeight()-avlNode.Left.GetHeight() == 2 {
		if avlNode.Right.Right.GetHeight() > avlNode.Right.Left.GetHeight() {
			avlNode = avlNode.LeftRotate()
		} else {
			avlNode = avlNode.RightThenLeftRotate()
		}
	} else if avlNode.Left.GetHeight()-avlNode.Right.GetHeight() == 2 {
		if avlNode.Left.Left.GetHeight() > avlNode.Left.Right.GetHeight() {
			avlNode = avlNode.RightRotate()
		} else {
			avlNode = avlNode.LeftThenRightRotate()
		}
	}
	return avlNode
}

func (avlNode *AVLNode) Find(data interface{}) *AVLNode {
	var result *AVLNode
	switch compare(data, avlNode.Data) {
	case -1:
		result = avlNode.Left.Find(data)
	case 1:
		result = avlNode.Right.Find(data)
	case 0:
		return avlNode
	}
	return result
}

func (avlNode *AVLNode) FindMin() *AVLNode {
	var finded *AVLNode
	if avlNode.Left != nil {
		finded = avlNode.Left.FindMin()
	} else {
		finded = avlNode
	}
	return finded
}

func (avlNode *AVLNode) FindMax() *AVLNode {
	var finded *AVLNode
	if avlNode.Right != nil {
		finded = avlNode.Right.FindMax()
	} else {
		finded = avlNode
	}
	return finded
}

// 抓取数据
func (avlNode *AVLNode) GetData() interface{} {
	return avlNode.Data
}

// 设置
func (avlNode *AVLNode) SetData(data interface{}) {
	avlNode.Data = data
}

func (avlNode *AVLNode) GetLeft() *AVLNode {
	return avlNode.Left
}

func (avlNode *AVLNode) GetHeight() int {
	return avlNode.height
}

func (avlNode *AVLNode) GetRight() *AVLNode {
	return avlNode.Right
}
