package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Инициализация запуска сервера")
	http.HandleFunc("/", HomeHandler)
	http.HandleFunc("/about", AboutHandler)
	http.HandleFunc("/api/random", RandomHandler)

	fmt.Println("Сервер запущен на http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Ошибка запуска сервера:", err)
	}
}
