package main

import (
	"fmt"
	"strings"
)

// Функція для об'єднання клієнтів за їхніми іменами (для Reduce)
func combineClients(c1, c2 Client) Client {
	return Client{
		name:    c1.name + " & " + c2.name,
		parcel:  c1.parcel, // Беремо першу посилку для демонстрації
		address: c1.address,
	}
}

// Функція для порівняння клієнтів за вагою посилки (для Max)
func compareByWeight(c1, c2 Client) bool {
	return c1.parcel.weight > c2.parcel.weight
}

// Функція для отримання унікальних клієнтів за іменем (для Distinct)
func extractClientName(c Client) string {
	return c.name
}

func main() {
	// Створення набору даних
	parcel1 := Parcel{id: 101, description: "Books", weight: 2.5}
	parcel2 := Parcel{id: 102, description: "Electronics", weight: 1.2}
	parcel3 := Parcel{id: 103, description: "Clothes", weight: 3.0}
	parcel4 := Parcel{id: 104, description: "Toys", weight: 0.8}

	client1 := Client{name: "Ivan", parcel: parcel1, address: "Main St. 123"}
	client2 := Client{name: "Pavlo", parcel: parcel2, address: "Green St. 45"}
	client3 := Client{name: "Alina", parcel: parcel3, address: "Park St. 67"}
	client4 := Client{name: "Ivan", parcel: parcel4, address: "Main St. 456"} // Дублікат за ім'ям

	postOffice1 := PostOffice{name: "Post Office 1", address: "Main St. 1"}
	postOffice2 := PostOffice{name: "Post Office 2", address: "Green St. 2"}

	// Створюємо Stream для клієнтів
	clientStream := CreateStream([]Client{client1, client2, client3, client4})

	// 1. Використання Filter (фільтрація за адресою)
	fmt.Println("Filter: клієнти з адресою, що містить 'Main'")
	clientStream.
		Filter(func(c Client) bool { return strings.Contains(c.address, "Main") }).
		Display()

	fmt.Println("-----")

	// 2. Використання Map (зміна опису посилки)
	fmt.Println("Map: додавання до опису посилки слова 'Updated'")
	clientStream.
		Map(func(c Client) Client {
			c.parcel.description = "Updated " + c.parcel.description
			return c
		}).
		Display()

	fmt.Println("-----")

	// 3. Використання Max (пошук клієнта з найважчою посилкою)
	fmt.Println("Max: клієнт з найважчою посилкою")
	heaviestClient := clientStream.Max(compareByWeight)
	fmt.Println(heaviestClient.Display())

	fmt.Println("-----")

	// 4. Використання Reduce (об'єднання імен клієнтів)
	fmt.Println("Reduce: об'єднання імен клієнтів")
	reducedClient := clientStream.Reduce(combineClients)
	fmt.Println(reducedClient.Display())

	fmt.Println("-----")

	// 5. Використання Distinct (уникнення дублікатів за ім'ям)
	fmt.Println("Distinct: унікальні клієнти за іменем")
	clientStream.
		Distinct(extractClientName).
		Display()

	fmt.Println("-----")

	// Створюємо Stream для поштових відділень
	officeStream := CreateStream([]PostOffice{postOffice1, postOffice2})

	// Використання Filter для поштових відділень
	fmt.Println("Post Offices with 'Main' in the address:")
	officeStream.
		Filter(func(p PostOffice) bool { return strings.Contains(p.address, "Main") }).
		Display()
}
