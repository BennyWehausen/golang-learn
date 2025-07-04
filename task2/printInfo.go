package task2

import "fmt"

// Person 结构体包含姓名和年龄
type Person struct {
	Name string
	Age  int
}

// Employee 结构体组合了Person并添加员工ID
type Employee struct {
	Person     // 嵌入Person结构体（组合）
	EmployeeID string
}

func (e Employee) PrintInfo() {
	fmt.Printf("员工信息:\n")
	fmt.Printf("姓名: %s\n", e.Name)       // 可以直接访问嵌入结构的字段
	fmt.Printf("年龄: %d\n", e.Age)        // 同上
	fmt.Printf("工号: %s\n", e.EmployeeID) // 访问自己的字段
}
