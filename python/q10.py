class Solution:
    def isMatch(self, s: str, p: str) -> bool:
        return recur(s, p)

    

def recur(s, p):
    while p:        
        next_pattern = p[-1]
        if not s and next_pattern != '*':
            return False

        if next_pattern == '*':
            star_char = p[-2]
            for i in range(len(s)+1):
                if star_char != '.' and s[len(s)-i:] != star_char*i:
                    return False
                if recur(s[:len(s)-i], p[:-2]):
                    return True
            return False

        next_character = s[-1]
        if next_pattern != next_character and next_pattern != '.':
            return False

        p = p[:-1]
        s = s[:-1]

    if s:
        return False
    return True