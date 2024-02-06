class Solution(object):
    def convert(self, s, numRows):
        """
        :type s: str
        :type numRows: int
        :rtype: str
        """
        if numRows == 1:
            return s
        
        step = numRows * 2 - 2
        if numRows == 2:
            step = 2
        start = ""
        end = ""
        middle = ""

        for i in range(0, len(s), step):
            start += s[i]
            if len(s) > i + numRows - 1:
                end += s[i + numRows - 1]
        
        for row in range(1, numRows - 1):
            middle_step = (numRows - row - 1) * 2
            for i in range(row, len(s), step):
                middle += s[i]
                if len(s) > i + middle_step:
                    middle += s[i + middle_step]
                
        return start + middle + end