package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"log"
)

var db *pgx.Conn

func InitializeDatabase() error {
	var err error
	dsn := "postgres://postgres:postgres@localhost:5432/parcel_db"
	db, err = pgx.Connect(context.Background(), dsn)
	if err != nil {
		return fmt.Errorf("не вдалося підключитися до бази даних: %w", err)
	}
	log.Println("Підключення до бази даних успішне")
	return nil
}

func CloseDatabase() {
	if db != nil {
		db.Close(context.Background())
		log.Println("Підключення до бази даних закрите")
	}
}

func LoadParcels() ([]Parcel, error) {
	rows, err := db.Query(context.Background(), "SELECT id, description, weight, sender, receiver FROM parcels")
	if err != nil {
		return nil, fmt.Errorf("помилка при завантаженні посилок: %w", err)
	}
	defer rows.Close()

	var parcels []Parcel
	for rows.Next() {
		var parcel Parcel
		if err := rows.Scan(&parcel.ID, &parcel.Description, &parcel.Weight, &parcel.Sender, &parcel.Receiver); err != nil {
			return nil, fmt.Errorf("помилка при читанні рядка: %w", err)
		}
		parcels = append(parcels, parcel)
	}
	return parcels, nil
}

func SaveParcel(parcel Parcel) error {
	_, err := db.Exec(context.Background(), "INSERT INTO parcels (description, weight, sender, receiver) VALUES ($1, $2, $3, $4)",
		parcel.Description, parcel.Weight, parcel.Sender, parcel.Receiver)
	if err != nil {
		return fmt.Errorf("помилка при збереженні посилки: %w", err)
	}
	return nil
}

func UpdateParcel(parcel Parcel) error {
	_, err := db.Exec(context.Background(), "UPDATE parcels SET description=$1, weight=$2, sender=$3, receiver=$4 WHERE id=$5",
		parcel.Description, parcel.Weight, parcel.Sender, parcel.Receiver, parcel.ID)
	if err != nil {
		return fmt.Errorf("помилка при оновленні посилки: %w", err)
	}
	return nil
}

func DeleteParcel(id int) error {
	_, err := db.Exec(context.Background(), "DELETE FROM parcels WHERE id=$1", id)
	if err != nil {
		return fmt.Errorf("помилка при видаленні посилки: %w", err)
	}
	return nil
}
