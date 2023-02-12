package quick


func QuickSort(nums []int, left, right int) {
	mid := PartSort(nums, left, right)
	if left >= right {
		return
	}
	QuickSort(nums, left, mid-1)
	QuickSort(nums, mid+1, right)
}

func PartSort(nums []int, left, right int) int {
	mid := left
	for left < right {
		for left < right && nums[mid] <= nums[right] {
			right--
		}
		for left < right && nums[mid] > nums[left] {
			left++
		}
		nums[left], nums[right] = nums[right], nums[left]
	}
	nums[left], nums[mid] = nums[mid], nums[left]
	return left
}