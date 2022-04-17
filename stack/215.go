package stack

func findKthLargest(nums []int, k int) int {
	stack0, stack1 := Stack{}, Stack{}

	for _, num := range nums {
		if stack0.Len() == 0 {
			stack0.Push(num)
		} else if num > stack0.Peek() {
			for stack0.Len() > 0 && num > stack0.Peek() {
				stack1.Push(stack0.Pop())
			}
			stack0.Push(num)
			for stack0.Len() < k && stack1.Len() > 0 {
				stack0.Push(stack1.Pop())
			}
			for stack1.Len() > 0 {
				stack1.Pop()
			}
		} else if stack0.Len() < k {
			stack0.Push(num)
		}
	}
	return stack0.Peek()
}
