class Solution {
public:
    int calPoints(vector<string>& ops) {
        int srk[1001],tt=0,res=0;
        for(auto o:ops){
            if(o=="C"){
                tt--;
            }
            else if(o=="D"){
                srk[tt]=srk[tt-1]*2;
                tt++;
            }
            else if(o=="+"){
                srk[tt]=srk[tt-1]+srk[tt-2];
                tt++;
            }
            else{
                srk[tt]=stoi(o);
                tt++;
            }
        }
        for(int i=0;i<tt;i++){
            res+=srk[i];
        }
        return res;
    }
};