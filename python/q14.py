class Solution:
    def longestCommonPrefix(self, strs: List[str]) -> str:
        output = ""
        base = strs[0]
        for i in range(len(base)):
            for s in strs:
                try:
                    if s[i] != base[i]:
                        return output
                except:
                    return output
            output += base[i]
        return output