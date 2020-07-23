package ArrayList

import "errors"

type Iterator interface {
	HasNext() bool              // 是否有下一个
	Next() (interface{}, error) // 下一个
	Remove()                    // 删除
	GetIndex() int              // 得到索引
}

type Iterable interface {
	Iterator() Iterator
}

// 构造指针访问数组
type ArrayListIterator struct {
	list         *ArrayList // 数组指针
	currentIndex int        // 当前的索引
}

func (it *ArrayListIterator) HasNext() bool {
	return it.currentIndex < it.list.TheSize // 是否有下一个
}

func (it *ArrayListIterator) Next() (interface{}, error) {
	if !it.HasNext() {
		return nil, errors.New("没有下一个")
	}
	value, err := it.list.Get(it.currentIndex)
	it.currentIndex++
	return value, err
}

func (it *ArrayListIterator) Remove() {
	it.currentIndex--
	it.list.Delete(it.currentIndex) // 删除一个元素
}

func (it *ArrayListIterator) GetIndex() int {
	return it.currentIndex
}
