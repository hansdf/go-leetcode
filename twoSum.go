/*
Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.
You may assume that each input would have exactly one solution, and you may not use the same element twice.
You can return the answer in any order.
*/

package main

import "fmt"

func main() {
	twoSum([]int{2, 7, 15, 19}, 9)
}

func twoSum(nums []int, target int) []int { // takes an array of int and an int as parameters, returns an array of ints as results
	fmt.Println("hi")
	result := []int{}
	for i := 0; i < len(nums); i++ {
		if nums[i]+nums[i+1] == target {
			fmt.Print("found match. The two numbers on the list ")
			fmt.Print(nums)
			fmt.Print(" that add up to ")
			fmt.Print(target)
			fmt.Print(" are ")
			fmt.Print([]int{nums[i], nums[i+1]})
			result = []int{nums[i], nums[i+1]}
			return result
		} else {
			fmt.Println("no sum can be found to match target")
			return result
		}
	}
	return result
}
