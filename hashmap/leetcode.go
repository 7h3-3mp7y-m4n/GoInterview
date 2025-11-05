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

// Prefix Sum ???
func preFixSum(num []int) ([]int, int) {
	if len(num) == 0 {
		return []int{}, 0
	}
	temp := make([]int, len(num))
	sum := 0

	for i, val := range num {
		sum += val
		temp[i] = sum
	}
	return temp, sum
}

func subarraySum(nums []int, k int) int {
	count := 0
	prefixSum := 0
	freq := make(map[int]int)
	freq[0] = 1 // base case
	for _, num := range nums {
		prefixSum += num
		// check if prefixSum - k has been seen
		if val, exists := freq[prefixSum-k]; exists {
			count += val
		}

		freq[prefixSum]++
	}

	return count
}
