package main

import (
	"fmt"
	"strings"
)

func main() {
	// Створення набору даних
	parcel1 := Parcel{id: 101, description: "Books", weight: 2.5}
	parcel2 := Parcel{id: 102, description: "Electronics", weight: 1.2}
	parcel3 := Parcel{id: 103, description: "Clothes", weight: 3.0}

	client1 := Client{name: "Ivan", parcel: parcel1, address: "Main St. 123"}
	client2 := Client{name: "Pavlo", parcel: parcel2, address: "Green St. 45"}
	client3 := Client{name: "Alina", parcel: parcel3, address: "Park St. 67"}

	postOffice1 := PostOffice{name: "Post Office 1", address: "Main St. 1"}
	postOffice2 := PostOffice{name: "Post Office 2", address: "Green St. 2"}

	// Створюємо Stream для клієнтів
	clientStream := CreateStream([]Client{client1, client2, client3})

	// Використання Stream API для клієнтів
	clientStream.
		Filter(func(c Client) bool { return strings.Contains(c.address, "Main") }).
		Map(func(c Client) Client {
			c.address = "Updated " + c.address
			return c
		}).
		Display()

	fmt.Println("-----")

	// Створюємо Stream для поштових відділень
	officeStream := CreateStream([]PostOffice{postOffice1, postOffice2})

	// Використання Stream API для поштових відділень
	officeStream.
		Filter(func(p PostOffice) bool { return strings.Contains(p.address, "Main") }).
		Display()
}
