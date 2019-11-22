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

We will go though a real Returning advice. This advice enriches returned errors by the intercepted functions. 


{: .text-yellow-300}
### Prerequisites

Let's check that our environment is ready to follow the tutorial!
 
- Install goa tool & clone the goa-examples repository
```bash
>> go get github.com/wesovilabs/goa
>> git clone https://github.com/wesovilabs/goa-examples.git
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

Open file [advice/error.go](https://github.com/wesovilabs/goa-examples/blob/master/advice/error.go#L10) and have a look at type `ErrorsEnrichAdvice`.

```go
type ErrorsEnrichAdvice struct {
}

func (a *ErrorsEnrichAdvice) Returning(ctx *context.GoaContext) {
  if index, result := ctx.Results().Find(isError);index>=0{
    ctx.Results().SetAt(index, &CustomError{
      err:      result.Value().(error),
      pkg:      ctx.Pkg(),
      function: ctx.Function(),
      params:   ctx.Params(),
    })
  }
}

func isError(_ int, arg *context.Arg) bool{
  if val := arg.Value(); val != nil {
    if _, ok := val.(*CustomError);!ok{
      return arg.IsError()
    }
  }
  return false
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

**Type ErrorsEnrichAdvice**

This is our advice. It implements `Returning` interface. 

**Method Returning**

It contains the code to be executed after intercepted functions are executed.

**CustomError**

It implements interface `error` and It's used to wrap the returning errors by the functions.


{: .text-yellow-300}
### > Register the advice 

- Write a function (or many) that returns the Returning advice

The function signature must be:

```go
func() Returning
```

Check the following functions, in file [advice/error.go](https://github.com/wesovilabs/goa-examples/blob/master/advice/error.go#L50),

```go
func NewErrorsEnrichAdviceAdvice() api.Returning {
  return &ErrorsEnrichAdviceAdvice{}
}
```

Keep in mind that Goa ignores non-exported functions.

- Register the above function

Open file [cmd/returning/main.go](https://github.com/wesovilabs/goa-examples/blob/master/cmd/returning/main.go) and have a look at function `Goa()`.

```go
func Goa() *api.Goa {
  return api.New().
    WithReturning(advice.NewErrorsEnrichAdviceAdvice, "*.*(...)error")
}
func main() {
  checkError(greeting.Greetings("Hello", ""))
  checkError(greeting.Greetings("Bye", ""))
  checkError(greeting.Greetings("--", "John"))
}

func checkError(err error){
  if err!=nil{
    fmt.Println(err.Error())
  }
}
```

- Only functions, with an error result, will be intercepted.

*We will learn more about how to register advices in section [JoinPoint Expressions](/joinpoints)*


{: .text-yellow-300}
### > Goa in action

This would be the normal behavior

```bash
>> go run cmd/returning/main.go
[ERR] invalid firstName
[ERR] invalid firstName
[ERR] unexpected greeting
```

but when we execute **goa** command ...

```bash
>> goa run cmd/returning/main.go
[greeting.Hello(firstName:)] => [ERR] invalid firstName
[greeting.Bye(firstName:)] => [ERR] invalid firstName
[greeting.Greetings(mode:--,firstName:John)] => [ERR] unexpected greeting
```

{: .text-green-300}
## Challenge

- Make changes in code to obtain the below output:

```bash
>> goa run main.go
[greeting.Greetings(mode:hello,firstName:)] => [greeting.hello(firstName:)] => [ERR] invalid firstName
[greeting.Greetings(mode:bye,firstName:)] => [greeting.bye(firstName:)] => [ERR] invalid firstName
[greeting.Greetings(mode:--,firstName:John)] => [ERR] unexpected greeting
```

We should print to the console the full list of invoked function until the error was thrown.

**Hint** *You could do it by removing a code statement* 


If you found any problem to resolve this challenge, don't hesitate to drop me an email at `ivan.corrales.solera@gmail.com` and I will
be happy to give you some help.

---
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
