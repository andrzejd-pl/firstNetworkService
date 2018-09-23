package main

import (
	"fmt"
	"net/http"
	"strconv"
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
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
