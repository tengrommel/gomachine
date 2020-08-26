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

// 1	6
// 123456789
func (bpt *BPlusTree) Insert(key int, value string) {
	_, oldIndex, leaf := search((*bpt)[-1], key)
	p := leaf.Parent() // 保存父亲节点
	// 插入叶子节点，判断是否分裂
	mid, nextLeaf, bump := leaf.insert(key, value)
	if !bump { // 没有分谢
		return
	}
	// f分裂的节点插入B+树
	(*bpt)[mid] = nextLeaf
	var midNode node
	midNode = leaf.next                  // 设置父亲节点
	leaf.next.SetParent(p)               // 分裂的节点设置父亲节点
	interior, interiorP := p, p.Parent() // 获取中间点解，父亲节点
	//平衡过程，迭代向上判断，是否需要平衡
	for {
		var newInterior *interiorNode // 新的节点
		// 判断是否到达根节点
		isRoot := interiorP == nil
		if !isRoot {
			oldIndex, _ = interiorP.find(key) // 查找
		}
		// 叶子节点分裂后的中间节点传入父亲的中间节点，传入分裂
		mid, newInterior, bump = interior.insert(mid, midNode)
		if !bump {
			return
		}
		(*bpt)[newInterior.kcs[0].key] = newInterior // 插入并填充好了map
		if !isRoot {
			interiorP.kcs[oldIndex].child = newInterior // 没有到根节点，直接插入父亲节点
			newInterior.SetParent(interiorP)
			midNode = interior
		} else {
			// 更新节点
			(*bpt)[interior.kcs[0].key] = (*bpt)[-1]       // 备份下根节点
			(*bpt)[-1] = NewInteriorNode(nil, newInterior) // 根节点插入新的根节点
			node := (*bpt)[-1].(*interiorNode)
			node.insert(mid, interior) // 重新插入
			(*bpt)[-1] = node
			newInterior.SetParent(node)
		}
		interior, interiorP = interiorP, interiorP.Parent()

	}

}
