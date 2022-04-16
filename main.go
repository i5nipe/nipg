package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/", HelloHandler)

	http.HandlerFunc("/dado", RollDice)

	fmt.Println("Server started at port 1313")
	http.ListenAndServe(":1313", nil)
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello Word")
}

func RollDice(w http.ResponseWriter, r *http.Request) {

	keys, ok := r.URL.Query()["faces"]

	faces := 20

	if ok {
		faces = keys[0]
	}

	fmt.Fprintln(w, faces)

}
