package two_sum

import "fmt"

// https://leetcode.com/problems/two-sum/description/
func main() {
	fmt.Println("Pattern: two pointers")
	fmt.Println("Question: two sum 0001 e")

}

// TwoSum returns the indices of the two numbers that add up to the target.
func TwoSum(nums []int, target int) []int {
	seen := make(map[int]int) // Store number → index

	for i, num := range nums {
		complement := target - num
		if j, exists := seen[complement]; exists {
			return []int{j, i}
		}
		seen[num] = i
	}
	return nil
}
