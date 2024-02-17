package main

func twoSum(numbers []int, target int) []int {
	for i, j := range numbers {
		for l, m := range numbers {
			if j+m == target {
				return []int{i, l}
			}
		}
	}
	return []int{}
}
