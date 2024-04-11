class Solution:
    def threeSum(self, nums: List[int]) -> List[List[int]]:
        n, p = [],[]
        zero = 0
        for i in nums:
            if i > 0:
                p.append(i)
            elif i < 0:
                n.append(i)
            else:
                zero += 1
        
        N_set, P_set = set(n), set(p)

        output = set()
        if zero:
            for v in p:
                if -v in N_set:
                    output.add((-v, 0, v))
            if zero >= 3:
                output.add((0,0,0))

        for i in range(len(n)):
            for j in range(i+1,len(n)):
                target = -1*(n[i]+n[j])
                if target in P_set:
                    output.add(tuple(sorted([n[i],n[j],target])))
        for i in range(len(p)):
            for j in range(i+1,len(p)):
                target = -1*(p[i]+p[j])
                if target in N_set:
                    output.add(tuple(sorted([p[i],p[j],target])))
        
        return output