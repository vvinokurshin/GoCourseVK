package stack

type Stack struct {
	size int
	data []interface{}
}

func (st *Stack) IsEmpty() bool {
	return st.size == 0
}

func (st *Stack) Push(value interface{}) {
	st.data = append(st.data, value)
	st.size++
}

func (st *Stack) Pop() bool {
	if st.IsEmpty() {
		return false
	}

	st.data = st.data[:st.size-1]
	st.size--

	return true

}

func (st *Stack) Top() interface{} {
	var result interface{}

	if !st.IsEmpty() {
		result = st.data[st.size-1]
	}

	return result
}
