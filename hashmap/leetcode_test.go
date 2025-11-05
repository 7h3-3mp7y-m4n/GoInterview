package hashmap

import (
	"reflect"
	"testing"
)

func TestClassic2Sum(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		target   int
		expected []int
	}{
		{"basic case", []int{2, 7, 11, 15}, 9, []int{0, 1}},
		{"multiple pairs", []int{3, 2, 4}, 6, []int{1, 2}},
		{"no match", []int{1, 2, 3}, 7, nil},
		{"duplicate values", []int{3, 3}, 6, []int{0, 1}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := classic2Sum(tt.nums, tt.target)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}

// --- Tests for countFreq ---
func TestCountFreq(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{"all unique", []int{1, 2, 3, 4}, 1},
		{"some duplicates", []int{1, 2, 2, 3, 3, 3}, 3},
		{"single element", []int{5}, 1},
		{"empty slice", []int{}, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := countFreq(tt.nums)
			if result != tt.expected {
				t.Errorf("expected %d, got %d", tt.expected, result)
			}
		})
	}
}

// ---- Test anagram
func TestIsAnagram(t *testing.T) {
	tests := []struct {
		name     string
		s, t     string
		expected bool
	}{
		{"simple anagram", "listen", "silent", true},
		{"not anagram", "hello", "world", false},
		{"different lengths", "abc", "ab", false},
		{"same letters diff count", "aabb", "ab", false},
		{"speical chars", "1$$", "1$$", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := isAnagram(tt.s, tt.t)
			if result != tt.expected {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}

func TestSubarraySum(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		k        int
		expected int
	}{
		{
			name:     "single match",
			nums:     []int{1, 1, 1},
			k:        2,
			expected: 2,
		},
		{
			name:     "no subarray sums to k",
			nums:     []int{1, 2, 3},
			k:        7,
			expected: 0,
		},
		{
			name:     "multiple matches",
			nums:     []int{1, 2, 3},
			k:        3,
			expected: 2, // subarrays [1,2] and [3]
		},
		{
			name:     "negative numbers",
			nums:     []int{1, -1, 0},
			k:        0,
			expected: 3, // subarrays: [1,-1], [0], [1,-1,0]
		},
		{
			name:     "single element equal to k",
			nums:     []int{3},
			k:        3,
			expected: 1,
		},
		{
			name:     "empty array",
			nums:     []int{},
			k:        0,
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := subarraySum(tt.nums, tt.k)
			if result != tt.expected {
				t.Errorf("for input nums=%v, k=%d: expected %d, got %d", tt.nums, tt.k, tt.expected, result)
			}
		})
	}
}
