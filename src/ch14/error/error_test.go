package error_test

import "testing"

func GetFibonacci(n int) []int {
	fibList := []int{1, 1}

	for i := 2; /*短变量声明 := */ i < n; i++ {
		fibList = append(fibList, fibList[i-2]+fibList[i-1])
	}
	return fibList
}

func TestGetFibnacci(t *testing.T) {
	t.Log(GetFibonacci(10))
}
