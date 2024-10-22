package main

import (
	"log"
	"net/http"
)

func main() {
	// Ініціалізація сховища
	if err := InitializeStorage(); err != nil {
		log.Fatalf("Помилка ініціалізації сховища: %v", err)
	}

	http.HandleFunc("/parcels", GetParcelsHandler)
	http.HandleFunc("/parcel", CreateParcelHandler)
	http.HandleFunc("/parcel/", ParcelByIDHandler)

	log.Println("Сервер запущений на http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
