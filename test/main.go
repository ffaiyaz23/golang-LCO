package main

import "fmt"

func containsDuplicate(nums []int) bool {
	mp := make(map[int]int)

	for _, num := range nums {
		if value, exists := mp[num]; exists {
			fmt.Println(value)
			fmt.Println(exists)
			return true
		}

		mp[num] = 1
		// fmt.Println(mp[num])
	}

	return false
}

func main() {
	fmt.Println("Maps in Golang")

	// Example 1:
	nums1 := []int{1, 2, 3, 1}
	fmt.Println("\nExample 1:")
	fmt.Printf("Input: nums = %v\n", nums1)
	fmt.Printf("Output: %v\n", containsDuplicate(nums1))
	// Explanation: The element 1 occurs at the indices 0 and 3.

	// Example 2:
	nums2 := []int{1, 2, 3, 4}
	fmt.Println("\nExample 2:")
	fmt.Printf("Input: nums = %v\n", nums2)
	fmt.Printf("Output: %v\n", containsDuplicate(nums2))
	// Explanation: No element occurs more than once.
}
