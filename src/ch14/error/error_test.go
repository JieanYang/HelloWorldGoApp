package error_test

import (
	"errors"
	"testing"
)

func GetFibonacci(n int) ([]int, error) {
	if n < 2 || n > 100 {
		return nil, errors.New("n should be in [2, 100]")
	}

	fibList := []int{1, 1}

	for i := 2; /*短变量声明 := */ i < n; i++ {
		fibList = append(fibList, fibList[i-2]+fibList[i-1])
	}
	return fibList, nil
}

func TestGetFibnacci(t *testing.T) {
	if value, err := GetFibonacci(10); err != nil {
		t.Error(err)
	} else {
		t.Log(value)
	}
	if value, err := GetFibonacci(-10); err != nil {
		t.Error(err)
	} else {
		t.Log(value)
	}
}
