package goroutine

import (
	"fmt"
	"time"
)

/*
File name    : goroutine.go
Author       : miaoyc
Create Date  : 2023/9/12 14:05
Update Date  : 2023/9/12 14:05
Description  : goroutine常见错误示例，参考文章：https://juejin.cn/post/7121754891533418504
*/

// example1Failed, 示例1，忘记向go routine传参, 错误示例
func example1Failed() {
	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println(i)
		}()
	}

	time.Sleep(2 * time.Second)
}

// example1Success, 示例1，忘记向go routine传参，正确示例
func example1Success() {
	for i := 0; i < 10; i++ {
		go func(j int) {
			fmt.Println(j)
		}(i)
	}

	time.Sleep(2 * time.Second)
}
