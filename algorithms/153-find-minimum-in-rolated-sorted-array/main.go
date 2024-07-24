package main

import (
    "github.com/fatih/color"
)

func findMinimumInRolatedSortedArray(nums []int) int {
    left, right := 0, len(nums) - 1
    for left < right {
        mid := left + (right - left) / 2
        if nums[mid] > nums[right] {
            left = mid + 1
        } else {
            right = mid
        }
    }
    return nums[left]
}

func main() {
    result := findMinimumInRolatedSortedArray([]int{3,4,5,1,2})
    color.Green("Result: %v", result) // Your expected output here
}
