class Solution {
    public void sortColors(int[] nums) {
        int zeros = 0;
        int end = nums.length-1;
        int i = 0;

        while (i <= end) {
            if (nums[i] == 2) { 
                int temp = nums[i];
                nums[i] = nums[end];
                nums[end] = temp;
                end--;
            } else if (i > zeros && nums[i] == 0) {
                int temp = nums[i];
                nums[i] = nums[zeros];
                nums[zeros] = temp;
                zeros++;
            } else {
                i++;
            }
        }
    }
}