package BPlusTree

type BPlusTree map[int]node // 定义

func NewBPlusTree() *BPlusTree {
	bt := BPlusTree{}
	leaf := NewLeafNode(nil)        // 叶子节点
	r := NewInteriorNode(nil, leaf) // 中间节点
	leaf.parent = r                 // 设定父亲节点
	bt[-1] = r
	bt[0] = leaf
	return &bt
}

// 返回根节点
func (bpt *BPlusTree) Root() node {
	return (*bpt)[-1]
}

func (bpt *BPlusTree) First() node {
	return (*bpt)[0]
}

// 统计下数量
func (bpt *BPlusTree) Count() int {
	count := 0
	leaf := (*bpt)[0].(*LeafNode)
	for {
		count += leaf.CountNum() // 数量叠加
		if leaf.next == nil {
			break
		}
		leaf = leaf.next
	}
	return count
}

func (bpt *BPlusTree) Values() []*LeafNode {
	nodes := make([]*LeafNode, 0) // 开辟节点
	leaf := (*bpt)[0].(*LeafNode)
	for {
		nodes = append(nodes, leaf)
		if leaf.next == nil {
			break
		}
		leaf = leaf.next
	}
	return nodes
}

func (bpt *BPlusTree) Insert(key int, value string) {

}

// 查找数据
func (bpt *BPlusTree) Search(key int) (string, bool) {
	kv, _, _ := search((*bpt)[-1], key) // 查找
	if kv == nil {
		return "", false
	} else {
		return kv.value, true
	}
}

// 搜索数据
func search(n node, key int) (*kv, int, *LeafNode) {
	curr := n
	oldIndex := -1
	for {
		switch t := curr.(type) {
		case *LeafNode:
			i, ok := t.find(key)
			if !ok {
				return nil, oldIndex, t
			} else {
				return &t.kvs[i], oldIndex, t
			}
		case *interiorNode:
			i, _ := t.find(key) // 中间节点查找
			curr = t.kcs[i].child
			oldIndex = i
		default:
			panic("异常节点")
		}
	}
}
