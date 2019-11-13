//go:generate  go run github.com/wesovilabs/goa --project github.com/wesovilabs/goa/testdata/basic   --output ../generated --verbose
package main

import (
	"fmt"
	"github.com/wesovilabs/goa/api"
	"github.com/wesovilabs/goa/testdata/basic/aspects"
	"github.com/wesovilabs/goa/testdata/basic/database"
	"github.com/wesovilabs/goa/testdata/basic/model"
)

var beforeFn = func() api.Before {
	return nil
}

func Goa() *api.Goa {
	return api.New().
		WithBefore("*database.CreatePerson(*model.Person)string", aspects.NewTracingAspect)
	/**
	WithReturning("*.Create*(...)string", aspects.NewNormalizeID).
	WithAround("*.*(*)...", time.TimingMicroSeconds).
	WithAround("*.ListPeople()...", time.TimingNanoSeconds).
	WithBefore("*.*(...)...", aspects.NewTracingAspect)
	*/
}

func main() {
	person := model.NewPerson("John", 20)
	fmt.Println("-----------")
	personID := database.CreatePerson(person)
	fmt.Printf("Created person with id %s\n", personID)
	fmt.Println("-----------")
	people := database.ListPeople()
	fmt.Printf("There are %v people\n", len(people))
	fmt.Println("-----------")

}
