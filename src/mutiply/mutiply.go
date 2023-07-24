package src

import (

	"fmt"
	"sync"
)

func FuncGoroutine() {
    a := 1
    // 创建四个无缓冲的 channel，用于协程之间的同步
    ch1, ch2, ch3, ch4:= make(chan struct{}), make(chan struct{}), make(chan struct{}),make(chan struct{})
	// chs:=make([]chan struct{},10)
	
    var wg sync.WaitGroup
	

    wg.Add(4) // WaitGroup 需要等待四个协程执行结束

    // 启动 goroutine1，使用 ch1 和 ch2 进行同步
    go goroutine1(&a, ch1, ch2,&wg)
    // 启动 goroutine2，使用 ch2 和 ch3 进行同步
    go goroutine2(&a, ch2, ch3,&wg)
    // 启动 goroutine3，使用 ch3 进行同步
    go goroutine3(&a, ch3,ch4,&wg)
	go goroutine4(&a,ch4,&wg)

    ch1 <- struct{}{} // 向 ch1 发送一个空结构体，启动 goroutine1 的执行

    wg.Wait() // 等待所有协程执行结束

    fmt.Println(a)
}

func goroutine1(i *int, ch1, ch2 chan struct{},wg *sync.WaitGroup) {
    defer func ()  {
		// close(ch1)
		wg.Done()	
	}() // 在 goroutine1 结束时关闭 ch1
    <-ch1            // 等待 ch1 上接收到数据
    *i++             // 修改变量 i 的值
    fmt.Println("goroutine1 ", *i)
    ch2 <- struct{}{} // 向 ch2 发送一个空结构体，启动 goroutine2 的执行
}

func goroutine2(i *int, ch2, ch3 chan struct{},wg *sync.WaitGroup) {
	defer func ()  {
		// close(ch2)
		wg.Done()	
	}() // 在 goroutine2 结束时关闭 ch2
    // defer close(ch2) // 在 goroutine2 结束时关闭 ch2
    <-ch2            // 等待 ch2 上接收到数据
    *i *= *i         // 修改变量 i 的值
    fmt.Println("goroutine2 ", *i)
    ch3 <- struct{}{} // 向 ch3 发送一个空结构体，启动 goroutine3 的执行
}

func goroutine3(i *int, ch3,ch4 chan struct{},wg *sync.WaitGroup) {
	defer func ()  {
		// close(ch3)
		wg.Done()	
	}() // 在 goroutine3 结束时关闭 ch3
    // defer close(ch3) // 在 goroutine3 结束时关闭 ch3
    <-ch3            // 等待 ch3 上接收到数据
    *i--             // 修改变量 i 的值
    fmt.Println("goroutine3 ", *i)
	ch4<-struct{}{}
}
func goroutine4(i *int, ch4 chan struct{},wg *sync.WaitGroup) {
	defer func ()  {
		// close(ch4)
		wg.Done()	
	}() // 在 goroutine3 结束时关闭 ch3
    // defer close(ch3) // 在 goroutine3 结束时关闭 ch3
    <-ch4           // 等待 ch4 上接收到数据
    *i+=10             // 修改变量 i 的值
    fmt.Println("goroutine4 ", *i)
}