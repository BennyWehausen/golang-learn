package main

import (
	"fmt"
	"golang-learn/task1"
)

func main() {
	//fmt.Println("Hello, World!")
	// 测试用例
	// testCases := []int{4, 1, 2, 1, 2}
	// // 运行测试用例
	// result :=  task1.SingleNumber(testCases)
	// fmt.Printf("测试结果： %d\n", result)

	//res :=  task1.Palindrome(121)
	//fmt.Printf("palindromec返回的结果： %d\n", res)
	//s := "()[]{}"
	//res :=  task1.ValidParentheses(s)
	//fmt.Printf("ValidParentheses返回的结果： %d\n", res)

	//var s = [3]string{"flower", "flow", "flight"}
	//res :=  task1.LongestCommonPrefix(s)
	//fmt.Printf("longestCommonPrefix返回的结果： %d\n", res)

	//var s = []int{1, 2, 3}
	//res :=  task1.PlusOne(s)
	//fmt.Printf("plusOne返回的结果： %d\n", res)
	//var s2 = []int{1, 2, 9}
	//res2 :=  task1.PlusOne(s2)
	//fmt.Printf("plusOne返回的结果2： %d\n", res2)

	//var s2 = []int{1, 1, 2, 2, 9}
	//res2 :=  task1.RemoveDuplicates(s2)
	//fmt.Printf("removeDuplicates返回的结果2： %d\n, array: %v \n", res2, s2[:res2])

	//var arr = [][]int{{1, 4}, {4, 5}}
	//res2 :=  task1.Merge(arr)
	//fmt.Printf("removeDuplicates返回的结果2： %d\n, array: %v \n", res2, arr)

	nums := []int{7, 11, 2, 15}
	target := 9
	result := task1.TwoSum(nums, target)
	fmt.Printf("Input: nums=%v, target=%d -> Output: %v\n,result:%v和%v", nums, target, result, nums[result[0]], nums[result[1]])
}
