// https://leetcode.com/problems/house-robber-ii/

// Concise
class Solution {
    vector<int> rob1(int i, bool first, vector<int>& nums){
        if((i==nums.size()-1 && first) || i==nums.size()){
            return {0,0};
        }
        auto exclude=rob1(i+1, first, nums);
        vector<int> res(2, 0);
        // exclude
        res[0]=max(exclude[0], exclude[1]);
        // include
        res[1]=exclude[0]+nums[i];
        return res;
    }
public:
    int rob(vector<int>& nums) {
        if(nums.size()==1)
            return nums[0];
        // robbery starts from house 1
        vector<int> r1=rob1(0, true, nums);
        // robbery starts from house 2
        vector<int> r2=rob1(1, false, nums);
        return max({r1[0], r1[1], r2[0], r2[1]});
    }
};

// easy to understand
class Solution {
    vector<int> rob1_(int i, vector<int>& nums){
        if(i==nums.size()-1)
            return {0,0};
        auto excl = rob1_(i+1, nums);
        vector<int> res(2,0);
        res[0]=max(excl[0], excl[1]); // exclude
        res[1]=nums[i]+excl[0]; //include
        cout<<i<<"->"<<res[0]<<" "<<res[1]<<"\n";
        return res;
    }
    vector<int> rob2_(int i, vector<int>& nums){
        if(i==nums.size())
            return {0,0};
        auto excl = rob2_(i+1, nums);
        vector<int> res(2,0);
        res[0]=max(excl[0], excl[1]); // exclude
        res[1]=nums[i]+excl[0]; //include
        cout<<i<<"->"<<res[0]<<" "<<res[1]<<"\n";
        return res;
    }
public:
    int rob(vector<int>& nums) {
        if(nums.size()==1)
            return nums[0];
        auto res1 = rob1_(0, nums);
        auto res2 = rob2_(1, nums);
        return max({res1[0], res1[1], res2[0], res2[1]});
    }
};

/*
    Time Complexity: O(n)
    Space Complexity: O(n)
*/