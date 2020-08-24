package BPlusTree

// 存储数据
type kc struct {
	key   int  // 数据
	child node // 接口类型
}

type kcs [MaxKC + 1]kc // 数组，并非切片

// 中间节点数据结构
type interiorNode struct {
	kcs    kcs           // 数据存储数据
	count  int           // 存储元素的数量
	parent *interiorNode // 指向父亲节点
}
