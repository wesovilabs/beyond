---
layout: post
parent: Advices
title: Returning
nav_order: 1
---

{: .text-green-300}
# Returning Advice
{: .fs-9 }

{: .text-green-200}
The guy that always has the last word....
{: .fs-6 .fw-300 }

---

{: .text-green-200}
## About

We will go though an advice that enriches returned errors with info from the function invocation that through
the error.

{: .text-yellow-300}
### Prerequisites

Let's check that our environment is ready to follow the tutorial!
 
- Install goa tool 
```bash
>> go get github.com/wesovilabs/goa
```

- Clone the [goa-examples repository](https://github.com/wesovilabs/goa-examples.git)
```bash
>> git clone https://github.com/wesovilabs/goa-examples.git
>> cd goa-examples
>> git checkout feature/advice-returning
```

{: .text-green-200}
## Let's do it!

{: .text-yellow-300}
### > Define the advice

Returning advices must implement the interface Returning (`github.com/wesovilabs/goa/api.Returning`). 
```go
type Returning interface {
  Returning(ctx *context.GoaContext)
}
```

Let's a have a look at type `ErrorsEnrichAdvice` that is declared in file `advice/errors.go`.

```go
type ErrorsEnrichAdvice struct {}

func (a *ErrorsEnrichAdvice) Returning(ctx *context.GoaContext) {
  if index, result := ctx.Results().Find(func(_ int, arg *context.Arg) bool {
    if val := arg.Value(); val != nil {
      if _, ok := val.(*CustomError);!ok{
        return arg.IsError()
      }
    }
    return false
  });index>=0{
    ctx.Results().SetAt(index, &CustomError{
      err:      result.Value().(error),
      pkg:      ctx.Pkg(),
      function: ctx.Function(),
      params:   ctx.Params(),
    })
  }
}

type CustomError struct {
  err      error
  pkg      string
  function string
  params   *context.Args
}

func (e *CustomError) Error() string {
  params := make([]string, e.params.Count())
  e.params.ForEach(func(index int, arg *context.Arg) {
    params[index] = fmt.Sprintf("%s:%v", arg.Name(), arg.Value())
  })
  errDetail := fmt.Sprintf("%s.%s(%s)", e.pkg, e.function, strings.Join(params, ","))
  return fmt.Sprintf("[%s] => %s", errDetail, e.err.Error())
}
```

**CustomError**

This type will be used to wrap the returning error from the function. As you can observe `CustomError implements interface error`. The
method `Error()` will print the function invocation.


**Type ErrorsEnrichAdvice**

Type that implements interface `Returning` 

**Method Returning**

1. It iterates over the results searching for a result that implements interface `error` and whose values is not nil. In case
of the error was a `CustomError` the function would return false.
```go
func(_ int, arg *context.Arg) bool {
  if val := arg.Value(); val != nil {
    if _, ok := val.(*CustomError);!ok{
      return arg.IsError()
    }
  }
  return false
}
```

2. In case of finding a result, then wwe modify its value. As you see, the returned error is override.
```go
ctx.Results().SetAt(index, &CustomError{
  err:      result.Value().(error),
  pkg:      ctx.Pkg(),
  function: ctx.Function(),
  params:   ctx.Params(),
})
```

**GoaContext** is the guy that provides us with the **joinpoint details**.
You can find the full list of provided methods by GoaContext in section [The GoaContext API](/goacontext)

{: .text-yellow-300}
### > Write a function (or many) that returns the advice
To register a Returning advice,  we need to provide functions that matches with the below signature
```go
func() Returning
```

In file `advice/errors.go` we can find a function that we will use to register the advice.
```go
func NewErrorsEnrichAdviceAdvice() api.Returning {
  return &ErrorsEnrichAdviceAdvice{}
}
```

These functions must be `public`. If not, It will be ignored when registering the advice.

{: .text-yellow-300}
### > Register the advice

Function `Goa() * api.Goa` in file`main.go` is used to register the advices.

```go
func Goa() *api.Goa {
  return api.New().
    WithReturning(advice.NewErrorsEnrichAdviceAdvice, "*.*(...)error")
}
```

Only functions that returns an error will be intercepted

We will dive into registering advices in [JoinPoint Expressions](/joinpoints)


{: .text-yellow-300}
### > Execution

This would be the normal behavior

```bash
>> go run main.go
[ERR] invalid firstName
[ERR] invalid firstName
[ERR] unexpected greeting
```

but if you make use of Goa ...

```bash
>> goa run main.go
[greeting.hello(firstName:)] => [ERR] invalid firstName
[greeting.bye(firstName:)] => [ERR] invalid firstName
[greeting.Greetings(mode:--,firstName:John)] => [ERR] unexpected greeting
```

{: .text-green-300}
## Challenge

You should make the aspect print the full list of functions invocation until the error is thrown. So the output would be

```bash
>> goa run main.go
[greeting.Greetings(mode:hello,firstName:)] => [greeting.hello(firstName:)] => [ERR] invalid firstName
[greeting.Greetings(mode:bye,firstName:)] => [greeting.bye(firstName:)] => [ERR] invalid firstName
[greeting.Greetings(mode:--,firstName:John)] => [ERR] unexpected greeting
```

Just like a tip... maybe you only need to remove a code statemen to complete the challenge... 


If you found any problem to resolve this challenge, don't hesitate to drop me an email at `ivan.corrales.solera@gmail.com` and I will
be happy to give you some help.

If you enjoyed this article, I would really appreciate if you shared it with your networks


<div class="socialme">
    <ul>
        <li class="twitter">
            <a href="https://twitter.com/intent/tweet?via={{site.data.social.twitter.username}}&url={{ site.data.social.twitter.url | uri_escape}}&text={{ site.data.social.twitter.message | uri_escape}}" target="_blank">
                {% include social/twitter.svg %}
            </a>
        </li>
    </ul>
</div>
