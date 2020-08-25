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
	Parent() *interiorNode    // 返回父亲节点
	SetParent(*interiorNode)  // 设置父亲节点
	full() bool               // 判断是否满了
	CountNum() int            // 统计元素的数量
}
