package Queue

type MyQueue interface {
	Size() int                // 大小
	Front() interface{}       // 第一个元素
	End() interface{}         // 最后一个
	IsEmpty() bool            // 是否为空
	EnQueue(data interface{}) // 入队
	DeQueue() interface{}     // 出队
	Clear()                   // 清空
}

type Queue struct {
	dataStore []interface{} // 队列的数据存储
	theSize   int           // 队列的大小
}

func (q *Queue) Size() int {
	return q.theSize
}

func (q *Queue) Front() interface{} {
	if q.Size() == 0 {
		return nil
	}
	return q.dataStore[0]
}

func (q *Queue) End() interface{} {
	if q.Size() == 0 {
		return nil
	}
	return q.dataStore[q.Size()-1]
}

func (q *Queue) IsEmpty() bool {
	return q.theSize == 0
}

func (q *Queue) EnQueue(data interface{}) {
	q.dataStore = append(q.dataStore, data) // 入队
	q.theSize++
}

func (q *Queue) DeQueue() interface{} {
	if q.Size() == 0 {
		return nil
	}
	data := q.dataStore[0]
	if q.Size() > 1 {
		q.dataStore = q.dataStore[1:q.Size()]
	} else {
		q.dataStore = make([]interface{}, 0)
	}
	q.theSize--
	return data // 返回数据
}

func (q *Queue) Clear() {
	q.dataStore = make([]interface{}, 0)
	q.theSize = 0
}

func NewQueue() *Queue {
	myQueue := new(Queue) // 初始化，开辟结构体
	myQueue.dataStore = make([]interface{}, 0)
	myQueue.theSize = 0
	return myQueue
}
