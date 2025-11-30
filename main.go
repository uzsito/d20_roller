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
			<head><title>D20 dobás</title></head>
			<body style="font-family: Arial; text-align: center; margin-top: 50px;">
				<h1>D20 Kockadobó</h1>
				<h2 style="color:%s;">Eredmény: %d</h2>
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