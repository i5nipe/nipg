package main

import (
	"fmt"
	"html/template"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var tmpl = template.Must(template.ParseGlob("form/*"))

func Index(w http.ResponseWriter, r *http.Request) {
	db := dbConn()

	selDB, err := db.Query("SELECT * FROM users ORDER BY id DESC")
	if err != nil {
		fmt.Println("Falha ao conectar no banco de dados INDEX")
		panic(err.Error())
	}
	emp := Users{}
	res := []Users{}

	for selDB.Next() {
		var id int
		var user, password, role string

		err = selDB.Scan(&id, &user, &password, &role)
		if err != nil {
			panic(err.Error())
		}
		emp.Id = id
		emp.User = user
		emp.Password = password
		emp.Role = role
		res = append(res, emp)
	}

	fmt.Println(res)
	tmpl.ExecuteTemplate(w, "Index", res)
	defer db.Close()
}

func main() {

	http.HandleFunc("/", Index)
	http.HandleFunc("/api/dado", RollDicehttp)

	fmt.Println("Server started at port 1313")
	http.ListenAndServe(":1313", nil)
}

func RollDicehttp(w http.ResponseWriter, r *http.Request) {
	var faces int

	keys, ok := r.URL.Query()["d"]

	if ok {
		faces_str := keys[0]
		faces, err := strconv.Atoi(faces_str)
		if err != nil {
			fmt.Println("'?d' não é um Inteiro")
			faces = 20
		}

		if faces >= 10000 {
			fmt.Println("9999 é a quantidade maxima de lados")
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
