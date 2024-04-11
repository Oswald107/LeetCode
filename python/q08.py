class Solution:
    def myAtoi(self, s: str) -> int:
        n = len(s)
        i = 0
        while i < n and s[i] == ' ':
            i+=1
        
        if i == n:
            return 0

        negative = False
        if s[i] == '-':
            negative = True
            i+=1
        elif s[i] == '+':
            i+=1

        output = 0
        while i < n:
            c = s[i]
            if c < '0' or c > '9':
                break
            output = output*10 + int(c)
            i+=1
        
        if negative:
            output = -output

        if output > 2**31 - 1:
            return 2**31 - 1
        
        if output < -2**31:
            return -2**31
        return int(output)
