package main

//link:https://leetcode-cn.com/problems/combination-sum/

func combinationSum(candidates []int, target int) [][]int {
	result :=[][]int{}
	l,r :=0,len(candidates)-1
	dict :=map[int]struct{}{}
	for i:=0;i<=r;i++ {
       dict[candidates[i]] = struct{}{}
	}
	for i:=0;i<=r;i++ {
		 sum :=target
		 for {
		 	 if sum <0 {
		 	 	break
			 }
			 sum -= candidates[i]
			 
		 }
	}
}
