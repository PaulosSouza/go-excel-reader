package repository

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type PersonRepository struct {
	repository *mongo.Collection
}

func (personRepository *PersonRepository) InsertMany(personArr []interface{}, c chan bool) (bool, error) {
	repository := personRepository.repository

	repository.InsertMany(context.TODO(), personArr)

	c <- true

	return true, nil
}

func NewPersonRepository(mongoDbInstance *mongo.Collection) (*PersonRepository) {
	return &PersonRepository{mongoDbInstance}
}
