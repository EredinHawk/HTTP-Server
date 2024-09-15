package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	server := constructServer()

	fmt.Println("Сервер localhost:8090 запущен и прослушивает входящие запросы...")
	err := server.ListenAndServe()
	if err != nil {
		log.Fatalf("ошибка при запуске сервера (%v)", err)
	}
}

// constructServer возвращает инициализированный сервер типа *http.Server
func constructServer() *http.Server {
	router := http.NewServeMux()
	router.HandleFunc("GET /{year}/{month}/{day}/", getAgeHandler)

	srv := &http.Server{
		Addr:         "localhost:8090",
		Handler:      router, // HTTP мультиплексор, или по другому роутер
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	return srv
}

// getAgeHandler обработчик принимает параметром дату и возвращает время, прошедшее с её начала
func getAgeHandler(w http.ResponseWriter, r *http.Request) {
	birthday, err := ScanDate(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "возраст  - %s", AgeString(birthday))
}
