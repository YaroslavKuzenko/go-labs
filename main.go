package main

import (
	"log"
	"net/http"
	"strings"
)

func main() {
	// Ініціалізація сховища
	if err := InitializeStorage(); err != nil {
		log.Fatalf("Помилка ініціалізації сховища: %v", err)
	}

	http.Handle("/", withAuthorization(withLogging(http.HandlerFunc(router))))

	log.Println("Сервер запущений на http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func router(w http.ResponseWriter, r *http.Request) {
	switch {
	case r.Method == http.MethodGet && r.URL.Path == "/parcels":
		GetParcelsHandler(w, r)
	case r.Method == http.MethodPost && r.URL.Path == "/parcel":
		CreateParcelHandler(w, r)
	case strings.HasPrefix(r.URL.Path, "/parcel/"):
		ParcelByIDHandler(w, r)
	default:
		http.NotFound(w, r)
	}
}
