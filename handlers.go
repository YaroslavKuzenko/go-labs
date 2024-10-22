package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// GetParcelsHandler обробляє запити на отримання всіх посилок з можливістю фільтрації
func GetParcelsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Метод не дозволений", http.StatusMethodNotAllowed)
		return
	}
	parcels, err := LoadParcels()
	if err != nil {
		http.Error(w, "Помилка при завантаженні посилок", http.StatusInternalServerError)
		return
	}

	// Фільтрація за query parameters
	query := r.URL.Query()
	nameFilter := query.Get("sender")
	weightFilter := query.Get("weight")

	filteredParcels := filterParcels(parcels, nameFilter, weightFilter)

	// Повертаємо відфільтровані посилки
	json.NewEncoder(w).Encode(filteredParcels)
}

// створення нової посилки
func CreateParcelHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Метод не дозволений", http.StatusMethodNotAllowed)
		return
	}
	var parcel Parcel
	if err := json.NewDecoder(r.Body).Decode(&parcel); err != nil {
		http.Error(w, "Невірні дані", http.StatusBadRequest)
		return
	}
	parcels, err := LoadParcels()
	if err != nil {
		http.Error(w, "Помилка при завантаженні посилок", http.StatusInternalServerError)
		return
	}
	parcel.ID = fmt.Sprintf("%d", len(parcels)+1)
	parcels = append(parcels, parcel)
	if err := SaveParcels(parcels); err != nil {
		http.Error(w, "Помилка при збереженні посилок", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(parcel)
}

// оновлення або видалення посилки за ID
func ParcelByIDHandler(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/parcel/")
	parcels, err := LoadParcels()
	if err != nil {
		http.Error(w, "Помилка при завантаженні посилок", http.StatusInternalServerError)
		return
	}

	var parcel *Parcel
	for i := range parcels {
		if parcels[i].ID == id {
			parcel = &parcels[i]
			break
		}
	}
	if parcel == nil {
		http.Error(w, "Посилка не знайдена", http.StatusNotFound)
		return
	}

	switch r.Method {
	case http.MethodGet:
		json.NewEncoder(w).Encode(parcel)
	case http.MethodPut:
		if err := json.NewDecoder(r.Body).Decode(parcel); err != nil {
			http.Error(w, "Невірні дані", http.StatusBadRequest)
			return
		}
		if err := SaveParcels(parcels); err != nil {
			http.Error(w, "Помилка при збереженні посилок", http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(parcel)
	case http.MethodDelete:
		for i := range parcels {
			if parcels[i].ID == id {
				parcels = append(parcels[:i], parcels[i+1:]...)
				break
			}
		}
		if err := SaveParcels(parcels); err != nil {
			http.Error(w, "Помилка при збереженні посилок", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusNoContent)
	default:
		http.Error(w, "Метод не дозволений", http.StatusMethodNotAllowed)
	}
}

func filterParcels(parcels []Parcel, nameFilter, weightFilter string) []Parcel {
	var filtered []Parcel
	for _, parcel := range parcels {
		if nameFilter != "" && !strings.Contains(parcel.Sender, nameFilter) {
			continue
		}
		if weightFilter != "" {
			weight := parcel.Weight
			weightMatch := weightFilter == "light" && weight <= 1.0 || weightFilter == "heavy" && weight > 1.0
			if !weightMatch {
				continue
			}
		}
		filtered = append(filtered, parcel)
	}
	return filtered
}
