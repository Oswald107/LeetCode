class Solution:
    def letterCombinations(self, digits: str) -> List[str]:
        m = {
            '2': ['a', 'b', 'c'],
            '3': ['d', 'e', 'f'],
            '4': ['g', 'h', 'i'],
            '5': ['j', 'k', 'l'],
            '6': ['m', 'n', 'o'],
            '7': ['p', 'q', 'r', 's'],
            '8': ['t', 'u', 'v'],
            '9': ['w', 'x', 'y', 'z'],
        }
        if digits == "":
            return []

        output = m[digits[0]]
        for i in range(1, len(digits)):
            c = digits[i]
            temp = []
            for val in output:
                for letter in m[c]:
                    temp.append(val + letter)
            output = temp

        return output