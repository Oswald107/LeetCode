func merge(nums1 []int, m int, nums2 []int, n int) {
	for {
		if 0 == n {
			break
		}
		if 0 == m || nums1[m-1] <= nums2[n-1] {
			nums1[n+m-1] = nums2[n-1]
			n -= 1
		} else {
			nums1[n+m-1] = nums1[m-1]
			m -= 1
		}
	}
}