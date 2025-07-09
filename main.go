package main

import (
	"github.com/gin-gonic/gin"
	"golang-learn/router"
)

func main() {
	//fmt.Printf("TAG=", foo.TAG)
	//test.InterfaceTest()
	///*Task1 work */
	//test.TestTask1()
	/////*Task2 work */
	/////*Goroutine work */
	//test.TestTask2()
	/////*面向对象 work*/
	//test.TestShape()
	//test.PrintInfoTest()
	/////*Channel work*/
	//test.ChannelTest()
	//test.ChannelValueTest()
	/////*锁机制 work */
	//test.SafeCounterTest()
	//task2.SafeCounterAtomic()
	////test.SqlxTest()
	//test.ModelSetTest()
	//test.GormQueryTest(5)
	//test.HookFuncTest()
	r := gin.Default()
	//路由
	router.Router(r)
	// 启动服务
	if err := r.Run(":8080"); err != nil {
		panic(err)
	}
}
