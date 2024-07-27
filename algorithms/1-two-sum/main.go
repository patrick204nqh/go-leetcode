package main

import (
	"github.com/fatih/color"
	"github.com/patrick204nqh/go-leetcode/utils/perf"
)

func twoSum(nums []int, target int) []int {
	numMap := make(map[int]int)
	for i, num := range nums {
		if j, ok := numMap[target-num]; ok {
			return []int{j, i}
		}
		numMap[num] = i
	}
	return nil
}

func main() {
	perf.Analyze("Description", func() {
		nums := []int{2, 7, 11, 15}
		target := 9
		result := twoSum(nums, target)
		color.Green("Result: %v", result) // Your expected output here
	})
}
