package main

type Parcel struct {
	ID          int     `json:"id"`
	Description string  `json:"description"`
	Weight      float64 `json:"weight"`
	Sender      string  `json:"sender"`
	Receiver    string  `json:"receiver"`
}
