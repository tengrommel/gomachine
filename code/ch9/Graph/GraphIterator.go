package Graph

// 迭代器接口
type iterator interface {
	Len() int
	Next() interface{}
	Value() interface{}
}

type LinkedMapIterator struct {
	m   *LinkedMap // 访问结构
	key interface{}
	iterator
}

func NewLinkedMapIterator(m *LinkedMap) *LinkedMapIterator {
	return new(LinkedMapIterator).init(m)
}

func (l *LinkedMapIterator) Len() int {
	return l.m.keyL.Len() // 长度
}

func (l *LinkedMapIterator) Next() interface{} {
	if l.key == nil {
		return nil
	}
	if l.key = l.m.nextKey(l.key); l.key == nil {
		return nil
	}
	return l.key
}

func (l *LinkedMapIterator) Value() interface{} {
	if l.key == nil {
		return nil
	}
	return l.key
}

func (l *LinkedMapIterator) init(m *LinkedMap) *LinkedMapIterator {
	l.m = m
	l.key = l.m.frontKey()
	return l
}
