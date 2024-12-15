package main

// некоторые импорты нужны для проверки
import (
	"fmt"
	"net/http" // пакет для поддержки HTTP протокола
)

// Обработчик HTTP-запросов

func main() {
	http.HandleFunc("/api/user", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		name := r.URL.Query().Get("name")
		fmt.Fprintf(w, "Hello,%s!", name)

	})
	http.ListenAndServe(":8083", nil)

	err := http.ListenAndServe(":8083", nil)
	if err != nil {
		fmt.Println("Ошибка запуска сервера:", err)
	}
}
