package collection

type Queue struct {
	vector []interface{}
}

func NewQueue() *Queue {
	return &Queue{
		vector: make([]interface{}, 0),
	}
}

func (q *Queue) Size() int {
	return len(q.vector)
}

func (q *Queue) Front() interface{} {
	if q.Size() == 0 {
		return nil
	}
	return q.vector[0]
}

func (q *Queue) Pop() interface{} {
	if q.Size() == 0 {
		return nil
	}
	ret := q.Front()
	q.vector = q.vector[1:]
	return ret
}

func (q *Queue) IsEmpty() bool {
	return len(q.vector) == 0
}
