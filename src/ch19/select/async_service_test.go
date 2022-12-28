package csp_test

import (
	"fmt"
	"testing"
	"time"
)

func service() string {
	time.Sleep(time.Millisecond * 50)
	// time.Sleep(time.Millisecond * 500)
	return "service - Done"
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

func TestSelect(t *testing.T) {
	select {
	case ret := <-AsyncServiceBufferModel():
		t.Log(ret)
	case <-time.After(time.Millisecond * 100):
		t.Error("Time out")
	}
}
