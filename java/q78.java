class Solution {
    public List<List<Integer>> subsets(int[] nums) {
        List<List<Integer>> output  = new ArrayList<List<Integer>>();
        recur(output, new ArrayList<Integer>(), nums, 0);
        return output;
    }

    public void recur(List<List<Integer>> output, List<Integer> cur, int[] nums, int ind) {
        output.add(new ArrayList<>(cur));
        if (nums.length == 0) {
            return;
        }

        for (int i = ind; i < nums.length; i++) {
            cur.add(nums[i]);
            recur(output, cur, nums, i+1);
            cur.remove(cur.size()-1);
        }
    }
}