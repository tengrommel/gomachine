package main

// 定义常量红黑
const (
	RED   = true
	BLACK = false
)

// 红黑树的结构
type RBNode struct {
	Left   *RBNode // 左节点
	Right  *RBNode // 右节点
	Parent *RBNode // 父亲节点
	Color  bool    // 颜色
	//DataItem interface{} // 数据
	Item // 数据接口
}

type Item interface {
	Less(than Item) bool
}

// 红黑树
type RBTree struct {
	NIL   *RBNode
	Root  *RBNode
	count uint
}

// 比大小
func less(x, y Item) bool {
	return x.Less(y)
}

// 初始化内存
func NewRBTree() *RBTree {
	return new(RBTree).Init()
}

// 初始化红黑树
func (rbt *RBTree) Init() *RBTree {
	node := &RBNode{nil, nil, nil, BLACK, nil}
	return &RBTree{node, node, 0}
}

// 获取红黑树的长度
func (rbt *RBTree) Len() uint {
	return rbt.count
}

// 取得红黑树的极大值
func (rbt *RBTree) max(x *RBNode) *RBNode {
	if x == rbt.NIL { // 只有一个元素
		return rbt.NIL
	}
	for x.Right != rbt.NIL {
		x = x.Right
	}
	return x
}

// 取得红黑树的极小值
func (rbt *RBTree) min(x *RBNode) *RBNode {
	if x == rbt.NIL { // 只有一个元素
		return rbt.NIL
	}
	for x.Left != rbt.NIL {
		x = x.Left
	}
	return x
}

// 搜索红黑树
func (rbt *RBTree) search(x *RBNode) *RBNode {
	pNode := rbt.Root
	for pNode != rbt.NIL {
		if less(pNode.Item, x.Item) {
			pNode = pNode.Right
		} else if less(x.Item, pNode.Item) {
			pNode = pNode.Left
		} else {
			break // 找到
		}
	}
	return pNode
}

func (rbt *RBTree) leftRotate(x *RBNode) {
	if x.Right == rbt.NIL {
		return // 左旋转,逆时针,右孩子不可以为0
	}
	y := x.Right
	x.Right = y.Left // 实现旋转的左旋
	if y.Left != rbt.NIL {
		y.Left.Parent = x // 设置父亲节点
	}
	y.Parent = x.Parent // 交换父节点
	if x.Parent == rbt.NIL {
		// 根节点
		rbt.Root = y
	} else if x == x.Parent.Left {
		x.Parent.Left = y
	} else {
		x.Parent.Right = y
	}
	y.Left = x
	x.Parent = y
}

func (rbt *RBTree) rightRotate(x *RBNode) {
	if x.Left == nil {
		return // 右边旋转,左子树不可以为空
	}
	y := x.Left
	x.Left = y.Right
	if y.Right != rbt.NIL {
		y.Right.Parent = x
	}
	y.Parent = x.Parent
	if x.Parent == rbt.NIL {
		rbt.Root = y
	} else if x == x.Parent.Left {
		x.Parent.Left = y
	} else {
		x.Parent.Right = y
	}
	y.Right = x
	x.Parent = y
}

// 插入
func (rbt *RBTree) insert(z *RBNode) *RBNode {
	// 寻找插入位置
	x := rbt.Root
	y := rbt.NIL
	for x != rbt.NIL {
		y = x                     // 备份位置,数据插入x,y之间
		if less(z.Item, x.Item) { // 小于
			x = x.Left
		} else if less(x.Item, z.Item) { // 大于
			x = x.Right
		} else {
			return x // 数据已经存在,无法插入
		}
	}
	z.Parent = y
	if y == rbt.NIL {
		rbt.Root = z
	} else if less(z.Item, y.Item) {
		y.Left = z // 小于左边插入
	} else {
		y.Right = z // 大于右边插入
	}
	rbt.count++
	rbt.insertFixup(z)
	return z
}

// 插入后调整平衡
func (rbt *RBTree) insertFixup(z *RBNode) {

}
