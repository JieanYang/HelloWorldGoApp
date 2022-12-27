package error_test

import (
	"errors"
	"fmt"
	"testing"
)

var LessThantTwoError = errors.New("n should be not less than 2")
var LargerThantHundredError = errors.New("n should be not larger than 100")

func GetFibonacci(n int) ([]int, error) {
	if n < 2 {
		return nil, LessThantTwoError
	}
	if n > 100 {
		return nil, LargerThantHundredError
	}

	fibList := []int{1, 1}

	for i := 2; /*短变量声明 := */ i < n; i++ {
		fibList = append(fibList, fibList[i-2]+fibList[i-1])
	}
	return fibList, nil
}

func TestGetFibnacci(t *testing.T) {
	if value, err := GetFibonacci(1); err != nil {
		if err == LessThantTwoError {
			fmt.Println("It is less.")
		}
		t.Error(err)
	} else {
		t.Log(value)
	}
	if value, err := GetFibonacci(10); err != nil {
		t.Error(err)
	} else {
		t.Log(value)
	}
}
