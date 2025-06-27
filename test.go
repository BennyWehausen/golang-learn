package main

import "fmt"

func main() {
	//fmt.Println("Hello, World!")
	// 测试用例
	testCases := []int{4, 1, 2, 1, 2}
	// 运行测试用例
	result := singleNumber(testCases)
	fmt.Printf("测试结果： %d\n", result)
}

/**
 * 任务1-梦的开始
 */
//136. 只出现一次的数字：
//给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。可以使用 for 循环遍历数组，
//结合 if 条件判断和 map 数据结构来解决，例如通过 map 记录每个元素出现的次数，然后再遍历 map 找到出现次数为1的元素。
func singleNumber(nums []int) int {
	// 创建一个map来记录每个数字出现的次数
	numCount := make(map[int]int)
	// 第一次遍历：统计每个数字出现的次数
	for _, num := range nums {
		numCount[num]++
	}
	// 第二次遍历：找出只出现一次的数字
	for num, count := range numCount {
		if count == 1 {
			return num
		}
	}
	// 如果没有找到只出现一次的元素（实际上根据题目描述这种情况不会发生）
	return -1
}
