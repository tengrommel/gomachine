package ArrayList

import (
	"errors"
	"fmt"
)

// 接口
type List interface {
	Size() int                                  // 数值大小
	Get(index int) (interface{}, error)         // 抓取第几个元素
	Set(index int, newVal interface{}) error    // 修改数据
	Insert(index int, newVal interface{}) error // 插入数据
	Append(newVal interface{})                  // 追加 没有error
	Clear()                                     // 清空
	Delete(index int) error                     // 删除
	String() string                             // 返回字符串
}

// 数据结构
type ArrayList struct {
	dataStore []interface{} // 数组的存储
	theSize   int           // 数组的大小
}

func (list *ArrayList) Size() int {
	return list.theSize // 返回数组的大小
}

// 抓取数据
func (list *ArrayList) Get(index int) (interface{}, error) {
	if index < 0 || index >= list.theSize {
		return nil, errors.New("索引越界")
	}
	return list.dataStore[index], nil
}

func (list *ArrayList) Set(index int, newVal interface{}) error {
	panic("implement me")
}

func (list *ArrayList) Insert(index int, newVal interface{}) error {
	panic("implement me")
}

func (list *ArrayList) Append(newVal interface{}) {
	list.dataStore = append(list.dataStore, newVal) // 叠加数据
	list.theSize++
}

func (list *ArrayList) Clear() {
	panic("implement me")
}

func (list *ArrayList) Delete(index int) error {
	panic("implement me")
}

func (list *ArrayList) String() string {
	return fmt.Sprint(list.dataStore)
}

// 开辟内存
func NewArrayList() *ArrayList {
	list := new(ArrayList)                      // 初始化结构体
	list.dataStore = make([]interface{}, 0, 10) // 开辟空间10个
	list.theSize = 0
	return list
}
