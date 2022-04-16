package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

func main() {

	http.HandleFunc("/", HelloHandler)

	http.HandleFunc("/api/dado", RollDicehttp)

	fmt.Println("Server started at port 1313")
	http.ListenAndServe(":1313", nil)
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello Word")
}

func RollDicehttp(w http.ResponseWriter, r *http.Request) {
	var faces int

	keys, ok := r.URL.Query()["lados"]

	if ok {
		faces_str := keys[0]
		faces, err := strconv.Atoi(faces_str)
		if err != nil {
			fmt.Println("'?lados' não é um Inteiro")
			faces = 20
		}

		fmt.Fprintln(w, faces)
	} else {
		faces = 20
		fmt.Fprintln(w, faces)
	}
}

func RollDice(faces int) {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(faces) + 1
}
