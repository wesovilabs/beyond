---
layout: post
parent: Advices
grand_parent: Guides & Tutorials
title: Returning
nav_order: 1
---

{: .text-blue-300}
# Returning advices

Along this guide we will learn how to code an advice that adds information, about the function that returns errors, to
the error itself. 

{: .text-blue-200}
## Pre-requisites

- Clone the [goa-examples repository](https://github.com/wesovilabs/goa-examples.git)
```bash
>> git clone https://github.com/wesovilabs/goa-examples.git
```

- Install goa tool 
```bash
>> go get github.com/wesovilabs/goa
```

{: .text-blue-200}
## Go through the code

{: .text-blue-300}
### 1. Write a type that implements interface Returning

This is the `Returning` interface definition 
```go
// Returning definition
type Returning interface {
    Returning(ctx *context.GoaContext)
}
```

Thus,  to implement a Returning advice we just need to create a new type that implements the method `Returning(ctx *context.GoaContext)`
It can be found in file `advice/errors`.
```go
package advice

import (
    "fmt"
    "github.com/wesovilabs/goa/api"
    "github.com/wesovilabs/goa/api/context"
)
type ErrorAdvice struct {
}

func (a *ErrorAdvice) Returning(ctx *context.GoaContext) {
    if index, result := ctx.Results().Find(func(_ int, arg *context.Arg) bool {
        if val := arg.Value(); val != nil {
            _, ok := val.(*InterceptedError)
            return !ok && arg.IsError()
        }
        return false
    });index>=0 {
        ctx.Results().SetAt(index, &InterceptedError{
            err:      result.Value().(error),
            pkg:      ctx.Pkg(),
            function: ctx.Function(),
            params:   ctx.Params(),
        })
    }
}
```

The `InterceptedError` type will be used to override interface error. 
```go
type InterceptedError struct {
	err      error
	pkg      string
	function string
	params   *context.Args
}

func (e *InterceptedError) Error() string {
	params := make([]string, e.params.Count())
	e.params.ForEach(func(index int, arg *context.Arg) {
		params[index] = fmt.Sprintf("%s:%v", arg.Name(), arg.Value())
	})
	errDetail := fmt.Sprintf("%s.%s(%s)", e.pkg, e.function, strings.Join(params, ","))
	return fmt.Sprintf("[%s] => '%s'", errDetail, e.err.Error())
}
```


{: .text-blue-300}
### 2. Write a function that returns the advice

This function will be used to register the advice

```go
package advice

func NewErrorAdvice() api.Before {
    return &ErrorAdvice{}
}
``` 

{: .text-blue-300}
### 3. Register the advice

Let's open the file `returning/main.go` and have a look at function `func Goa() *api.Goa`.  We register advice ErrorAdvice for any function that matches with expression `*.*(...)...`.  It actually means
that any function invocation will be intercepted by our advice. We'll learn more about the advice expressions in section [Expressions]()

```go
package main
import (
    "github.com/wesovilabs/goa/api"
    "github.com/wesovilabs/goa/examples/advice"
)
func Goa() *api.Goa {
    return api.New().WithReturning(advice.NewErrorAdvice, "*.*(...)...")
}
```


The full code of this example can be found on [Goa repository]()

{: .text-blue-300}
### 4. Code generation

From the root of the goa-examples project, we will run the below command

```bash
>> go generate returning/main.go
```

The code is generated in directory `.goa` (by default). 

{: .text-blue-300}
### 5. Running the code

If we run the main function in directory `returning` 

```bash
>> go run returning/main
invalid firstName
unexpected greeting
invalid firstName
How're you, John?
```

On the other hand,  if we run the generated code the output will look different

```bash
>> go run .goa/before/main
[returning.sayHello(firstName:)] => invalid firstName
[returning.greetings(mode:unknown,firstName:John)] => unexpected greeting
[returning.AskTo(firstName:)] => invalid firstName
How're you, John?
```

{: .text-blue-300}
## Challenge

The ErrorAdvice, it just prints the function call that returns the error. I purpose you to make the advice
 print the full list of functions invocations that were called until the error was returned.

The output after generating the code should be 

```bash
[returning.greetings(mode:hello,firstName:)] => [returning.sayHello(firstName:)] => invalid firstName
[returning.greetings(mode:unknown,firstName:John)] => unexpected greeting
[returning.AskTo(firstName:)] => invalid firstName
How're you, John?
```

Be focused on first line 
```bash
[returning.greetings(mode:hello,firstName:)] => [returning.sayHello(firstName:)] => invalid firstName
```

If you find any problem to resolve this challenge, just drop me an email at `ivan.corrales.solera@gmail.com` and I will
be happy to give you some help.


If you enjoyed this article, I would really appreciate if you shared it with your networks


<div class="socialme">
    <ul>
        <li class="twitter">
            <a href="https://twitter.com/intent/tweet?via={{site.data.social.twitter.username}}&url={{ site.data.social.twitter.url | uri_escape}}&text={{ site.data.social.twitter.message2 | uri_escape}}" target="_blank">
                {% include social/twitter.svg %}
            </a>
        </li>
    </ul>
</div>
