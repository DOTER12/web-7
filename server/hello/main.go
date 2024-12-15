package main

// некоторые импорты нужны для проверки
import (
	"fmt"
	"net/http"
)

// Обработчик HTTP-запросов
func handlerGET(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write([]byte("Hello, web!"))
}

func main() {
	// Регистрируем обработчик для пути "/"
	http.HandleFunc("/", handlerGET)

	// Запускаем веб-сервер на порту 8080
	err := http.ListenAndServe(":8082", nil)
	if err != nil {
		fmt.Println("Ошибка запуска сервера:", err)
	}
}
