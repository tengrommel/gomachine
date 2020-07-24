package StackArray

type StackArray interface {
	Clear()                // 清空
	Size() int             // 大小
	Pop() interface{}      // 弹出
	Push(data interface{}) // 压入
	IsFull() bool          // 是否满了
	IsEmpty() bool         // 是否为空
}

type Stack struct {
	dataSource  []interface{}
	capsize     int // 最大范围
	currentSize int // 实际使用大小
}

func (s *Stack) Clear() {
	s.dataSource = make([]interface{}, 0, 10)
	s.currentSize = 0
	s.capsize = 10
}

func (s *Stack) Size() int {
	return s.currentSize
}

func (s *Stack) Pop() interface{} {
	if !s.IsEmpty() {
		last := s.dataSource[s.currentSize-1]         // 最后一个数据
		s.dataSource = s.dataSource[:s.currentSize-1] // 删除最后一个
		s.currentSize--
		return last
	}
	return nil
}

func (s *Stack) Push(data interface{}) {
	if !s.IsFull() {
		s.dataSource = append(s.dataSource, data)
		s.currentSize++
	}
}

func (s *Stack) IsFull() bool { // 判断满了
	if s.currentSize == s.capsize {
		return true
	} else {
		return false
	}
}

func (s *Stack) IsEmpty() bool { // 判断是否为空
	if s.currentSize == 0 {
		return true
	} else {
		return false
	}
}

func NewStack() *Stack {
	myStack := new(Stack)
	myStack.dataSource = make([]interface{}, 0, 10)
	myStack.capsize = 1000
	myStack.currentSize = 0
	return myStack
}
