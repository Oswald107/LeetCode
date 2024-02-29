class Solution {
    public int removeDuplicates(int[] nums) {
        int[] prev = new int[2];
        int k = 0; 
        for (int i = 0; i < nums.length; i++) {
            if (i < 2  || prev[0] != nums[i]) {
                nums[k] = nums[i];
                k += 1;
            }
            prev[0] = prev[1];
            prev[1] = nums[i];
        }
        return k;
    }
}