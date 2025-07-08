package main

import (
	"fmt"
	"golang-learn/common"
	"golang-learn/dao"
	"golang-learn/task1"
	"golang-learn/task2"
	"golang-learn/task3"
	"gorm.io/gorm"
	"strconv"
	"sync"
	"time"
)

func main() {
	//fmt.Printf("TAG=", foo.TAG)
	//interfaceTest()
	/*Task1 work */
	//testTask1()
	/*Task2 work */
	/*Goroutine work */
	//testTask2()
	/*面向对象 work*/
	//testShape()
	//printInfoTest()
	/*Channel work*/
	//channelTest()
	//channelValueTest()
	/*锁机制 work */
	//safeCounterTest()
	//task2.SafeCounterAtomic()
	//sqlxTest()
	//modelSetTest()
	//gormQueryTest(5)
	hookFuncTest()
}

func hookFuncTest() {
	// 示例：创建一个用户
	user := task3.User{Name: "mjz44", Email: "mjz44@example.com", Password: "123"}
	dao.DB.Create(&user)

	// 创建一篇文章
	post := task3.Post{Title: "Hello World44", Content: "This is my first post44.", UserID: user.ID}
	dao.DB.Create(&post)

	// 创建两条评论
	comment1 := task3.Comment{Content: "Great post44!", PostID: post.ID}
	//comment2 := task3.Comment{Content: "Nice article33.", PostID: post.ID}
	dao.DB.Create(&comment1)
	//dao.DB.Create(&comment2)

	// 删除其中一条评论
	dao.DB.Delete(&comment1)

	// 再次查询 post 查看评论状态是否变化
	var p task3.Post
	dao.DB.Preload("Comments").First(&p, post.ID)
	fmt.Printf("Post Title: %s\n", p.Title)
	fmt.Printf("Comment Status: %s\n", p.CommentStatus)
}

func gormQueryTest(userId uint) {
	var user task3.User
	dao.DB.Preload("Posts.Comments").First(&user, userId)
	fmt.Printf("userId %v 用户发布的所有文章及其对应的评论信息", userId)
	fmt.Printf("User: %s\n", user.Name)
	for _, post := range user.Posts {
		fmt.Printf("文章信息-Post: %s\n", post.Title)
		for _, comment := range post.Comments {
			fmt.Printf("评论信息-Comment: %s\n", comment.Content)
		}
	}
	getPostWithMostComments(dao.DB)
}

// 使用原生 SQL 查询评论最多的帖子
func getPostWithMostComments(db *gorm.DB) {
	type Result struct {
		PostID   uint
		Title    string
		Comments int64
	}

	var result Result

	// 查询并按评论数排序
	db.Table("comments").
		Select("posts.id as post_id, posts.title as title, COUNT(comments.id) as comments").
		Joins("JOIN posts ON posts.id = comments.post_id").
		Group("comments.post_id").
		Order("comments DESC").
		Limit(1).
		Scan(&result)
	fmt.Printf("评论数量最多的文章信息:")
	fmt.Printf("Post with most comments:\n")
	fmt.Printf("Title: %s\n", result.Title)
	fmt.Printf("Total Comments: %d\n", result.Comments)
}

func modelSetTest() {
	// 自动迁移模型，创建对应的数据库表
	dao.DB.AutoMigrate(&task3.User{}, &task3.Post{}, &task3.Comment{})
}

/*
	func sqlxTest() {
		// 查询1: 获取技术部所有员工
		techEmployees, err := task3.GetEmployeesByDepartment(dao.DB, "技术部")
		if err != nil {
			fmt.Printf("查询技术部员工失败: %v", err)
		} else {
			fmt.Println("技术部员工:")
			for _, emp := range techEmployees {
				fmt.Printf("ID:%d Name:%s Salary:%.2f\n", emp.ID, emp.Name, emp.Salary)
			}
		}
		//查询工资最高的员工信息
		employeeInfo := task3.EmployeeInfo{}
		employeeInfo, _ = task3.GetHighestPaidEmployee(dao.DB)
		// 打印结果
		fmt.Printf("最高工资员工: ID: %d, Name: %s, Department: %s, Salary: %.2f\n",
			employeeInfo.ID, employeeInfo.Name, employeeInfo.Department, employeeInfo.Salary)
		defer dao.DB.Close()

		// 查询价格大于50的书籍
		books, err := task3.QueryExpensiveBooks(dao.DB, 50)
		if err != nil {
			panic(err)
		}
		fmt.Println("价格大于 50 元的书籍:")
		// 打印结果
		for _, book := range books {
			fmt.Printf("%d: %s (%s) ￥%.2f\n",
				book.ID, book.Title, book.Author, book.Price)
		}
	}
*/
func safeCounterTest() {
	// 创建带互斥锁的计数器
	counter := task2.SafeCounter{}
	var wg sync.WaitGroup
	wg.Add(10) // 等待10个goroutine
	// 启动10个goroutine
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			// 每个goroutine执行1000次递增
			for j := 0; j < 1000; j++ {
				counter.Increment()
			}
		}()
	}
	wg.Wait() // 等待所有goroutine完成
	// 输出最终结果
	fmt.Printf("最终计数器值: %d\n", counter.Count)
}
func printInfoTest() {
	// 创建一个Employee实例
	emp := task2.Employee{
		Person: task2.Person{
			Name: "张三",
			Age:  33,
		},
		EmployeeID: "p0187330",
	}
	// 调用方法打印信息
	emp.PrintInfo()
}
func testShape() {
	r := task2.Rectangle{Width: 3, Height: 4}
	fmt.Printf("Rectangle area:%.2f\n", r.Area())
	fmt.Printf("Rectangle Perimeter:%.2f\n", r.Perimeter())
	c := task2.Circle{Radius: 5}
	fmt.Printf("Circle area:%.2f\n", c.Area())
	fmt.Printf("Circle Perimeter:%.2f\n", c.Perimeter())
}

