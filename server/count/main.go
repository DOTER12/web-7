package main

// некоторые импорты нужны для проверки
import (
	"fmt"
	"net/http"
	"strconv" // вдруг понадобиться вам ;)
)

func main() {
	var counter = 0
	http.HandleFunc("/count", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		if r.Method == "GET" {
			w.Write([]byte(strconv.Itoa(counter)))
		}
		if r.Method == "POST" {
			r.ParseForm()
			numberString := r.Form.Get("count")
			number, err := strconv.Atoi(numberString)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)

				w.Write([]byte("это не число"))
				return
			}
			counter += number
		}
	})
	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		fmt.Println("Ошибка запуска сервера:", err)
	}

}
