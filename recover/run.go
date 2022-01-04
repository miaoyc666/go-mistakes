package main

import (
	"fmt"
	"sync"
)

/*
File name    : run.go
Author       : miaoyc
Create date  : 2022/1/4 11:42 下午
Description  : 推荐的并发处理模式
*/

var wg sync.WaitGroup

func div(num int) {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err) }
		wg.Done()
	}()
	fmt.Printf("10/%d=%d\n", num, 10/num)
}

func main() {
	for i:=0; i<10; i++ {
		wg.Add(i)
		go div(i)
	}
	wg.Wait()
}