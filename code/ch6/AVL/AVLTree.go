package main

import (
	"errors"
	"fmt"
)

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

func (avlNode *AVLNode) GetAll() []interface{} {
	values := make([]interface{}, 0)
	return AddValues(values, avlNode)
}

func AddValues(values []interface{}, avlNode *AVLNode) []interface{} {
	if avlNode != nil {
		values = AddValues(values, avlNode.Left)
		values = append(values, avlNode.Data)
		fmt.Println(avlNode.Data, avlNode.height)
		values = AddValues(values, avlNode.Right)
	}
	return values
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

func (avlNode *AVLNode) Insert(value interface{}) *AVLNode {
	if avlNode == nil {
		newNode := &AVLNode{value, nil, nil, 0}
		return newNode
	}
	switch compare(value, avlNode.Data) {
	case -1:
		avlNode.Left = avlNode.Left.Insert(value)
		avlNode = avlNode.adjust() // 自动平衡
	case 1:
		avlNode.Right = avlNode.Right.Insert(value)
		avlNode = avlNode.adjust()
	case 0:
		fmt.Println("数据已经存在")
	}
	avlNode.height = Max(avlNode.Left.GetHeight(), avlNode.Right.GetHeight()) + 1
	return avlNode
}

func (avlNode *AVLNode) Delete(value interface{}) *AVLNode {
	if avlNode == nil {
		return nil
	}
	switch compare(value, avlNode.Data) {
	case -1:
		avlNode.Left = avlNode.Left.Delete(value)
	case 1:
		avlNode.Right = avlNode.Right.Delete(value)
	case 0:
		// 删除在这里
		if avlNode.Left != nil && avlNode.Right != nil { // 左右都有节点
			avlNode.Data = avlNode.Right.FindMin()
			avlNode.Right = avlNode.Right.Delete(avlNode.Data)
		} else if avlNode.Left != nil { // 左孩子存在，右孩子存在或者不存在
			avlNode = avlNode.Left
		} else {
			// 只有一个右孩子，或者无孩子
			avlNode = avlNode.Right
		}
	}
	if avlNode != nil {
		avlNode.height = Max(avlNode.Left.GetHeight(), avlNode.Right.GetHeight()) + 1
		avlNode = avlNode.adjust() // 自动平衡
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
	if avlNode == nil {
		return nil
	}
	return avlNode.Data
}

// 设置
func (avlNode *AVLNode) SetData(data interface{}) {
	avlNode.Data = data
}

func (avlNode *AVLNode) GetLeft() *AVLNode {
	if avlNode == nil {
		return nil
	}
	return avlNode.Left
}

func (avlNode *AVLNode) GetHeight() int {
	if avlNode == nil {
		return 0
	}
	return avlNode.height
}

func (avlNode *AVLNode) GetRight() *AVLNode {
	if avlNode == nil {
		return nil
	}
	return avlNode.Right
}
