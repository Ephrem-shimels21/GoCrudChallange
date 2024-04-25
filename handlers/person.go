package handlers

import (
	"github.com/Ephrem-shimels21/GoCrudChallenge/models"
	"github.com/Ephrem-shimels21/GoCrudChallenge/storage"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// GetPersons returns all persons in the storage
func GetPersons(storage *storage.InMemoryPersonStorage) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		persons := storage.GetPersons()
		ctx.JSON(200, persons)
	}
}

// GetPerson returns a specific person by their ID
func GetPerson(storage *storage.InMemoryPersonStorage) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		person, err := storage.GetPersonByID(id)
		if err != nil {
			ctx.JSON(404, gin.H{"error": "Person not found"})
			return
		}
		ctx.JSON(200, person)
	}
}

// AddPerson creates a new person
func AddPerson(storage *storage.InMemoryPersonStorage) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var newPerson models.Person
		if err := ctx.BindJSON(&newPerson); err != nil {
			ctx.JSON(400, gin.H{"error": "Invalid data"})
			return
		}

		newPerson.ID = uuid.New().String()
		storage.AddPerson(newPerson)

		ctx.JSON(201, newPerson)
	}
}

// UpdatePerson updates an existing person by their ID
func UpdatePerson(storage *storage.InMemoryPersonStorage) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")

		// Get the existing person
		existingPerson, err := storage.GetPersonByID(id)
		if err != nil {
			ctx.JSON(404, gin.H{"error": "Person not found"})
			return
		}

		var updatedFields models.Person
		if err := ctx.BindJSON(&updatedFields); err != nil {
			ctx.JSON(400, gin.H{"error": "Invalid data"})
			return
		}

		// Update only the fields that are present in the request body
		if updatedFields.Name != "" {
			existingPerson.Name = updatedFields.Name
		}
		if updatedFields.Age != 0 {
			existingPerson.Age = updatedFields.Age
		}
		if updatedFields.Hobbies != nil {
			existingPerson.Hobbies = updatedFields.Hobbies
		}

		err = storage.UpdatePerson(id, *existingPerson)
		if err != nil {
			ctx.JSON(404, gin.H{"error": "Person not found"})
			return
		}

		ctx.JSON(200, gin.H{"message": "Person updated successfully", "person": existingPerson})
	}
}

// DeletePerson deletes a specific person by their ID
func DeletePerson(storage *storage.InMemoryPersonStorage) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")

		err := storage.DeletePerson(id)
		if err != nil {
			ctx.JSON(404, gin.H{"error": "Person not found"})
			return
		}

		ctx.JSON(200, gin.H{"message": "Person deleted successfully"})
	}
}
