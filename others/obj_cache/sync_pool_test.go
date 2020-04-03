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

	v := pool.Get().(int) // 为空，先New
	//Get出来以后，里面就没有这个对象了
	fmt.Println(v)
	pool.Put(3)
	//runtime.GC() //GC 会清除sync.pool中缓存的对象
	//不调用GC的话，会输出100 3，符合预期
	//调用GC的话，会触发2次New，输出2个100
	v1, _ := pool.Get().(int)
	fmt.Println(v1)
}

func TestSyncPoolInMultiGoroutine(t *testing.T) {
	pool := &sync.Pool{
		New: func() interface{} {
			fmt.Println("Create a new object")
			return 10
		},
	}

	pool.Put(100)
	pool.Put(100)
	pool.Put(100)

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			fmt.Println(pool.Get())
			wg.Done()
		}(i)
		wg.Wait()
	}
}
