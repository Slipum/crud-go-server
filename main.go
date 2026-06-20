package main

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"os"
)

type User struct {
	ID    int16  `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

var globalUserID int16 = 0

func PingHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	fmt.Fprintf(w, "pong")
}

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	user := User{ID: globalUserID, Name: "tempUser", Email: "temp@email.com"}

	if err := json.NewEncoder(w).Encode(user); err != nil {
		slog.Error("Failed to encode user JSON", "error", err)
		return
	}
	globalUserID++
	// Логируем успешный запрос
	slog.Info("Get user success", "method", r.Method, "path", r.URL.Path)
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	fmt.Fprintf(w, "Hello World")
}

func main() {
	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stdout, nil)))

	// Регистрируем обработчики для разных маршрутов
	http.HandleFunc("/hello", HelloHandler)
	http.HandleFunc("/user", GetUserHandler)
	http.HandleFunc("/ping", PingHandler)

	// Запускаем сервер
	fmt.Println("Starting server at port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting the server:", err)
	}
}
