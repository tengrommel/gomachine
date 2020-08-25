package BPlusTree

import "sort"

// 存储数据
type kv struct {
	key   int // 数据
	value string
}

type kvs [MaxKV]kv

// 求数组的长度
func (kvs *kvs) Len() int {
	return len(kvs)
}

// 交换数据
func (kvs *kvs) Swap(i, j int) {
	kvs[i], kvs[j] = kvs[j], kvs[i]
}

// 判断大小
func (kvs *kvs) Less(i, j int) bool {
	return kvs[i].key < kvs[j].key
}

// 叶子节点
type LeafNode struct {
	kvs    kvs           // 数据
	count  int           // 数量
	next   *LeafNode     // 下一个节点
	parent *interiorNode // 父亲节点
}

// 创建叶子节点
func NewLeafNode(parent *interiorNode) *LeafNode {
	return &LeafNode{parent: parent}
}

// 查找数据
func (l *LeafNode) find(key int) (int, bool) {
	myFunc := func(i int) bool {
		return l.kvs[i].key >= key
	} // myFunc是一个函数，主要进行数据的对比
	i := sort.Search(l.count, myFunc) // 查询
	if i < l.count && l.kvs[i].key == key {
		return i, true
	}
	return i, false
}

// 数据插入叶子节点
func (l *LeafNode) insert(key int, value string) (int, *LeafNode, bool) {
	i, ok := l.find(key)
	if ok {
		// key已经存在，更新value
		l.kvs[i].value = value
		return 0, nil, false
	}
	if !l.full() {
		// 数组删除整体往后移动
		copy(l.kvs[i+1:], l.kvs[i:l.count])
		l.kvs[i].key = key
		l.kvs[i].value = value
		l.count++
		return 0, nil, false
	} else {
		next := l.split()
		if key < next.kvs[0].key {
			l.insert(key, value)
		} else {
			next.insert(key, value)
		}
		return next.kvs[0].key, next, true
	}

}

func (l *LeafNode) full() bool {
	return l.count == MaxKV // 判断是否满了
}

func (l *LeafNode) Parent() *interiorNode {
	return l.parent
}

func (l *LeafNode) SetParent(p *interiorNode) {
	l.parent = p
}

func (l *LeafNode) CountNum() int {
	return l.count
}

// 叶子节点，分裂 // 123	4567
func (l *LeafNode) split() *LeafNode {
	next := NewLeafNode(nil)                // 新建一个右边节点
	copy(next.kvs[0:], l.kvs[l.count/2+1:]) // 复制数据到右边节点
	l.InitArray(l.count/2 + 1)              // 后半部数据清空
	next.count = MaxKV - l.count/2 - 1      // 下一个节点的数量
	next.next = l.next                      // 调整指针的指向
	l.count = l.count/2 + 1
	l.next = next
	return next
}

// 初始化数组，数组每个元素初始化
func (l *LeafNode) InitArray(num int) {
	for i := num; i < len(l.kvs); i++ {
		l.kvs[i] = kv{}
	}
}
