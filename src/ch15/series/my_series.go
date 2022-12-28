package series

import "fmt"

// 统一源码中可以含有多个 init 函数
func init() {
	fmt.Println("init 1")
}

func init() {
	fmt.Println("init 2")
}

// 小写声明是不可以被包外访问
func square(n int) int {
	return n * n
}

func GetFibonacci(n int) []int {
	fibList := []int{1, 1}

	for i := 2; /*短变量声明 := */ i < n; i++ {
		fibList = append(fibList, fibList[i-2]+fibList[i-1])
	}
	return fibList
}
