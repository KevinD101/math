package main

import (
	"github.com/KevinD101/math/sort/quick"
	"log"
)

func main() {
	nums := []int{1,9,4,7,8, 5,8,5}
	quick.QuickSort(nums, 0, len(nums) - 1)
	log.Println(nums)
}
