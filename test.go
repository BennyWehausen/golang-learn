package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	//fmt.Println("Hello, World!")
	// 测试用例
	// testCases := []int{4, 1, 2, 1, 2}
	// // 运行测试用例
	// result := singleNumber(testCases)
	// fmt.Printf("测试结果： %d\n", result)

	//res := palindrome(121)
	//fmt.Printf("palindromec返回的结果： %d\n", res)
	//s := "()[]{}"
	//res := validParentheses(s)
	//fmt.Printf("ValidParentheses返回的结果： %d\n", res)

	//var s = [3]string{"flower", "flow", "flight"}
	//res := longestCommonPrefix(s)
	//fmt.Printf("longestCommonPrefix返回的结果： %d\n", res)

	//var s = []int{1, 2, 3}
	//res := plusOne(s)
	//fmt.Printf("plusOne返回的结果： %d\n", res)
	//var s2 = []int{1, 2, 9}
	//res2 := plusOne(s2)
	//fmt.Printf("plusOne返回的结果2： %d\n", res2)

	//var s2 = []int{1, 1, 2, 2, 9}
	//res2 := removeDuplicates(s2)
	//fmt.Printf("removeDuplicates返回的结果2： %d\n, array: %v \n", res2, s2[:res2])

	//var arr = [][]int{{1, 4}, {4, 5}}
	//res2 := merge(arr)
	//fmt.Printf("removeDuplicates返回的结果2： %d\n, array: %v \n", res2, arr)

	nums := []int{7, 11, 2, 15}
	target := 9
	result := twoSum(nums, target)
	fmt.Printf("Input: nums=%v, target=%d -> Output: %v\n,result:%v和%v", nums, target, result, nums[result[0]], nums[result[1]])
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

// 任务1-判断回文数
func palindrome(num int) bool {
	// 排除负数和以0结尾的非零数
	if num < 0 || num%10 == 0 && num != 0 {
		return false
	}
	// 反转整数
	reversedNum := 0
	originalNum := num
	for num > 0 {
		digit := num % 10
		reversedNum = reversedNum*10 + digit
		num = num / 10
	}
	// 比较反转前后的整数是否相等
	return originalNum == reversedNum
}

// 任务1-判断该字符串是否是有效括号字符串
func validParentheses(s string) bool {
	stack := []rune{}
	pairs := map[rune]rune{
		'(': ')',
		'{': '}',
		'[': ']'}
	for _, char := range s {
		if char == '(' || char == '{' || char == '[' {
			stack = append(stack, char)
		} else {
			if len(stack) == 0 {
				return false
			}
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if pairs[top] != char {
				return false
			}
		}
	}
	return len(stack) == 0
}

// 任务1-查找字符串数组中的最长公共前缀
func longestCommonPrefix(strs [3]string) string {
	if len(strs) == 0 {
		return ""
	}
	prefix := strs[0]
	for i := 1; i < len(strs); i++ {
		//检查字符串是否以指定前缀开头
		s := strs[i]
		for !strings.HasPrefix(s, prefix) {
			//通过 prefix[:len(prefix)-1] 逐步减少前缀长度，直到匹配或为空。
			prefix = prefix[:len(prefix)-1]
			if prefix == "" {
				return ""
			}
		}
	}
	return prefix
}

// 任务1-题目：给定一个由整数组成的非空数组所表示的非负整数，在该数的基础上加一
func plusOne(digits []int) []int {
	n := len(digits)
	// 从数组的最后一个元素开始向前遍历
	for i := n - 1; i >= 0; i-- {
		// 如果当前元素小于9，则将其加1，然后返回
		if digits[i] < 9 {
			digits[i]++
			return digits
		}
		// 如果当前元素等于9，则将其置为0
		digits[i] = 0
	}
	// 如果遍历完数组后，所有元素都等于9，则在数组开头添加一个1
	return append([]int{1}, digits...)
}

// 任务1-删除有序数组中的重复项
func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	i := 0
	for j := 1; j < len(nums); j++ {
		if nums[i] != nums[j] {
			i++
			nums[i] = nums[j]
		}
	}
	return i + 1
}

// 任务1-合并区间
func merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return [][]int{}
	}
	// 按照区间起始位置排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	merged := [][]int{intervals[0]}
	for i := 1; i < len(intervals); i++ {
		// 如果当前区间与上一个区间不重合，直接添加
		if merged[len(merged)-1][1] < intervals[i][0] {
			merged = append(merged, intervals[i])
		} else {
			// 否则的话，我们就可以与上一个区间进行合并
			merged[len(merged)-1][1] = max(merged[len(merged)-1][1], intervals[i][1])
		}
	}
	return merged
}

// 任务1-两数之和
func twoSum(nums []int, target int) []int {
	// 创建一个map来存储数值和对应的索引
	numMap := make(map[int]int)
	for i, num := range nums {
		// 计算需要的补数
		complement := target - num
		// 检查补数是否在map中
		if index, ok := numMap[complement]; ok {
			return []int{index, i}
		}
		// 将当前数值和索引存入map
		numMap[num] = i
	}
	// 如果没有找到解（根据题目描述这种情况不会发生）
	return nil
}
