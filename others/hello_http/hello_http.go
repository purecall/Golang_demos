package main

import (
	"fmt"
	"net/http"
	"time"
)



func main() {
	// 定义不同url pattern下的不同处理逻辑
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World!")
	})
	http.HandleFunc("/time", func(w http.ResponseWriter, r *http.Request) {
		t := time.Now()
		timeStr := fmt.Sprintf("{\"time\": \"%s\"}", t)
		w.Write([]byte(timeStr))
	})

	http.ListenAndServe(":8080", nil)
}
