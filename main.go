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
	fmt.Fprintln(w, "Pagina principal")
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

		if faces >= 10000 {
			fmt.Println("10000 é a quantidade maxima de lados")
			return
		}
		result := rollDice(faces)
		fmt.Fprintln(w, result)
	} else {
		faces = 20
		result := rollDice(faces)
		fmt.Fprintln(w, result)
	}
}

func rollDice(faces int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(faces) + 1
}