func interfaceTest() {
	creditCard := &CreditCard{balance: 0, limit: 1000}
	debitCard := &DebitCard{balance: 500}

	fmt.Println("使用信用卡购买:")
	purchaseItem(creditCard, 800)

	fmt.Println("\n使用借记卡购买:")
	purchaseItem(debitCard, 300)

	fmt.Println("\n再次使用借记卡购买:")
	purchaseItem(debitCard, 300)
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
	//testGoroutine()
	// 方式1，声明并初始化一个空的切片
	/*	var s1 []int
		fmt.Println(s1)
		s1 = append(s1, 1, 2)
		fmt.Println("append:", s1)
	*/
	//a := [5]int{6, 5, 4, 3, 2}
	//fmt.Println(a)
	//// 从数组下标2开始，直到数组的最后一个元素
	//s7 := a[2:]
	//fmt.Println(s7)
	//s8 := a[1:3]
	//// 从0到下标2的元素，创建一个新的切片
	//s9 := a[:2]
	//fmt.Println(s8)
	//fmt.Println(s9)
	//a[0] = 9
	//a[1] = 8
	//a[2] = 7
	//fmt.Println(s7)
	//fmt.Println(s8)
	//fmt.Println(s9)
	//s := make([]int, 2, 2)
	//s[0] = 10
	//s[1] = 20
	//fmt.Println("initial, s =", s)
	//s2 := append(s, 4, 5, 6, 7, 8, 9)
	//s2[0] = 100
	//s2[1] = 200
	//fmt.Println("initial, s2 =", s2)

	//str1 := "a1中文"
	//for index, value := range str1 {
	//	fmt.Printf("str1 -- index:%d, index value:%d\n", index, str1[index])
	//	fmt.Printf("str1 -- index:%d, range value:%d\n", index, value)
	//}
	//var i int32 = 17
	//var b byte = 5
	//var f float32
	//
	//// 数字类型可以直接强转
	//f = float32(i) / float32(b)
	//fmt.Printf("f 的值为: %f\n", f)
	//
	//// 当int32类型强转成byte时，高位被直接舍弃
	//var i2 int32 = 256
	//var b2 byte = byte(i2)
	//fmt.Printf("b2 的值为: %d\n", b2)

	str := "123"
	num, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}
	fmt.Printf("字符串转换为int: %d \n", num)
	str1 := strconv.Itoa(num)
	fmt.Printf("int转换为字符串: %s \n", str1)

	ui64, err := strconv.ParseUint(str, 10, 32)
	fmt.Printf("字符串转换为uint64: %d \n", num)

	str2 := strconv.FormatUint(ui64, 2)
	fmt.Printf("uint64转换为字符串: %s \n", str2)
}

// 只发送channel的函数
func sendOnlyTen(ch chan<- int) {
	for i := 0; i < 10; i++ {
		ch <- i
		fmt.Printf("发送: %d\n", i)
	}
	close(ch)
}

// 只接收channel的函数
func receiveOnlyTen(ch <-chan int) {
	for v := range ch {
		fmt.Printf("接收到: %d\n", v)
		//time.Sleep(500 * time.Millisecond)
	}
}

func channelTest() {
	ch := make(chan int)
	// 启动发送goroutine
	go sendOnlyTen(ch)
	// 启动接收goroutine
	go receiveOnlyTen(ch)
	// 使用select进行多路复用
	timeout := time.After(10 * time.Second)
	for {
		select {
		case v, ok := <-ch:
			if !ok {
				fmt.Println("Channel已关闭")
				return
			}
			fmt.Printf("主goroutine接收到: %d\n", v)
			time.Sleep(500 * time.Millisecond)
		case <-timeout:
			fmt.Println("操作超时")
			return
		default:
			fmt.Println("没有数据，等待中...")
			time.Sleep(500 * time.Millisecond)
		}
	}
}
func channelValueTest() {
	// 创建缓冲通道，缓冲区大小为10
	ch := make(chan int, 10)

	wg := common.Wg
	wg.Add(2) // 等待生产者和消费者两个goroutine

	// 生产者goroutine：发送100个整数到通道
	go func() {
		defer wg.Done()
		for i := 1; i <= 100; i++ {
			ch <- i
			fmt.Printf("发送: %d\n", i)
		}
		close(ch) // 发送完成后关闭通道
	}()

	// 消费者goroutine：从通道接收并打印
	go func() {
		defer wg.Done()
		for num := range ch {
			fmt.Printf("接收: %d\n", num)
		}
	}()
	wg.Wait()
	fmt.Println("程序结束")
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
