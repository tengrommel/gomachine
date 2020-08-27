package Graph

import "container/list"

type LinkedMap struct {
	keyL  *list.List                  // 链表
	myMap map[interface{}]interface{} // 映射关系
}

// 数据结构的
func (lm *LinkedMap) init() *LinkedMap {
	lm.keyL = list.New()
	lm.myMap = make(map[interface{}]interface{})
	return lm
}

func (lm *LinkedMap) exist(key interface{}) bool {
	_, ok := lm.myMap[key]
	return ok
}

// 插入数据， key value
func (lm *LinkedMap) add(key interface{}, value interface{}) {
	if !lm.exist(key) {
		e := lm.keyL.PushBack(key) // 压入数据
		lm.myMap[key] = []interface{}{e, value}
	} else {
		lm.myMap[key].([]interface{})[1] = value
	}
}

// 抓取数据
func (lm *LinkedMap) get(key interface{}) interface{} {
	if lm.exist(key) {
		return lm.myMap[key].([]interface{})[1]
	} else {
		return nil
	}
}

func (lm *LinkedMap) delete(key interface{}) {
	if lm.exist(key) {
		i := lm.myMap[key].([]interface{})
		lm.keyL.Remove(i[0].(*list.Element))
		delete(lm.myMap, key) // 删除
	}
}

func (lm *LinkedMap) frontKey() interface{} {
	if lm.keyL.Len() == 0 {
		return nil
	}
	return lm.keyL.Front().Value
}

func (lm *LinkedMap) nextKey(key interface{}) interface{} {
	if e := lm.myMap[key].([]interface{})[0].(*list.Element).Next(); e != nil {
		return e.Value
	}
	return nil
}

func (lm *LinkedMap) prevKey(key interface{}) interface{} {
	if e := lm.myMap[key].([]interface{})[0].(*list.Element).Prev(); e != nil {
		return e.Value
	}
	return nil
}
