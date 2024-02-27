class Solution {
    public int[] plusOne(int[] digits) {
        int i = digits.length-1;
        while (i >= 0) {
            digits[i] = (digits[i] + 1) % 10;
            if (digits[i] != 0) {
                break;
            }
            i--;
        }

        if (digits[0] == 0) {
            int[] output = new int[digits.length+1];
            output[0] = 1;
            for (i = 0; i < digits.length; i++)
                output[i+1] = digits[i];
            return output;
        }
            

        return digits;
    }
}