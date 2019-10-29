package testdata

import (
	"github.com/wesovilabs/goa/api"
	"github.com/wesovilabs/goa/api/context"
)

type Person struct {
	firstName string
	age       int
}

func CreatePerson(firstName string, age int) (error, *Person) {
	return nil, &Person{
		firstName: firstName,
		age:       age,
	}
}

type SampleAspect struct {

}

func (a *SampleAspect) Before(ctx *context.Context){

}

func (a *SampleAspect) Returning(ctx *context.Context){

}

func NewSampleAspect() api.Around{
	return &SampleAspect{

	}
}

type SampleBefore struct {

}

func (a *SampleBefore) Before(ctx *context.Context){

}

func NewSampleBefore() api.Before{
	return &SampleBefore{

	}
}

type SampleReturning struct {

}

func (a *SampleReturning) Returning(ctx *context.Context){

}

func NewSampleReturning() api.Returning{
	return &SampleReturning{

	}
}

func Goa()*api.Goa{
	return api.Init().
		WithAround("*.*(...)(error,...)",NewSampleAspect).
		WithBefore("*.*(...)...",NewSampleBefore).
		WithReturning("*.*(...)...",NewSampleReturning)
}
