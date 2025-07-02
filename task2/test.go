package task2

import "fmt"

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
