class Solution {
    public int mySqrt(int x) {
        int i = 0;
        while (x>=0) {
            x -= (2*i+1);
            i++;
        }
        return i-1;
    }
}