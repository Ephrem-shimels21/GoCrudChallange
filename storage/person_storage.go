package storage

import (
	"errors"
	"github.com/Ephrem-shimels21/GoCrudChallenge/models"
)

// InMemoryPersonStorage is an in-memory database for persons
type InMemoryPersonStorage struct {
	persons []models.Person
}

// NewInMemoryPersonStorage creates a new in-memory storage instance
func NewInMemoryPersonStorage() *InMemoryPersonStorage {
	return &InMemoryPersonStorage{
		persons: []models.Person{},
	}
}

// GetPersons returns all persons in the storage
func (storage *InMemoryPersonStorage) GetPersons() []models.Person {
	return storage.persons
}

// GetPersonByID retrieves a person by their ID
func (storage *InMemoryPersonStorage) GetPersonByID(id string) (*models.Person, error) {
	for _, person := range storage.persons {
		if person.ID == id {
			return &person, nil
		}
	}
	return nil, errors.New("person not found")
}

// AddPerson adds a new person to the storage
func (storage *InMemoryPersonStorage) AddPerson(person models.Person) {
	storage.persons = append(storage.persons, person)
}

// UpdatePerson updates an existing person
func (storage *InMemoryPersonStorage) UpdatePerson(id string, updatedPerson models.Person) error {
	for i, person := range storage.persons {
		if person.ID == id {
			storage.persons[i] = updatedPerson
			return nil
		}
	}
	return errors.New("person not found")
}

// DeletePerson removes a person from the storage by their ID
func (storage *InMemoryPersonStorage) DeletePerson(id string) error {
	for i, person := range storage.persons {
		if person.ID == id {
			storage.persons = append(storage.persons[:i], storage.persons[i+1:]...)
			return nil
		}
	}
	return errors.New("person not found")
}

// type PersonDB struct {
//     mu      sync.RWMutex
//     persons map[string]models.Person
// }

// func NewPersonDB() *PersonDB {
//     return &PersonDB{
//         persons: make(map[string]models.Person),
//     }
// }

// func (db *PersonDB) GetPerson(id string) (models.Person, bool) {
//     db.mu.RLock()
//     defer db.mu.RUnlock()

//     person, exists := db.persons[id]
//     return person, exists
// }

// func (db *PersonDB) GetAllPersons() []models.Person {
//     db.mu.RLock()
//     defer db.mu.RUnlock()

//     var allPersons []models.Person
//     for _, person := range db.persons {
//         allPersons = append(allPersons, person)
//     }
//     return allPersons
// }

// func (db *PersonDB) CreatePerson(person models.Person) {
//     db.mu.Lock()
//     defer db.mu.Unlock()

//     db.persons[person.ID] = person
// }

// func (db *PersonDB) UpdatePerson(id string, person models.Person) {
//     db.mu.Lock()
//     defer db.mu.Unlock()

//     db.persons[id] = person
// }

// func (db *PersonDB) DeletePerson(id string) {
//     db.mu.Lock()
//     defer db.mu.Unlock()

//     delete(db.persons, id)
// }
