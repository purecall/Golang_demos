package concat_string_test

import (
	"bytes"
	"testing"
)

func BenchmarkConcatStringByAdd(b *testing.B) {
	elems := []string{"1", "2", "3", "4", "5"}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ret := ""
		for _, elem := range elems {
			ret += elem
		}
	}
	b.StopTimer()
}

func BenchmarkConcatStringByBytesBuffer(b *testing.B) {
	elems := []string{"1", "2", "3", "4", "5"}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var buf bytes.Buffer

		for _, elem := range elems {
			buf.WriteString(elem)
		}
	}
	b.StopTimer()
}

/*
PS C:\Users\vct\Desktop\github\Golang_demos\concat_string_benchmark_test> go test -bench="." -benchmem
goos: windows
goarch: amd64
BenchmarkConcatStringByAdd-8             7750515               150 ns/op              16 B/op          4 allocs/op
BenchmarkConcatStringByBytesBuffer-8    16470370                81.2 ns/op            64 B/op          1 allocs/op
PASS
ok      _/C_/Users/vct/Desktop/github/Golang_demos/concat_string_benchmark_test 2.957s
*/
