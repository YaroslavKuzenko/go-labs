package main

import (
	"log"
	"net/http"
	"os"
	"time"
)

// Middleware для логування
func withLogging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		file, err := os.OpenFile("log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Println("Помилка при відкритті файлу логів:", err)
			next.ServeHTTP(w, r)
			return
		}
		defer file.Close()
		logger := log.New(file, "", log.LstdFlags)
		logger.Printf("Method: %s, Path: %s, Time: %s\n", r.Method, r.URL.Path, time.Now().Format(time.RFC3339))
		next.ServeHTTP(w, r)
	})
}

// Middleware для авторизації
func withAuthorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		apiKey := r.Header.Get("X-API-KEY")
		if apiKey != "secret123" {
			http.Error(w, "Невірний API ключ", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}
