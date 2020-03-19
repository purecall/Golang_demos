package object_pool

import (
	"fmt"
	"testing"
	"time"
)

func TestObjPool(t *testing.T) {
	pool := NewObjPool(10)
	if err := pool.ReleaseObj(&ReusableObj{}); err != nil {
		//尝试放置对象，超出池大小
		t.Error(err) // overflow
	}
	for i := 0; i < 11; i++ {
		if v, err := pool.GetObj(time.Second * 1); err != nil {
			t.Error(err)
		} else {
			fmt.Printf("%T\n", v)
			if err := pool.ReleaseObj(v); err != nil {
				// 10个对象，get了11次，如果不release，就会阻塞，导致error:time out
				t.Error(err)
			}
		}
	}
}
