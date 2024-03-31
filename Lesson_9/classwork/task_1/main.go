package main

import (
	"net/http"
)

// handler обробляє запити на кореневий маршрут
func handler(w http.ResponseWriter, r *http.Request) {
}

func main() {
	// Зареєструвати обробник для кореневого маршруту
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)

}
