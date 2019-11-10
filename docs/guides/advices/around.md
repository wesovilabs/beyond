---
layout: post
parent: Advices
grand_parent: Guides & Tutorials
title: Around advices
nav_order: 3
---

{: .text-blue-300}
# Around advices

Along this guide we will learn how to code an advice that shows the spent time in any function.

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
### 1. Write a type that implements interface Around

This is the `Around` interface definition 
```go
// Around definition
type Around interface {
	Before(ctx *context.GoaContext)
    Returning(ctx *context.GoaContext)
}
```

Thus, to build an Around advice we need to create a new type that implements both the method  `Before(ctx *context.GoaContext)` and the method `Returning(ctx *context.GoaContext)`

The code in this guide can be found in `advice/timer.go`
```go
package advice

import (
    "fmt"
    "github.com/wesovilabs/goa/api"
    "github.com/wesovilabs/goa/api/context"
)
const timeStartKey = "time.start"
type TimerMode int32
const (
    Nanoseconds TimerMode = iota
    Microseconds
)
type TimerAdvice struct {
    mode TimerMode
}

func (a *TimerAdvice) Before(ctx *context.GoaContext) {
    ctx.Set(timeStartKey, time.Now())
}

func (a *TimerAdvice) Returning(ctx *context.GoaContext) {
    start := ctx.Get(timeStartKey).(time.Time)
    timeDuration:="?"
    switch a.mode {
    case Nanoseconds:
        timeDuration = fmt.Sprintf("%v nanoseconds\n", time.Since(start).Nanoseconds())
    case Microseconds:
        timeDuration = fmt.Sprintf("%v microseconds\n", time.Since(start).Microseconds())
    }
    params := make([]string, ctx.Params().Count())
    ctx.Params().ForEach(func(index int, arg *context.Arg) {
        params[index] = fmt.Sprintf("%s:%v", arg.Name(), arg.Value())
    })
    fmt.Printf("%s.%s(%s) took %s", ctx.Pkg(), ctx.Function(), strings.Join(params, ","),timeDuration)
}


```

{: .text-blue-300}
### 2. Write a function that returns the advice

This function will be used to register the advice

```go
package advice

func NewTimerAdvice(mode TimerMode) func() api.Around {
    return func() api.Around{
        return &TimerAdvice{mode}
    }
}
``` 

{: .text-blue-300}
### 3. Register the advice

Let's open the file `around/main.go` and have a look at function `func Goa() *api.Goa`.  We register advice TimerAdvice for any function that matches with expression `*.*(...)...`.  It actually means
that any function invocation will be intercepted by our advice. We'll learn more about the advice expressions in section [Expressions]()

```go
package main
import (
    "github.com/wesovilabs/goa/api"
    "github.com/wesovilabs/goa/examples/advice"
)
func Goa() *api.Goa {
    return api.New().WithAround(advice.NewTimerAdvice(advice.Microseconds), "*.*(...)...")
}
```


The full code of this example can be found on [Goa repository]()

{: .text-blue-300}
### 4. Code generation

From the root of the goa-examples project, we will run the below command

```bash
>> go generate around/main.go
```

The code is generated in directory `.goa` (by default). 

{: .text-blue-300}
### 5. Running the code

If we run the main function in directory `around` 

```bash
>> go run around/main
Hey John Doe
Hey Jane Doe
Bye John
Bye Jane
```

On the other hand,  if we run the generated code the output will look different

```bash
>> go run .goa/around/main
Hey John Doe
around.sayHello(firstName:John,lastName:Doe) took 33 microseconds
Hey Jane Doe
around.sayHello(firstName:Jane,lastName:Doe) took 6 microseconds
Bye John
around.sayBye(firstName:John) took 2 microseconds
Bye Jane
around.sayBye(firstName:Jane) took 1 microseconds
```

{: .text-blue-300}
## Challenge

As you could realize, `around` is the most powerful type of advice. So this challenge I purpose you to build something
really great.

Would you dare to build a **"Memorization" advice**?  I enum some tips that could help you to code this advice.

- Implement an struct (or use a map) to cache keys and values
- Hash the params from the function to be cached
- What's about TTL? It would be great if functions calls/responses were only cached during a determined period of time.


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
