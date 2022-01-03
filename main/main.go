package main

import "fmt"

/*
File name    : main.go
Author       : miaoyc
Create date  : 2022/1/3 11:46 下午
Description  : main函数相关常见问题
*/

// 问题1，main函数提前退出，后台goroutine无法保证完成任务

func main()  {
	go fmt.Println("go hello")
	fmt.Println("hello")
}