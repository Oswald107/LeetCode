class Solution:
    def isValid(self, s: str) -> bool:
        OPEN = ['(', '{', '[']
        CLOSED = {'(':')', '{':'}', '[':']'}
        stack = []
        for c in s:
            if c in OPEN:
                stack.append(c)
            elif not stack or c != CLOSED[stack.pop()]:
                return False
        if stack:
            return False
        return True