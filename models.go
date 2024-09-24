package main

type PostOffice struct {
	name    string
	address string
}

type Client struct {
	name    string
	parcel  Parcel
	address string
}

type Parcel struct {
	id          int
	description string
	weight      float64
}
