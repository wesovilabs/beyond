package main

import (
	"context"
	"fmt"
	"github.com/wesovilabs/goa/api"
	goaContext "github.com/wesovilabs/goa/api/context"
)

type Person struct {
	firstName string
	age       int
}

func CreatePerson(firstName string, age int) (error, *Person) {
	return nil, &Person{firstName: firstName, age: age}
}

type SampleAspect struct {
}

func (a *SampleAspect) Before(ctx *goaContext.Context) {
}
func (a *SampleAspect) Returning(ctx *goaContext.Context) {
	ctx.Out().Set("out_1", &Person{
	        firstName: "Ivan",
	        age:34,
    })
}
func NewSampleAspect() api.Around {
	return &SampleAspect{}
}
func Goa() *api.Goa {
	return api.Init().WithAround("*.*(...)(error,...)", NewSampleAspect)
}
func CreatePerson_wrapper(firstName string, age int) (error, *Person) {
	aspect, goaContext := NewSampleAspect(), goaContext.NewContext(context.Background())
	goaContext.In().Set("firstName", firstName)
	goaContext.In().Set("age", age)
	aspect.Before(goaContext)
	firstName = goaContext.In().Get("firstName").(string)
	age = goaContext.In().Get("age").(int)
	out_0, out_1 := CreatePerson(firstName, age)
	goaContext.Out().Set("out_0", out_0)
	goaContext.Out().Set("out_1", out_1)
	aspect.Returning(goaContext)
	if val := goaContext.Out().Get("out_0"); val != nil {
		out_0 = val.(error)
	} else {
		out_0 = nil
	}
	if val := goaContext.Out().Get("out_1"); val != nil {
		out_1 = val.(*Person)
	} else {
		out_1 = nil
	}
	return out_0, out_1
}

func main() {
	_, p := CreatePerson("John", 22)
	fmt.Println(p)
	_, p = CreatePerson_wrapper("John", 22)
	fmt.Println(p)
}
