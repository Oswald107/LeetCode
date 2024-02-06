class Solution(object):
    def intToRoman(self, num):
        """
        :type num: int
        :rtype: str
        """
        num_map = {
            1: "I",
            5: "V",
            10: "X",
            50: "L",
            100: "C",
            500: "D",
            1000: "M",
        }

        output = ""
        cur_num = 1000
        cur_div = 2
        while num > 0:
            val = num // cur_num
            # 400, 40, 4
            if val == 4:
                output += num_map[cur_num] + num_map[cur_num * 5]
                num %= cur_num
            # 900, 90, 9
            elif cur_div == 5 and num // (cur_num / cur_div) == 9:
                output +=  num_map[cur_num/5] + num_map[cur_num*2]
                num %= (cur_num / cur_div)
            # all other numbers
            else:
                output += num_map[cur_num] * int(val)
                num %= cur_num

            # update
            cur_num = int(cur_num / cur_div)
            cur_div = 5 if cur_div == 2 else 2

        return output



print(Solution().intToRoman(1994))