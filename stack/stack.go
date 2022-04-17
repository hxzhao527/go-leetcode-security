package stack

type Stack []int

func (stack *Stack) Push(ele int) {
	*stack = append(*stack, ele)
}
func (stack *Stack) Peek() int {
	slice := []int(*stack)
	length := len(slice)
	return slice[length-1]
}
func (stack *Stack) Pop() int {
	slice := []int(*stack)
	length := len(slice)
	last := slice[length-1]
	*stack = slice[:length-1]
	return last
}
func (stack *Stack) Len() int {
	return len(*stack)
}
