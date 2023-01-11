package sync_pool_test

import (
	"fmt"
	"sync"
	"testing"
)

func TestSyncPool(t *testing.T) {
	pool := &sync.Pool{
		New: func() interface{} {

			fmt.Println("Create a new object.")
			return 100
		},
	}

	v := pool.Get().(int) // 断言类型
	fmt.Println(v)
	pool.Put(3)
	// runtime.GC() // GC 会清除 sync.pool 中缓存的对象
	v1, _ := pool.Get().(int)
	fmt.Println(v1)
}