class Solution:
    def fourSum(self, nums: List[int], target: int) -> List[List[int]]:
        n = len(nums)
        output = set()
        seen = set()

        for i in range(n-2):
            v1 = nums[i]
            for j in range(i+1, n-1):
                v2 = nums[j]
                for k in range(j+1, n):
                    v3 = nums[k]
                    v4 = target - (v1 + v2 + v3)
                    if v4 in seen:
                        valid = sorted([v1, v2, v3, v4])
                        output.add((valid[0], valid[1], valid[2], valid[3]))
            seen.add(nums[i])

        return output 

