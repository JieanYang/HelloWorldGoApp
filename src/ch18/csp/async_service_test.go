package csp_test

import (
	"fmt"
	"testing"
	"time"
)

func service() string {
	time.Sleep(time.Millisecond * 50)
	return "service - Done"
}

func otherTask() {
	fmt.Println("otherTask - working on something else")
	time.Sleep(time.Millisecond * 100)
	fmt.Println("otherTask - Task is done.")
}

func TestService(t *testing.T) {
	fmt.Println(service())
	otherTask()
}

func AsyncServiceClassModel() chan string {
	retCh := make(chan string)
	go func() {
		ret := service()
		fmt.Println("AsyncService - returned result.")
		retCh <- ret
		fmt.Println("AsyncService - service exited.")
	}()

	return retCh
}
func AsyncServiceBufferModel() chan string {
	retCh := make(chan string, 1)
	go func() {
		ret := service()
		fmt.Println("AsyncService - returned result.")
		retCh <- ret
		fmt.Println("AsyncService - service exited.")
	}()

	return retCh
}
func TestAsyncService(t *testing.T) {
	// fmt.Println(service())
	// retCh := AsyncServiceClassModel()
	retCh := AsyncServiceBufferModel()
	otherTask()
	fmt.Println(<-retCh)

}
