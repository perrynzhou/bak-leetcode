package main

import "fmt"

//link : https://leetcode-cn.com/problems/jump-game/
// solition: https://www.zhihu.com/column/p/93432919
/*
class Solution {
public:
    bool canJump(vector<int>& nums) {
        int n=nums.size();
        if(n<=1)//边界情况
            return true;
        int pos=n-1;//pos初始化为最后一个位置
        for(int i=n-2;i>=0;i--)//从后往前遍历
        {
            if(nums[i]+i>=pos)//如果可以到达，更新pos
                pos=i;
        }
        return pos==0;
    }
};
*/
func canJump(nums []int) bool {

	n, index := len(nums), len(nums)-1
	if n <= 1 {
		return true
	}
	for i := n - 2; i >= 0; i-- {
		if nums[i]+i >= index {
			index = i
		}
	}
	return index == 0
}
func main() {
	fmt.Println(canJump([]int{2, 3, 1, 1, 4}))
	fmt.Println(canJump([]int{30, 2, 1, 0, 4}))
	fmt.Println(canJump([]int{1}))
	fmt.Println(canJump([]int{2, 0}))

}
