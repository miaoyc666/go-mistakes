package main

import (
	"fmt"
	"time"
)

/*
File name    : main.go
Author       : miaoyc
Create date  : 2022/1/3 11:40 下午
Description  : recover相关常见问题，问题整理参考博文：https://blog.csdn.net/wohu1104/article/details/115497582
*/

// 结论：协程A发生panic，协程B无法recover到协程A的panic，只有协程自己内部的recover才能捕获自己抛出的panic

// ARunBPanic 在协程B触发panic之，协程A并没有继续打印，并且主协程的main end也没有打印出来，充分说明了在B协程触发panic之后，在A协程也会因此挂掉，且主协程也会挂掉。
func ARunBPanic() {
	// 协程 A
	go func() {
		for {
			fmt.Println("协程 A")
		}
	}()

	// 协程 B
	go func() {
		time.Sleep(10 * time.Microsecond) 	// 确保协程A先运行起来
		panic("协程 B panic")
	}()

	time.Sleep(1 * time.Second) 			// 充分等待协程B触发panic完成和协程A执行完毕
	fmt.Println("main end")
}

// BPanicMainRecover 开启1个协程B，并在主协程中增加recover机制，尝试在主协程中捕获协程B触发的panic，但是结果未能如愿。 故得出结论，哪个协程发生panic，就需要在哪个协程自身中recover。
func BPanicMainRecover() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("panic err is %s", err)
		}
	}()

	// 协程 B
	go func() {
		panic("协程 B panic")
	}()

	time.Sleep(1 * time.Second) // 充分等待协程 B 触发 panic 完成
	fmt.Println("main end")
}

// BPanicBRecover 开启一个协程B，在协程中panic和尝试recover，可以捕获异常
func BPanicBRecover() {
	go func() {
		defer func() {
			if err := recover(); err != nil {
				fmt.Printf("panic err is %s\n", err)
			}
		}()
		panic("panic")
	}()

	time.Sleep(1 * time.Second) // 充分等待协程触发 panic 完成
	fmt.Println("main end")
}



func main() {
	// ARunBPanic()
	// BPanicMainRecover()
	BPanicBRecover()
}
