package SingleLink

import "fmt"

type SingleLink interface {
	// 增删改查
	GetFirstNode() *Node            // 抓取头部节点
	InsertNodeFront(node *Node)     // 头部插入
	InsertNodeBack(node *Node)      // 尾部插入
	GetNodeAtIndex(index int) *Node // 根据索引抓取指定位置的节点
	DeleteNode(dest *Node) bool     // 删除一个节点
	DeleteAtIndex(index int)        // 删除指定位置的节点
	String() string                 // 返回链表字符串
}

// 链表的结构
type List struct {
	head   *Node // 链表的头部指针
	length int   // 链表的长度
}

func (l *List) GetFirstNode() *Node {
	return l.head.pNext
}

func (l *List) InsertNodeFront(node *Node) {
	if l.head == nil {
		l.head.pNext = node
		node.pNext = nil
		l.length++ // 插入节点，数据追加
	} else {
		bak := l.head
		node.pNext = bak.pNext
		bak.pNext = node
		l.length++ // 插入节点，数据追加
	}
}

func (l *List) InsertNodeBack(node *Node) {
	if l.head == nil {
		l.head.pNext = node
		node.pNext = nil
		l.length++
	} else {
		bak := l.head
		for bak.pNext != nil { // 循环到最后
			bak = bak.pNext
		}
		bak.pNext = node
		l.length++
	}
}

func (l *List) GetNodeAtIndex(index int) *Node {
	panic("implement me")
}

func (l *List) DeleteNode(dest *Node) bool {
	panic("implement me")
}

func (l *List) DeleteAtIndex(index int) {
	panic("implement me")
}

func (l *List) String() string {
	var listString string
	p := l.head
	for p.pNext != nil {
		listString += fmt.Sprintf("%v-->", p.pNext.value)
		p = p.pNext
	}
	listString += fmt.Sprintf("nil")
	return listString // 打印链表字符串
}

func NewList() *List {
	head := NewNode(nil)
	return &List{head: head, length: 0}
}
