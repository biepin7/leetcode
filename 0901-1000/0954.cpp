class Solution
{
public:
    bool canReorderDoubled(vector<int> &arr)
    {
        unordered_map<int, int> m;

        // arr 塞到 m 里 key：arr[i]，value：出现次数
        for (int i : arr)
        {
            m[i]++;
        }

        // 特殊：0*2 = 0 所以要判断 0 是否为偶数个
        if (m[0]%2)
        {
            return false;
        }

        //从 m 里取出所有的 key
        vector<int> keys;
        keys.reserve(m.size());
        for (auto &[k, _] : m)
        {
            keys.push_back(k);
        }

        // 排序 以绝对值大小从小到大排序，为啥：贪心的来说从0左右开始判断
        // ps : 使用 lambda ，进行一个加速，这样就会用时超过 100%
        sort(keys.begin(), keys.end(),[](int a,int b){return abs(a)<abs(b);});

        for(int x : keys){
            if(m[x]>m[2*x]) return false;  // 不可用 ！=
            m[2*x] -= m[x]; // 做完处理后要抵消 m[x*2]
        }
        return true;
    }
};