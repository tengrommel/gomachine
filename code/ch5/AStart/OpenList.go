package main

type OpenList []*AstarPoint

func (self OpenList) Len() int {
	return len(self)
}

func (self OpenList) Less(i, j int) bool {
	// 比较大小
	return self[i].fVal < self[j].fVal
}

func (self OpenList) Swap(i, j int) {
	// 数据的交换
	self[i], self[j] = self[j], self[i]
}

func (this *OpenList) Push(data interface{}) {
	*this = append(*this, data.(*AstarPoint)) // 节点加入到栈中
}

func (this *OpenList) Pop() interface{} {
	old := *this
	n := len(old)
	x := old[n-1]
	*this = old[0 : n-1] // 切片的截取，去掉最后一个
	return x
}
