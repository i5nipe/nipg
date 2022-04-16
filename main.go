package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/", HelloHandler)

	fmt.Println("Server started at port 1313")
	http.ListenAndServe(":1313", nil)
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintln(w, "Hello Word")
}
