package main_test

import (
	"fmt"
	"testing"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1) > len(nums2) {
		return findMedianSortedArrays(nums2, nums1)
	}

	n, m := len(nums1), len(nums2)
	left, right := 0, n
	halfLen := (n + m + 1) / 2

	for left <= right {
		i := (left + right) / 2
		j := halfLen - i

		if i < n && nums2[j-1] > nums1[i] {
			left = i + 1
		} else if i > 0 && nums1[i-1] > nums2[j] {
			right = i - 1
		} else {
			var maxLeft int
			if i == 0 {
				maxLeft = nums2[j-1]
			} else if j == 0 {
				maxLeft = nums1[i-1]
			} else {
				maxLeft = max(nums1[i-1], nums2[j-1])
			}

			if (n+m)%2 == 1 {
				return float64(maxLeft)
			}

			var minRight int
			if i == n {
				minRight = nums2[j]
			} else if j == m {
				minRight = nums1[i]
			} else {
				minRight = min(nums1[i], nums2[j])
			}

			return float64(maxLeft+minRight) / 2.0
		}
	}
	return 0.0
}

func TestFindMedianSortedArrays(t *testing.T) {
	tests := []struct{
		nums1 []int
		nums2 []int
		expect float64
	}{
		{
			nums1: []int{1,3},
			nums2: []int{2},
			expect: 2,
		},
		{
			nums1: []int{1,2},
			nums2: []int{3,4},
			expect: 2.5,
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("test. nums1: %v, nums2: %v", tt.nums1, tt.nums2), func(t *testing.T) {
			r := findMedianSortedArrays(tt.nums1, tt.nums2)
			if r != tt.expect {
				t.Errorf("test failed. expect: %f, actual: %f", tt.expect, r)
			}
		})
	}
}

