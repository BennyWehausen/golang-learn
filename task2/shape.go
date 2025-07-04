package task2

import "math"

/**
面向对象
题目 ：定义一个 Shape 接口，包含 Area() 和 Perimeter() 两个方法。然后创建 Rectangle 和 Circle 结构体，实现 Shape 接口。
在主函数中，创建这两个结构体的实例，并调用它们的 Area() 和 Perimeter() 方法。
*/
// 定义 Shape 接口
type Shape interface {
	Area() float64
	Perimeter() float64
}

// Rectangle 结构体
type Rectangle struct {
	Width, Height float64
}

// 实现 Area 方法
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// 实现 Perimeter 方法
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// Circle 结构体
type Circle struct {
	Radius float64
}

// 实现 Area 方法
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// 实现 Perimeter 方法
func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}
