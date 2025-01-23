package main

import "fmt"

func main() {
	twoSum([]int{2, 15, 7, 19, 22, 8}, 15)
	fizzbuzz(35)
}

/*
Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.
You may assume that each input would have exactly one solution, and you may not use the same element twice.
You can return the answer in any order.
*/
func twoSum(nums []int, target int) []int { // takes an array of int and an int as parameters, returns an array of ints as results
	fmt.Println("hi")
	result := []int{}
	for i := 0; i < len(nums); i++ {
		for j := 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				fmt.Print("found match. The two numbers on the list that match are ")
				fmt.Print([]int{nums[i], nums[j]})
				result = []int{nums[i], nums[j]}
				return result
			}
		}
	}
	return result
}

/*
The FizzBuzz problem  requires writing a function that prints numbers from 1 to a given limit, but with a twist:
	For multiples of 3, print "Fizz" instead of the number.
	For multiples of 5, print "Buzz" instead of the number.
	For numbers which are multiples of both 3 and 5, print "FizzBuzz".
*/
func fizzbuzz(max int) {
	for i := 0; i <= max; i++ {
		if i%3 == 0 && i%5 == 0 { // Check for divisibility by both 3 and 5 first
			fmt.Println("fizzbuzz")
		} else if i%3 == 0 { // Check for divisibility by 3
			fmt.Println("fizz")
		} else if i%5 == 0 { // Check for divisibility by 5
			fmt.Println("buzz")
		} else { // Default case for non-multiples of 3 or 5
			fmt.Println(i)
		}
	}
}
