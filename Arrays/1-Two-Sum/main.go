package leetcode

func TwoSum(nums []int, target int) []int {
	hash := make(map[int]int)

	for i, v := range nums {
		if _, ok := hash[v]; ok {
			return []int{hash[v], i}
		}
		hash[target-v] = i
	}
	return nil
}
