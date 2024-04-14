package __median_of_two_sorted_arrays

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	totalLen := len(nums1) + len(nums2)
	if totalLen&1 == 1 {
		return float64(getKthElem(nums1, nums2, totalLen/2+1))
	}
	midLen1, midLen2 := totalLen/2-1, totalLen/2
	return float64(getKthElem(nums1, nums2, midLen1+1)+getKthElem(nums1, nums2, midLen2+1)) / 2.0
}

func getKthElem(n1, n2 []int, k int) int {
	idx1, idx2 := 0, 0
	for {
		if len(n1) == idx1 {
			return n2[idx2+k-1]
		}
		if len(n2) == idx2 {
			return n1[idx1+k-1]
		}
		if k == 1 {
			return min(n1[idx1], n2[idx2])
		}
		half := k / 2
		nidx1 := min(idx1+half, len(n1)) - 1
		nidx2 := min(idx2+half, len(n2)) - 1
		pivot1, pivot2 := n1[nidx1], n2[nidx2]
		if pivot1 <= pivot2 {
			k -= nidx1 - idx1 + 1
			idx1 = nidx1 + 1
		} else {
			k -= nidx2 - idx2 + 1
			idx2 = nidx2 + 1
		}
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
