package main

import (
	"fmt"
	"net/http"
	"strconv"
	"sync"
)

func silniaInt(n int64) (result int64) {
	result = 1
	if (n < 1) {
		return
	} else {
		result = n * silniaInt(n-1)
		return
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	var n, _ = strconv.ParseInt(r.FormValue("n"), 10, 0)
	fmt.Fprintln(w, silniaInt(n))
}

func main() {
	var wg sync.WaitGroup

	wg.Add(20)
	for i := int64(1); i <= 20; i++ {
		go func(n int64) {
			wg.Done()
			fmt.Println("Silnia ", n, ": ", silniaInt(n))
		}(i)
	}

	wg.Wait()
	//http.HandleFunc("/", handler)
	//http.ListenAndServe(":8080", nil)
}
