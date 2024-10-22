package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

const fileName = "parcels.json"

// зчитуємо посилки з файлу
func LoadParcels() ([]Parcel, error) {
	file, err := os.Open(fileName)
	if err != nil {
		if os.IsNotExist(err) {
			return []Parcel{}, nil // Повертаємо порожній список, якщо файл не існує
		}
		return nil, err
	}
	defer file.Close()

	var parcels []Parcel
	if err := json.NewDecoder(file).Decode(&parcels); err != nil {
		return nil, err
	}
	return parcels, nil
}

// зберігаємо посилки у файл
func SaveParcels(parcels []Parcel) error {
	data, err := json.MarshalIndent(parcels, "", "  ")
	if err != nil {
		return err
	}
	return ioutil.WriteFile(fileName, data, 0644)
}

func InitializeStorage() error {
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		initialParcels := []Parcel{
			{ID: "1", Description: "Books", Weight: 2.5, Sender: "Ivan", Receiver: "Pavlo"},
			{ID: "2", Description: "Electronics", Weight: 1.2, Sender: "Maria", Receiver: "Olena"},
			{ID: "3", Description: "Clothes", Weight: 3.0, Sender: "Oleg", Receiver: "Dmytro"},
			{ID: "4", Description: "Toys", Weight: 0.8, Sender: "Anna", Receiver: "Maksym"},
			{ID: "5", Description: "Documents", Weight: 0.5, Sender: "Nadia", Receiver: "Roman"},
		}
		return SaveParcels(initialParcels)
	}
	return nil
}
