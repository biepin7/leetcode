class Solution {
public:
    int trailingZeroes(int n) {
        int div = 5;
        int res = 0;

        while (div <= n) {

            res += n / div;
            div *= 5;
        }

        return res;
    }
};