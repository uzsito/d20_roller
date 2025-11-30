package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	rand.Seed(time.Now().UnixNano())

	dice := 1
	if r.Method == http.MethodPost {
		r.ParseForm()
		if d := r.FormValue("dice"); d != "" {
			if n, err := strconv.Atoi(d); err == nil && (n == 1 || n == 2) {
				dice = n
			}
		}
	}

	rolls := make([]int, dice)
	colors := make([]string, dice)
	total := 0
	for i := 0; i < dice; i++ {
		rolls[i] = rand.Intn(20) + 1
		switch rolls[i] {
		case 20:
			colors[i] = "green"
		case 1:
			colors[i] = "red"
		default:
			colors[i] = "black"
		}
		total += rolls[i]
	}

	fmt.Fprintf(w, `
		<html>
			<head>
				<title>D20 dobás</title>
				<style>
					body { font-family: Arial; text-align: center; margin-top: 50px; background-color: #f0f8ff; }
					.dice-box {
						width: 100px;
						height: 100px;
						display: flex;
						align-items: center;
						justify-content: center;
						font-size: 36px;
						margin: 10px;
					}
					form { margin-bottom: 20px; }
				</style>
			</head>
			<body>
				<h1 style="color: #1e90ff;">D20 Kockadobó</h1>
				<form method="post">
					<label><input type="radio" name="dice" value="1" %s> 1 kocka</label>
					<label><input type="radio" name="dice" value="2" %s> 2 kocka</label>
					<button type="submit">Dobás</button>
				</form>
	`, checkedAttr(dice, 1), checkedAttr(dice, 2))

	fmt.Fprintf(w, `<div style="display: flex; justify-content: center; flex-wrap: wrap;">`)
	for i := 0; i < dice; i++ {
		fmt.Fprintf(w, `<div class="dice-box" style="border: 3px solid %s;">%d</div>`, colors[i], rolls[i])
	}
	fmt.Fprintf(w, `</div>`)

	fmt.Fprintf(w, `<p>Összeg: %d</p></body></html>`, total)
}

func checkedAttr(dice int, value int) string {
	if dice == value {
		return "checked"
	}
	return ""
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Fut: http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}