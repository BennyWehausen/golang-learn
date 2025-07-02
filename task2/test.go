package task2

import (
	"fmt"
	"golang-learn/common"
	"time"
)

func Test() {
	a := 2
	var p *int
	fmt.Println(&a)
	p = &a
	fmt.Println(p, &a)

	var pp **int
	pp = &p
	fmt.Println(pp, p)
	**pp = 3
	fmt.Println(pp, *pp, p)
	fmt.Println(**pp, *p)
	fmt.Println(a, &a)
}

/*指针*/
// 任务2-1编写一个Go程序，定义一个函数，该函数接收一个整数指针作为参数，在函数内部将该指针指向的值增加10，然后在主函数中调用该函数并输出修改后的值
func AddTen(num *int) {
	// addTen 函数接收一个整数指针，将其指向的值增加 10
	*num += 10 // 解引用指针并增加 10
}

//	func AddTen2(num int) int {
//		// addTen 函数接收一个整数指针，将其指向的值增加 10
//		return num + 10 // 解引用指针并增加 10
//	}
//
// 任务2-2实现一个函数，接收一个整数切片的指针，将切片中的每个元素乘以2。
func DoubleSlice(sli *[]int) {
	// 遍历切片中的每个元素，将其乘以 2
	for i := range *sli {
		(*sli)[i] *= 2
	}
}

// 打印奇数
func PrintOdd() {
	for i := 1; i <= 10; i += 2 {
		fmt.Println(i)
	}
	time.Sleep(1 * time.Second)
}

// 打印偶数
func PrintEven() {
	for i := 2; i <= 10; i += 2 {
		fmt.Println(i)
	}
	time.Sleep(2 * time.Second)
}

// 打印奇数
func PrintOdd2() {
	defer common.Wg.Done() // Goroutine 完成时调用 Done()
	for i := 1; i <= 10; i += 2 {
		fmt.Println(i)
	}
}

// 打印偶数
func PrintEven2() {
	defer common.Wg.Done() // Goroutine 完成时调用 Done()
	for i := 2; i <= 10; i += 2 {
		fmt.Println(i)
	}
}

// 任务调度器函数
func ScheduleTasks(tasks []func()) {
	var wg = common.Wg
	resultChan := make(chan common.TaskResult, len(tasks)) // 缓冲通道

	for i, task := range tasks {
		wg.Add(1)
		go func(taskID int, f func()) {
			defer wg.Done()
			start := time.Now()
			f()
			dur := time.Since(start)
			resultChan <- common.TaskResult{ID: taskID, Duration: dur}
		}(i, task)
	}

	// 等待所有任务完成
	wg.Wait()

	// 关闭结果通道，防止死锁
	close(resultChan)

	// 收集所有任务执行结果
	var results []common.TaskResult
	for res := range resultChan {
		results = append(results, res)
	}

	// 输出任务执行时间
	for _, r := range results {
		fmt.Printf("Task %d took %v\n", r.ID, r.Duration)
	}
}
