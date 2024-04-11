class Solution:
    def findMedianSortedArrays(self, nums1: List[int], nums2: List[int]) -> float:
        if len(nums1) > len(nums2):
            temp = nums1
            nums1 = nums2
            nums2 = temp
        l1, l2 = len(nums1), len(nums2)

        total = l1 + l2
        mid = (total+1) // 2

        low, high = 0, l1
        while True:
            mid1 = (low + high) // 2
            mid2 = mid - mid1
            x1, x2, y1, y2 = float('-inf'),float('-inf'),float('inf'),float('inf')
            if mid1-1 >= 0: x1 = nums1[mid1-1]
            if mid2-1 >= 0: x2 = nums2[mid2-1]
            if mid1 < l1:   y1 = nums1[mid1]
            if mid2 < l2:   y2 = nums2[mid2]
            
            if x1 > y2: 
                high = mid1 - 1
            elif x2 > y1:
                low = mid1 + 1
            else:
                return max(x1, x2) if total % 2 else (max(x1, x2) + min(y1, y2)) / 2.0