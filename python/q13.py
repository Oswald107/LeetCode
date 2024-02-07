class Solution:
    NUM_MAP = {
        "I": 1,
        "V": 5,
        "X": 10,
        "L": 50,
        "C": 100,
        "D": 500,
        "M": 1000,
    }

    REV = {
        "I":["V","X"],
        "X":["L","C"],
        "C":["D","M"],
    }

    def romanToInt(self, s: str) -> int:
        output = 0
        prev =  ""
        for c in s:
            if prev in self.REV.keys() and c in self.REV[prev]:
                output -= 2 * self.NUM_MAP[prev]
            output += self.NUM_MAP[c]
            prev = c
        
        return output