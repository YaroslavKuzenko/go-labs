package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

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

	query := r.URL.Query()
	senderFilter := query.Get("sender")
	weightFilter := query.Get("weight")

	filteredParcels := filterParcels(parcels, senderFilter, weightFilter)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(filteredParcels)
}

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

	if err := SaveParcel(parcel); err != nil {
		http.Error(w, "Помилка при збереженні посилки", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(parcel)
}

func ParcelByIDHandler(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/parcel/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Невірний ID", http.StatusBadRequest)
		return
	}

	switch r.Method {
	case http.MethodGet:
		parcels, err := LoadParcels()
		if err != nil {
			http.Error(w, "Помилка при завантаженні посилок", http.StatusInternalServerError)
			return
		}

		var foundParcel *Parcel
		for _, p := range parcels {
			if p.ID == id {
				foundParcel = &p
				break
			}
		}

		if foundParcel == nil {
			http.Error(w, "Посилка не знайдена", http.StatusNotFound)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(foundParcel)

	case http.MethodPut:
		var parcel Parcel
		if err := json.NewDecoder(r.Body).Decode(&parcel); err != nil {
			http.Error(w, "Невірні дані", http.StatusBadRequest)
			return
		}
		parcel.ID = id

		if err := UpdateParcel(parcel); err != nil {
			http.Error(w, "Помилка при оновленні посилки", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(parcel)

	case http.MethodDelete:
		if err := DeleteParcel(id); err != nil {
			http.Error(w, "Помилка при видаленні посилки", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusNoContent)

	default:
		http.Error(w, "Метод не дозволений", http.StatusMethodNotAllowed)
	}
}

func filterParcels(parcels []Parcel, senderFilter, weightFilter string) []Parcel {
	var filtered []Parcel
	for _, parcel := range parcels {
		if senderFilter != "" && !strings.Contains(parcel.Sender, senderFilter) {
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
