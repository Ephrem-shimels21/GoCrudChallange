package models

import (
	"github.com/google/uuid"
)

// Person represents a person object in the system
type Person struct {
	ID      string   `json:"id"` // UUID for unique identifier
	Name    string   `json:"name"`
	Age     int      `json:"age"`
	Hobbies []string `json:"hobbies"`
}

// NewPerson creates a new person instance
func NewPerson(name string, age int, hobbies []string) Person {
	return Person{
		ID:      uuid.New().String(),
		Name:    name,
		Age:     age,
		Hobbies: hobbies,
	}
}
