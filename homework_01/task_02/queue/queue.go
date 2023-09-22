package queue

type Queue struct {
	size int
	data []interface{}
}

func (q *Queue) IsEmpty() bool {
	return q.size == 0
}

func (q *Queue) Push(value interface{}) {
	q.data = append(q.data, value)
	q.size++
}

func (q *Queue) Pop() bool {
	if q.IsEmpty() {
		return false
	}

	q.data = q.data[1:]
	q.size--

	return true
}

func (q *Queue) Front() interface{} {
	var result interface{}

	if !q.IsEmpty() {
		result = q.data[0]
	}

	return result
}
