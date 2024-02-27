class Solution {
    public String addBinary(String a, String b) {
        int i = a.length()-1;
        int j = b.length()-1;
        if (j > i) {
            String temp_str = a;
            int temp_len = i;
            a = b;
            b = temp_str;
            i = j;
            j = temp_len;
        }

        int carry = 0;
        String output = "";
        while (i >= 0) {
            int val = carry + Character.getNumericValue(a.charAt(i));
            if (j >= 0)
                val += Character.getNumericValue(b.charAt(j));
            
            carry = 0;
            if (val > 1)
                carry = 1;
            
            output = val%2 + output;
            
            i--;
            j--;
        }
        
        if (carry == 1) {
            output = "1" + output;
        }

        return output;
    }
}