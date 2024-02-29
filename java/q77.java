class Solution {
    public List<List<Integer>> combine(int n, int k) {
        List<List<Integer>> output  = new ArrayList<List<Integer>>();
        recur(output, new ArrayList<Integer>(), n, k, 0);
        return output;
    }

    public void recur(List<List<Integer>> output, List<Integer> cur, int n, int k, int ind) {
        if (k == 0) {
            output.add(new ArrayList<>(cur));
            return;
        }


        for (int i = ind+1; i <= n; i++) {
            if (k > n-ind)
                return;
            cur.add(i);
            recur(output, cur, n, k-1, i);
            cur.remove(cur.size()-1);
        }
    }
}