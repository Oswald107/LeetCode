class Solution:
    def threeSumClosest(self, nums: List[int], target: int) -> int:
        nums.sort()
        output = float('inf')
        n = len(nums)
        for i1 in range(n-2):
            i2 = i1+1
            i3 = n-1
            v1 = nums[i1]
            v2 = nums[i2]
            v3 = nums[i3]
            while i2 < i3:
                total = v1 + v2 + v3
                if abs(target - total) < abs(target - output):
                    output = total
                if total == target:
                    return target
                elif total < target:
                    i2 += 1
                    v2 = nums[i2]
                else:
                    i3 -= 1
                    v3 = nums[i3]
        
        return int(output)