package mergesortedarray

func merge(nums1 []int, m int, nums2 []int, n int) {
	for i, j, t := m-1, n-1, m+n-1; i >= 0 || j >= 0; t-- {
		var cur int
		if i == -1 {
			cur = nums2[j]
			j--
		} else if j == -1 {
			cur = nums1[i]
			i--
		} else if nums1[i] > nums2[j] {
			cur = nums1[i]
			i--
		} else {
			cur = nums2[j]
			j--
		}
		nums1[t] = cur
	}
}

func merge1(nums1 []int, m int, nums2 []int, n int) {
	sorted := make([]int, 0, m+n)
	i, j := 0, 0
	for {
		if i == m {
			sorted = append(sorted, nums2[j:]...)
			break
		}
		if j == n {
			sorted = append(sorted, nums1[i:]...)
			break
		}

		if nums1[i] < nums2[j] {
			sorted = append(sorted, nums1[i])
			i++
		} else {
			sorted = append(sorted, nums2[j])
			j++
		}
	}
	copy(nums1, sorted)
}
