class Solution {
public:
    vector<int> missingRolls(vector<int>& rolls, int mean, int n) {
        // 求 m+n 之和
        int sum = mean*(n+rolls.size());
        
        // sum 减去 m 
        for(int i=0;i<rolls.size();i++) 
            sum -= rolls[i];

        // 边界条件
        if(sum<n || sum > 6*n) return {};

        // 定义 res ，均分田地 PS：使用贪心来分配
        vector<int> res( n,sum/n);

        // 讲均分后剩余的数，贪心分配
        int rest = sum%n;
        for (int i = 0; i < rest; i++)
        {
            res[i]++;
        }
        
        return res;
    }
};