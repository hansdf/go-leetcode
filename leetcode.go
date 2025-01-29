package main

import "fmt"

func main() {
	// twoSum([]int{2, 15, 7, 19, 22, 8}, 15)
	// strStr("sadbutsad", "sad") //haystack = "sadbutsad", needle = "sad"
	singleNumber([]int{2, 22, 15, 15, 8, 7, 19, 19, 22, 2, 8})
}

/*
Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.
You may assume that each input would have exactly one solution, and you may not use the same element twice.
You can return the answer in any order.
*/
func twoSum(nums []int, target int) []int { // takes an array of int and an int as parameters, returns an array of ints as results
	result := []int{}
	for i := 0; i < len(nums); i++ {
		for j := 1; j < len(nums); j++ {
			if nums[i] != nums[j] && nums[i]+nums[j] == target {
				fmt.Printf("Found match. The two numbers on the list that match up to target %v are %v, located in indexes %v", target, []int{nums[i], nums[j]}, []int{i, j})
				result = []int{i, j}
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

/*
Given two strings needle and haystack, return the index of the first occurrence of needle in haystack, or -1 if needle is not part of haystack.
haystack and needle consist of only lowercase English characters.
*/
func strStr(haystack string, needle string) int {

	haystackLen := len(haystack)
	needleLen := len(needle)

	for i := 0; i <= haystackLen; i++ {
		if haystack[i:i+needleLen] == needle {
			fmt.Printf("Index for needle found at %v", i)
			return i // index
		}
	}
	return -1
}

// Given a non-empty array of integers nums, every element appears twice except for one. Find that single one.
// Each element in the array appears twice except for one element which appears only once.
func singleNumber(nums []int) int {
	for i := 0; i < len(nums); i++ {
		counter := 0

		for j := 0; j < len(nums); j++ {
			if nums[i] == nums[j] {
				counter++
			}
		}

		if counter == 1 {
			fmt.Printf("Found unique item in list: %v", nums[i])
			return nums[i]
		}
	}

	return -1 // Return -1 if no unique number is found
}
