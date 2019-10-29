package testdata

import (
        "github.com/wesovilabs/goa/api"
        "github.com/wesovilabs/goa/api/context"
        "github.com/wesovilabs/goa/wrapper/testdata"
)

type Person struct {
        firstName string
        age       int
}

func CreatePerson(firstName string, age int) (error, *Person) {
        return nil, &Person{firstName: firstName, age: age}
}

type SampleAspect struct{}

func (a *SampleAspect) Before(ctx *context.Context) {
}
func (a *SampleAspect) Returning(ctx *context.Context) {
}
func NewSampleAspect() api.Around {
        return &SampleAspect{}
}

type SampleBefore struct{}

func (a *SampleBefore) Before(ctx *context.Context) {
}
func NewSampleBefore() api.Before {
        return &SampleBefore{}
}

type SampleReturning struct{}

func (a *SampleReturning) Returning(ctx *context.Context) {
}
func NewSampleReturning() api.Returning {
        return &SampleReturning{}
}
func Goa() *api.Goa {
        return api.Init().WithAround("*.*(...)(error,...)", NewSampleAspect).WithBefore("*.*(...)...", NewSampleBefore).WithReturning("*.*(...)...", NewSampleReturning)
}
func CreatePersonWrapper(firstName string, age int) (error, *Person) {
        goaContext := goaContext.NewContext(context.Background())
        aspect0 := NewSampleReturning()
        aspect1 := NewSampleBefore()
        aspect2 := NewSampleAspect()
        goaContext.In().Set("firstName", firstName)
        goaContext.In().Set("age", age)
        aspect1.Before(goaContext)
        aspect2.Before(goaContext)
        firstName = goaContext.Out().Get("firstName").(string)
        age = goaContext.Out().Get("age").(int)
        result0, result1 := testdata.CreatePerson(firstName, age)
        goaContext.Out().Set("result0", result0)
        goaContext.Out().Set("result1", result1)
        aspect0.Returning(goaContext)
        aspect2.Returning(goaContext)
        result0 = goaContext.Out().Get("result0").(error)
        result1 = goaContext.Out().Get("result1").(*Person)
        return result0, result1
}
