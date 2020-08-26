package Graph

import "container/list"

type LinkedMap struct {
	keyL  *list.List                  // 链表
	myMap map[interface{}]interface{} // 映射关系
}
