package main

import (
	"fmt"
	"net/http"
)

func test01(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintln(w, "<h1>Hello Fuck<h1>")
}

func main() {
	http.HandleFunc("/hello", test01)

	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		fmt.Printf("http serve failed ,err:%v\n", err)
		return
	}
}
