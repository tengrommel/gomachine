package ArrayList

type StackArrayX interface {
	Clear()                // 清空
	Size() int             // 大小
	Pop() interface{}      // 弹出
	Push(data interface{}) // 压入
	IsFull() bool          // 是否满了
	IsEmpty() bool         // 是否为空
}

type StackX struct {
	myArray *ArrayList
	MyIt    Iterator
}

func (s *StackX) Clear() {
	s.myArray.Clear()
	s.myArray.TheSize = 0
}

func (s *StackX) Size() int {
	return s.myArray.TheSize
}

func (s *StackX) Pop() interface{} {
	if !s.IsEmpty() {
		last := s.myArray.dataStore[s.myArray.TheSize-1]
		s.myArray.Delete(s.myArray.TheSize - 1)
		return last
	}
	return nil
}

func (s *StackX) Push(data interface{}) {
	if !s.IsFull() {
		s.myArray.Append(data)
	}
}

func (s *StackX) IsFull() bool { // 判断满了
	if s.myArray.TheSize >= 10 {
		return true
	} else {
		return false
	}
}

func (s *StackX) IsEmpty() bool { // 判断是否为空
	if s.myArray.TheSize == 0 {
		return true
	} else {
		return false
	}
}

func NewArrayListStackX() *StackX {
	myStack := new(StackX)
	myStack.myArray = NewArrayList()          // 数组
	myStack.MyIt = myStack.myArray.Iterator() // 迭代
	return myStack
}
