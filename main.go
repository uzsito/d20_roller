package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	rand.Seed(time.Now().UnixNano())
	roll := rand.Intn(20) + 1

	color := "black"
	if roll == 20 {
		color = "green"
	} else if roll == 1 {
		color = "red"
	}

	fmt.Fprintf(w, `
		<html>
			<head>
				<title>D20 dobás</title>
				<style>
					body { font-family: Arial; text-align: center; margin-top: 50px; background-color: #f0f8ff; }
					.circle {
						width: 100px;
						height: 100px;
						border-radius: 50%%;
						border: 3px solid %s;
						display: flex;
						align-items: center;
						justify-content: center;
						font-size: 36px;
						margin: 0 auto 20px auto;
					}
				</style>
			</head>
			<body>
				<h1 style="color: #1e90ff;">D20 Kockadobó</h1>
				<div class="circle">%d</div>
				<p>Frissítsd az oldalt egy új dobáshoz!</p>
				<p>A kritikus dobások színesek.</p>
			</body>
		</html>
	`, color, roll)
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Fut: http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}