package divide

import "math/rand"

func partition(nums []int, left, right, mid int) int {
	if left >= right {
		return left
	}

	swap(nums, mid, right)

	insertIdx := left
	for j := left; j < right; j++ {
		if nums[j] <= nums[right] {
			swap(nums, insertIdx, j)
			insertIdx += 1
		}
	}

	swap(nums, insertIdx, right)

	return insertIdx
}

func swap(nums []int, i, j int) {
	if i >= len(nums) || j >= len(nums) {
		return
	}
	tmp := nums[j]
	nums[j] = nums[i]
	nums[i] = tmp
}

func quickSort(nums []int) {
	left, right := 0, len(nums)-1
	quickSortRng(nums, left, right)
}

func quickSortRng(nums []int, from, to int) {
	if from >= to {
		return
	}
	mid := random(from, to)
	mid = partition(nums, from, to, mid)

	quickSortRng(nums, from, mid-1)
	quickSortRng(nums, mid+1, to)

}

func random(from, to int) int {
	return rand.Intn(to-from+1) + from
}

func findKthLargest(nums []int, k int) int {
	return quickSelect(nums, k, 0, len(nums)-1)
}

func quickSelect(nums []int, k int, from, to int) int {
	mid := random(from, to)
	mid = partition(nums, from, to, mid)

	if mid == k {
		return nums[k]
	} else if mid > k {
		return quickSelect(nums, k, from, mid-1)
	}
	return quickSelect(nums, k, mid+1, to)
}
