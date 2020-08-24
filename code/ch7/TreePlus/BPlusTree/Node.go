package BPlusTree

// 定义一个常量
const (
	// 叶子节点的存储数量
	MaxKV = 255
	// 中间节点最大存储数量
	MaxKC = 511
)

// 接口设计
type node interface {
	find(key int) (int, bool) // 查找key
	parent() *interiorNode    // 返回父亲节点
	setParent(*interiorNode)  // 设置父亲节点
	full() bool               // 判断是否满了
	countNum() int            // 统计元素的数量
}
