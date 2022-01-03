package main

/*
File name    : main.go
Author       : miaoyc
Create date  : 2022/1/3 11:40 下午
Description  : recover相关常见问题
*/

func main() {
	defer func() {
		recover()
	}()
	panic(1)
}
