class Solution(object):
    def reverse(self, x):
        """
        :type x: int
        :rtype: int
        """
        if x == -2147483648:
            return 0

        neg = x < 0
        if neg:
            x = x * -1
        
        output = ""
        x_str = str(x)
        for c in x_str:
            output = c + output
        

        might_be_too_big = len(output) == 10
        if might_be_too_big:
            MAX = "214748364"
            for i in range(len(MAX)):
                if int(output[i]) > int(MAX[i]):
                    return 0
                elif int(output[i]) < int(MAX[i]):
                    might_be_too_big = False
                    break

            if might_be_too_big:
                if int(output[-1]) > 8:
                    return 0
                if int(output[-1]) == 8 and neg:
                    return 0
                
        output = int(output)    
        if neg:
            output *= -1
        return output