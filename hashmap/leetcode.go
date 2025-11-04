package hashmap

func classic2Sum(nums []int, target int) []int {
	seen := make(map[int]int)
	for i, num := range nums {
		complement := target - num
		if j, found := seen[complement]; found {
			return []int{j, i}
		}
		seen[num] = i
	}
	return nil
}

func countFreq(nums []int) int {
	count := make(map[int]int)
	maxFreq := 0
	for _, val := range nums {
		count[val]++
		if count[val] > maxFreq {
			maxFreq = count[val]
		}
	}
	return maxFreq
}

func isAnagram(s, t string) bool {
	if len(s) != len(t) {
		return false
	}
	count := make(map[rune]int)
	for _, ch := range s {
		count[ch]++
	}
	for _, ch := range t {
		count[ch]--
		if count[ch] < 0 {
			return false
		}
	}
	return true
}
