package groutine_test

import (
	"fmt"
	"testing"
)

func TestGroutine(t *testing.T) {
	for i := 0; i < 10; i++ {
		// 匿名函数运行
		// func(i int) {
		// 	fmt.Println(i)
		// }(i)
		// Groutine 协程运行
		go func(i int) {
			fmt.Println(i)
		}(i) // 由于 Go 的值传递，所以在每个协程中拥有的 i 的地址是不同的，所以可以正确运行
		// time.Sleep(time.Millisecond * 50)
		// Groutine 常见错误
		// go func() {
		// 	fmt.Println(i) // 每次都会打印 10，因为 i 变量 在 test 所在的协程以及其他的其他协程共享，存在竞争条件，我们需要锁的机制才能修复
		// }()
	}
}
