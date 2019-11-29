package storage

import (
	"errors"
	"github.com/wesovilabs/beyond/testdata/model"
	"github.com/wesovilabs/beyond/testdata/storage/helper"
)

var database map[string]*model.Person

func SetUpDatabase() {
	database = make(map[string]*model.Person)
}

func InsertPerson(person *model.Person, _ *helper.Test) error {
	uid := helper.RandomUID(8)
	person.ID = uid
	if person.FirstName == "" {
		return errors.New("missing firstName")
	}
	database[uid] = person
	return nil
}

func FindPerson(uid string) (*model.Person, error) {
	if person, ok := database[uid]; ok {
		return person, nil
	}
	return nil, errors.New("person not found")
}

func DeletePerson(uid string) ([]*model.Person, error) {
	if _, ok := database[uid]; ok {
		delete(database, uid)
		return ListPeople()
	}
	return nil, errors.New("person not found")
}

func ListPeople() ([]*model.Person, error) {
	people := make([]*model.Person, 0)
	for _, person := range database {
		people = append(people, person)
	}
	return people, nil
}
