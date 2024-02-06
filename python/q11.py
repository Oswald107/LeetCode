class Solution(object):
    def maxArea(self, height):
        """
        :type height: List[int]
        :rtype: int
        """
        i = 0
        j = len(height) - 1
        output = 0
        while j > i:
            area = (j - i) * min(height[i], height[j])
            output = max(area, output)
            if height[j] < height[i]:
                j -= 1
            else:
                i += 1
        return output