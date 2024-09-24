package main

import "fmt"

// Інтерфейс Displayable
type Displayable interface {
	Display() string
}

// Реалізація інтерфейсу Displayable для PostOffice
func (p PostOffice) Display() string {
	return fmt.Sprintf("Post Office: %s, Address: %s", p.name, p.address)
}

// Реалізація інтерфейсу Displayable для Client
func (c Client) Display() string {
	return fmt.Sprintf("Client: %s, Parcel: %s, Address: %s", c.name, c.parcel.Display(), c.address)
}

// Реалізація інтерфейсу Displayable для Parcel
func (p Parcel) Display() string {
	return fmt.Sprintf("Parcel ID: %d, Description: %s, Weight: %.2fkg", p.id, p.description, p.weight)
}
