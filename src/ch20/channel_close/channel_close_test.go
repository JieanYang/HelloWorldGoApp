package csp_test

import (
	"fmt"
	"sync"
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

func dataProducer(ch chan int, wg *sync.WaitGroup) {
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		wg.Done()
	}()
}

func dataReceiver(ch chan int, wg *sync.WaitGroup) {
	go func() {
		for i := 0; i < 10; i++ {
			data := <-ch
			fmt.Println(data)
		}
		wg.Done()
	}()
}

func TestCloseChannel(t *testing.T) {
	var wg sync.WaitGroup
	ch := make(chan int)
	wg.Add(1)
	dataProducer(ch, &wg)
	wg.Add(1)
	dataReceiver(ch, &wg)
	wg.Wait()
}
