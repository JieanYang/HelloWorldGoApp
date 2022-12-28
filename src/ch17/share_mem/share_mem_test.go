package sharemem_test

import (
	"sync"
	"testing"
	"time"
)

func TestCounterWithUnsafeThread(t *testing.T) {
	counter := 0
	for i := 0; i < 5000; i++ {
		go func() {
			counter++
		}()
	}
	time.Sleep(1 * time.Second)
	t.Logf("counter = %d", counter)
}

func TestCounterWithSafeThread(t *testing.T) {
	var mut sync.Mutex

	counter := 0
	for i := 0; i < 5000; i++ {

		go func() {
			defer func() {
				mut.Unlock()
			}()
			mut.Lock()
			counter++
		}()

	}
	time.Sleep(1 * time.Second) // test 文件所在的协程运行速度太快，以至于生成的所有协程没法正确执行完，所以需要在test文件加入此行保证结果正确
	t.Logf("counter = %d", counter)
}

func TestCounterWithWaitGroup(t *testing.T) {
	var mut sync.Mutex
	var wg sync.WaitGroup

	counter := 0
	for i := 0; i < 5000; i++ {
		wg.Add(1)

		go func() {
			defer func() {
				mut.Unlock()
			}()
			mut.Lock()
			counter++
			wg.Done()
		}()

	}
	wg.Wait()
	t.Logf("counter = %d", counter)
}
