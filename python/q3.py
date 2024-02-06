class Solution(object):
    def lengthOfLongestSubstring(self, s):
        """
        :type s: str
        :rtype: int
        """
        output = 0
        sub_count = 0
        substring = []
        for c in s:
            in_sub = c in substring
            substring.append(c)
            sub_count += 1

            if in_sub:
                while True:
                    sub_char = substring.pop(0)
                    sub_count -= 1
                    if sub_char == c:
                        break

            else:
                if sub_count > output:
                    output = sub_count

    
        return output
