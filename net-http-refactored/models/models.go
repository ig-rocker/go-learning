package models

import "time"

type Book struct{
	ID string `json:"id"`
	Name string	`json:"name"`
	Author string	`json:"author"`
	DOP time.Time() `json:"date_of_publication"`
}

type BookCollection struct{
	collection map[string]Book
}

