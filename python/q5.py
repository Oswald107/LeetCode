class Solution(object):
    def longestPalindrome(self, s):
        """
        :type s: str
        :rtype: str
        """
        # make a map of characters to indices
        char_dict = {}
        for i in range(len(s)):
            c = s[i]
            if c not in char_dict:
                char_dict[c] = []
            char_dict[c].append(i)

        output = s[0]
        print(char_dict)
        for char_lst in char_dict.values():
            done_with_char = False
            for y in range(len(char_lst)-1):
                i = char_lst[y]
                for z in range(len(char_lst)-1):
                    j = char_lst[-(z+1)]
                    print(i, j)
                    if j - i + 1 <= len(output):
                        break
                    substr = s[i:j+1]
                    if isPal(substr):
                        output = substr
                        break

        return output

def isPal(s):
    for i in range(len(s)//2):
        if s[i] != s[len(s) - i - 1]:
            return False
    return True