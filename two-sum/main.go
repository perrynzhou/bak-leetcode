package main

import "fmt"

func twoSum(nums []int, target int) []int {
	indexes := make([]int, 2)
	mp := make(map[int][]int)
	for i, v := range nums {
		if mp[v] == nil {
			mp[v] = make([]int, 0)
		}
		mp[v] = append(mp[v], i)
	}
	//处理一个数组中相同的元素
	mid := target / 2
	if target%2 == 0 && len(mp[mid]) == 2 {
		indexes[0] = mp[mid][0]
		indexes[1] = mp[mid][1]
		return indexes
	}
	//处理数组中target减去另外一个元素没有重复的情况
	for i, v := range nums {
		rest := target - v
		if rest != v && mp[rest] != nil {
			indexes[0] = i
			indexes[1] = mp[rest][0]
			break
		}
	}
	return indexes
}
func main() {
	fmt.Println(twoSum([]int{3, 2, 4}, 6))
}
