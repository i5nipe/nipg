package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func main() {

	http.HandleFunc("/", HelloHandler)

	http.HandleFunc("/dado", RollDice)

	fmt.Println("Server started at port 1313")
	http.ListenAndServe(":1313", nil)
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello Word")
}

func RollDice(w http.ResponseWriter, r *http.Request) {
	var faces int

	keys, ok := r.URL.Query()["faces"]

	if ok {
		faces_str := keys[0]
		faces, err := strconv.Atoi(faces_str)
		if err != nil {
			fmt.Println("'?faces' não é um Inteiro")
			faces = 20
		}

		fmt.Fprintln(w, faces)
	} else {
		faces = 20
		fmt.Fprintln(w, faces)
	}
}
