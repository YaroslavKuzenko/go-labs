package main

import (
	"fmt"
	"strings"
)

// Інтерфейс Displayable
type Displayable interface {
	Display() string
}

// Структура для поштового відділення
type PostOffice struct {
	name    string
	address string
}

func (p PostOffice) Display() string {
	return fmt.Sprintf("Post Office: %s, Address: %s", p.name, p.address)
}

// Структура для клієнта
type Client struct {
	name    string
	parcel  Parcel
	address string
}

func (c Client) Display() string {
	return fmt.Sprintf("Client: %s, Parcel: %s, Address: %s", c.name, c.parcel.Display(), c.address)
}

// Структура для посилки
type Parcel struct {
	id          int
	description string
	weight      float64
}

func (p Parcel) Display() string {
	return fmt.Sprintf("Parcel ID: %d, Description: %s, Weight: %.2fkg", p.id, p.description, p.weight)
}

// Generic структура Stream
type Stream[T Displayable] struct {
	items []T
}

func CreateStream[T Displayable](items []T) Stream[T] {
	return Stream[T]{items}
}

func (s Stream[T]) Filter(predicate func(T) bool) Stream[T] {
	var result []T
	for _, item := range s.items {
		if predicate(item) {
			result = append(result, item)
		}
	}
	return Stream[T]{result}
}

func (s Stream[T]) Map(transform func(T) T) Stream[T] {
	var result []T
	for _, item := range s.items {
		result = append(result, transform(item))
	}
	return Stream[T]{result}
}

func (s Stream[T]) Max(compare func(T, T) bool) T {
	maxItem := s.items[0]
	for _, item := range s.items[1:] {
		if compare(item, maxItem) {
			maxItem = item
		}
	}
	return maxItem
}

func (s Stream[T]) Reduce(accumulator func(T, T) T) T {
	result := s.items[0]
	for _, item := range s.items[1:] {
		result = accumulator(result, item)
	}
	return result
}

func (s Stream[T]) Distinct(keyExtractor func(T) string) Stream[T] {
	seen := make(map[string]bool)
	var result []T
	for _, item := range s.items {
		key := keyExtractor(item)
		if !seen[key] {
			seen[key] = true
			result = append(result, item)
		}
	}
	return Stream[T]{result}
}

func (s Stream[T]) Display() {
	for _, item := range s.items {
		fmt.Println(item.Display())
	}
}

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

	// Використання Stream API
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

	// Використання Stream API
	officeStream.
		Filter(func(p PostOffice) bool { return strings.Contains(p.address, "Main") }).
		Display()
}
