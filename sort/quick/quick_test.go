package quick

import (
	"log"
	"testing"
)

func TestQuickSort(t *testing.T){
	nums := []int{1,9,4,7,8, 5,8,5}
	QuickSort(nums, 0, len(nums) - 1)
	log.Println(nums)
}