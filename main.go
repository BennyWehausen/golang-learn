package main

import (
	"fmt"
	"golang-learn/task1"
	"golang-learn/task2"
	"time"
)

func main() {
	//testTask1()
	testTask2()
}

func testTask2() {
	//task2.Test()
	/*	var num int = 5
		fmt.Println("原始值:", num) // 输出原始值
		//rs := task2.AddTen2(num)
		//fmt.Println("修改后的值2:", num, rs)
		// 传递变量的地址给函数
		task2.AddTen(&num)
		fmt.Println("修改后的值:", num)*/
	//nums := []int{1, 2, 3}
	//fmt.Println("原始切片:", nums)
	//// 调用函数，传递切片的指针
	//task2.DoubleSlice(&nums)
	//fmt.Println("修改后的切片:", nums)
	//testGoFunc()
	testGoroutine()
}

// 同时启动两个goroutine来打印奇数和偶数，然后主程序等待一段时间以确保输出完成。
func testGoFunc() {
	go task2.PrintOdd()
	go task2.PrintEven()
	time.Sleep(1 * time.Second)
}

/*同步多个 Goroutine*/
func testGoroutine() {
	//common.Wg.Add(2)
	//// 启动 Goroutine
	//go task2.PrintOdd2()
	//go task2.PrintEven2()
	//// 等待所有 Goroutine 完成
	//common.Wg.Wait()

	// 定义任务列表
	tasks := []func(){
		task2.PrintOdd,
		task2.PrintEven,
	}
	// 调度任务
	task2.ScheduleTasks(tasks)
}
func testTask1() {
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
