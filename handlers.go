package main

import (
	"fmt"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	message := GetMessage("home")
	fmt.Fprintf(w, "Главная страница сервиса \n%s", message)
}

func AboutHandler(w http.ResponseWriter, r *http.Request) {
	message := GetMessage("about")
	fmt.Fprintf(w, "О приложении\n%s", message)
}
