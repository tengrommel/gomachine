package SingleLink

// 链表的节点
type Node struct {
	value interface{}
	pNext *Node
}

// 构造一个节点
func NewNode(data interface{}) *Node {
	return &Node{
		data,
		nil,
	}
}

// 返回节点
func (node *Node) Value() interface{} {
	return node.value
}

// 返回一个指针
func (node *Node) PNext() *Node {
	return node.pNext
}
