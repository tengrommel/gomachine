package BPlusTree

import "sort"

// 存储数据
type kv struct {
	key   int // 数据
	value string
}

type kvs [MaxKC]kv

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
type leafNode struct {
	kvs    kvs           // 数据
	count  int           // 数量
	next   *leafNode     // 下一个节点
	parent *interiorNode // 父亲节点
}

// 创建叶子节点
func NewLeafNode(parent *interiorNode) *leafNode {
	return &leafNode{parent: parent}
}

// 查找数据
func (l *leafNode) find(key int) (int, bool) {
	myFunc := func(i int) bool {
		return l.kvs[i].key >= key
	} // myFunc是一个函数，主要进行数据的对比
	i := sort.Search(l.count, myFunc) // 查询
	if i < l.count && l.kvs[i].key == key {
		return i, true
	}
	return 0, false
}
