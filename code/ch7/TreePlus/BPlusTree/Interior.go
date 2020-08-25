package BPlusTree

import "sort"

// 存储数据
type kc struct {
	key   int  // 数据
	child node // 接口类型
}

type kcs [MaxKC + 1]kc // 数组，并非切片

// 求数组的长度
func (kcs *kcs) Len() int {
	return len(kcs)
}

// 交换数据
func (kcs *kcs) Swap(i, j int) {
	kcs[i], kcs[j] = kcs[j], kcs[i]
}

// 判断大小
func (kcs *kcs) Less(i, j int) bool {
	if kcs[i].key == 0 { // 中间节点的数组第一个会空着，search
		return false
	}
	if kcs[j].key == 0 {
		return true
	}
	return kcs[i].key < kcs[j].key
}

// 中间节点数据结构
type interiorNode struct {
	kcs    kcs           // 数据存储数据
	count  int           // 存储元素的数量
	parent *interiorNode // 指向父亲节点
}

// 新建一本中间节点
func NewInteriorNode(p *interiorNode, largestChild node) *interiorNode {
	in := &interiorNode{parent: p, count: 1}
	if largestChild != nil {
		in.kcs[0].child = largestChild // 设置插入节点
	}
	return in
}

func (interiorNode *interiorNode) find(key int) (int, bool) {
	myFunc := func(i int) bool {
		return interiorNode.kcs[i].key >= key
	} // myFunc是一个函数，主要进行数据的对比
	i := sort.Search(interiorNode.count, myFunc) // 查询
	if i < interiorNode.count && interiorNode.kcs[i].key == key {
		return i, true
	}
	return i, false
}

func (interiorNode *interiorNode) full() bool {
	return interiorNode.count == MaxKV // 判断是否满了
}

func (interiorNode *interiorNode) Parent() *interiorNode {
	return interiorNode.parent // 返回父亲节点
}

func (interiorNode *interiorNode) SetParent(p *interiorNode) {
	interiorNode.parent = p // 设置父亲节点
}

func (interiorNode *interiorNode) CountNum() int {
	return interiorNode.count
}

// 初始化数组，数组每个元素初始化
func (interiorNode *interiorNode) InitArray(num int) {
	for i := num; i < len(interiorNode.kcs); i++ {
		interiorNode.kcs[i] = kc{}
	}
}

func (in *interiorNode) insert(key int, child node) (int, *interiorNode, bool) {
	// 确定位置
	i, _ := in.find(key)
	if !in.full() {
		// 数据插入之前，整体向后移动
		copy(in.kcs[i+1:], in.kcs[i:in.count])
		// 设置子节点分裂以后的元素，设置key
		in.kcs[i].key = key
		in.kcs[i].child = child
		child.setParent(in)
		in.count++
		return 0, nil, false
	} else {
		in.kcs[MaxKC].key = key
		in.kcs[MaxKC].child = child // 存储到最后
		child.setParent(in)         // 设置父亲节点
		next, midKey := in.split()  // 切割
		return midKey, next, true   // 返回中间节点
	}
}

// 123 45
func (in *interiorNode) split() (*interiorNode, int) {
	// 节点分裂，节点插入正确位置
	sort.Sort(&in.kcs) // 确保有序
	// 取得中间元素
	midIndex := MaxKC / 2
	midChild := in.kcs[midIndex].child // 取得中间节点
	midKey := in.kcs[midIndex].key     // 取得键值

	// 新建一个中间节点
	next := NewInteriorNode(nil, nil)
	copy(next.kcs[0:], in.kcs[midIndex+1:]) // 拷贝数据
	in.InitArray(midIndex + 1)              // 数据初始化
	next.count = MaxKC - midIndex           // 下一个节点的数量
	for i := 0; i < next.count; i++ {
		next.kcs[i].child.setParent(next)
	}
	in.count = midIndex + 1
	in.kcs[in.count-1].key = 0
	in.kcs[in.count-1].child = midChild
	midChild.setParent(in)
	return next, midKey
}
