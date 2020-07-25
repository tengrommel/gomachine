package main

type Node struct {
	data  interface{}
	pNext *Node
}

func (n *Node) IsEmpty() bool {
	if n.pNext == nil {
		return true
	} else {
		return false
	}
}

func (n *Node) Push(data interface{}) {
	newNode := &Node{data, nil}
	newNode.pNext = n.pNext
	n.pNext = newNode // 头部插入
}

// 栈的操作都是一样的
func (n *Node) Pop() interface{} {
	if n.IsEmpty() == true {
		return nil
	}
	value := n.pNext.data // 要弹出的数据
	n.pNext = n.pNext.pNext
	return value
}

func (n *Node) length() int {
	pNext := n
	length := 0
	for pNext.pNext != nil {
		pNext = pNext.pNext // 节点循环跳跃
		length++
	}
	return length
}

type LinkStack interface {
	IsEmpty() bool
	Push(data interface{})
	Pop() interface{}
	length() int
}

func NewStack() *Node {
	return &Node{} // 返回一个节点的指针
}
