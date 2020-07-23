package ArrayList

type StackArray interface {
	Clear()                // 清空
	Size() int             // 大小
	Pop() interface{}      // 弹出
	Push(data interface{}) // 压入
	IsFull() bool          // 是否满了
	IsEmpty() bool         // 是否为空
}

type Stack struct {
	myArray *ArrayList
	capsize int // 最大范围
}

func (s *Stack) Clear() {
	s.myArray.Clear()
	s.capsize = 10
}

func (s *Stack) Size() int {
	return s.myArray.TheSize
}

func (s *Stack) Pop() interface{} {
	if !s.IsEmpty() {
		last := s.myArray.dataStore[s.myArray.TheSize-1]
		s.myArray.Delete(s.myArray.TheSize - 1)
		return last
	}
	return nil
}

func (s *Stack) Push(data interface{}) {
	if !s.IsFull() {
		s.myArray.Append(data)
	}
}

func (s *Stack) IsFull() bool { // 判断满了
	if s.myArray.TheSize >= s.capsize {
		return true
	} else {
		return false
	}
}

func (s *Stack) IsEmpty() bool { // 判断是否为空
	if s.myArray.TheSize == 0 {
		return true
	} else {
		return false
	}
}

func NewArrayListStack() *Stack {
	myStack := new(Stack)
	myStack.myArray = NewArrayList() // 数组
	myStack.capsize = 10
	return myStack
}
