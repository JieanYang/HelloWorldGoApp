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
	// runtime.GC() // GC 会清除 sync.pool 中缓存的对象, 这里的描述与测试不一样，有可能是因为 Go 后期版本对特性做了修改
	v1, _ := pool.Get().(int)
	fmt.Println(v1)
	v2, _ := pool.Get().(int) // Get 拿出后，池里就没数据了，所以需要重新产生
	fmt.Println(v2)
}
