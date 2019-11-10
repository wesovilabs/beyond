---
layout: post
parent: Advices
grand_parent: Guides & Tutorials
title: Before advices
nav_order: 1
---

{: .text-blue-300}
# Before advices

Along this guide we will learn how to code an advice that traces all the functions invocations.

{: .text-blue-200}
## Pre-requisites

- Clone the [goa-examples repository](https://github.com/wesovilabs/goa-examples.git)
```bash
>> git clone https://github.com/wesovilabs/goa-examples.git
```

- Install goa tool 
```bash
>> go get gihub.com/wesovilabs/goa
```

{: .text-blue-200}
## Go through the code

{: .text-blue-300}
### 1. Write a type that implements interface Before

This is the `Before` interface definition 
```go
// Before definition
type Before interface {
    Before(ctx *context.GoaContext)
}
```

Thus,  to implement a Before advice we just need to create a new type that implements the method `Before(ctx *context.GoaContext)`
It can be found in file `advice/tracing`.

```go
package advice

import (
    "fmt"
    "github.com/wesovilabs/goa/api"
    "github.com/wesovilabs/goa/api/context"
)
type TracingAdvice struct{}
func (c *TracingAdvice) Before(ctx *context.GoaContext) {
    params := make([]string, ctx.ParamsLen())
    ctx.Params().ForEach(func(index int, arg *context.Arg) {
    params[index] = fmt.Sprintf("%s:%v", arg.Name(), arg.Value())
    })
    fmt.Printf("[advice.tracing] %s.%s(%s)\n", ctx.Pkg(), ctx.Function(), strings.Join(params, ","))
}
```

As you can appreciate in the code, we use `ctx.Params().ForEachParam` to iterate over the list of params for the intercepted functions.

You can find the full list of  provided GoaContext methods in section [GoaContext]()


{: .text-blue-300}
### 2. Write a function that returns the advice

This function will be used to register the advice

```go
package advice

func NewTracingAdvice() api.Before {
    return &TracingAdvice{}
}
``` 

{: .text-blue-300}
### 3. Register the advice

Let's open the file `before/main.go` and have a look at function `func Goa() *api.Goa`.  We register advice Tracing for any function that matches with expression `*.*(...)...`.  It actually means
that any function invocation will be intercepted by our advice. We'll learn more about the advice expressions in section [Expressions]()

```go
package main
import (
    "github.com/wesovilabs/goa/api"
    "github.com/wesovilabs/goa/examples/advice"
)
func Goa() *api.Goa {
    return api.New().WithBefore(advice.NewTracingAdvice, "*.*(...)...")
}
```


The full code of this example can be found on [Goa repository]()

{: .text-blue-300}
### 4. Code generation

From the root of the goa-examples project, just run the below command

```bash
>> go generate before/main.go
```

The code is generated in directory `.goa` (by default). 

{: .text-blue-300}
### 5. Running the code

If we run the main function in directory `before` 

```bash
>> go run before/main
Hey John Doe
Hey Jane Doe
Bye John
Bye Jane
```

On the other hand,  if we run the generated code the output will look different

```bash
>> go run .goa/before/main
[advice.tracing] before.sayHello(firstName:John,lastName:Doe)
Hey John Doe
[advice.tracing] before.sayHello(firstName:Jane,lastName:Doe)
Hey Jane Doe
[advice.tracing] before.sayBye(firstName:John)
Bye John
[advice.tracing] before.sayBye(firstName:Jane)
Bye Jane
```

{: .text-blue-300}
## Challenge

I purpose you to implement a new advice to put in practice what we learnt in this article.
 
1. Create a new advice that transform the string params to uppercase. 
2. The new advice must intercept all the functions invocations.
3. Generate the code and then execute it.

The output after generating the code should be 

```bash
>> go run .goa/before/main
[advice.tracing] before.sayHello(firstName:John,lastName:Doe)
Hey JOHN DOE
[advice.tracing] before.sayHello(firstName:Jane,lastName:Doe)
Hey JANE DOE
[advice.tracing] before.sayBye(firstName:John)
Bye JOHN
[advice.tracing] before.sayBye(firstName:Jane)
Bye JANE
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
