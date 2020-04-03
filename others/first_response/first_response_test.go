package first_response_test

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

func runTask(id int) string {
	time.Sleep(10 * time.Millisecond)
	return fmt.Sprintf("The result is from %d", id)
}

func FirstResponse() string {
	numOfRunner := 10
	ch := make(chan string, numOfRunner)
	for i := 0; i < numOfRunner; i++ {
		go func(i int) {
			ret := runTask(i)
			ch <- ret
		}(i)
	}
	return <-ch
}

func TestFirstResponse(t *testing.T) {
	t.Log("Before:", runtime.NumGoroutine())
	t.Log(FirstResponse())
	time.Sleep(time.Second * 1)
	t.Log("After:", runtime.NumGoroutine())
}

/*
如果第17行是 ch:= make(chan string)，
可以看到，结束后有11条协程，如果是一个服务器程序，每次调用都会有协程被阻塞。如果有很多协程被阻塞，系统资源就会被耗尽

PS C:\Users\vct\Desktop\go_src> go test -v .\first_package_test.go
=== RUN   TestFirstResponse
--- PASS: TestFirstResponse (1.01s)
    first_package_test.go:28: Before: 2
    first_package_test.go:29: The result is from 0
    first_package_test.go:31: After: 11
PASS
ok      command-line-arguments  1.195s
*/

/*
使用现在的buffered channel，使其解耦合
PS C:\Users\vct\Desktop\go_src> go test -v .\first_package_test.go
=== RUN   TestFirstResponse
--- PASS: TestFirstResponse (1.01s)
    first_package_test.go:28: Before: 2
    first_package_test.go:29: The result is from 0
    first_package_test.go:31: After: 2
PASS
*/
