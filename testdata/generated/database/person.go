package database

import "github.com/wesovilabs/goa/testdata/basic/model"

func CreatePerson(p *model.Person) string {
        id := id()
        db.people[id] = p
        return id
}
func ListPeople() []*model.Person {
        people := make([]*model.Person, len(db.people))
        count := 0
        for _, p := range db.people {
                people[count] = p
                count++
        }
        return people
}
