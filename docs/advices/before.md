---
layout: post
parent: Advices
title: Before
nav_order: 1
---

{: .text-green-300}
# Before Advice
{: .fs-9 }

{: .text-green-200}
Born to be the bouncer of your functions
{: .fs-6 .fw-300 }

---

{: .text-green-200}
## About

Along this section, we wil learn to implement an advice that will
print the function invocations. 


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

Before advices must implement the interface Before (`github.com/wesovilabs/goa/api.Before`). 
```go
type Before interface {
  Before(ctx *context.GoaContext)
}
```

Open file [advice/tracing.go](https://github.com/wesovilabs/goa-examples/blob/master/advice/tracing.go) and have a look at type `TracingAdvice`.

```go
type TracingAdvice struct {
  prefix string
}

func (a *TracingAdvice) Before(ctx *context.GoaContext) {
  params := make([]string, ctx.Params().Count())
  ctx.Params().ForEach(func(index int, arg *context.Arg) {
    params[index] = fmt.Sprintf("%s:%v", arg.Name(), arg.Value())
  })
  printTrace(ctx,a.prefix,params)
}
```

**Type TracingAdvice** 

This is our advice. We can build more reusable and customizable advices by making use of attributes (`TracingAdvice` has a `prefix` attribute)

**Method Before**: 

We need to implement this since is required by interface `Before`

1. Define a list to put the params info (`name:value`)
```go 
params := make([]string, ctx.Params().Count())
```
2. Iterate over the params and put them into the list
```go
ctx.Params().ForEach(func(index int, arg *context.Arg) {
  params[index] = fmt.Sprintf("%s:%v", arg.Name(), arg.Value())
})
```
3. Call function the will print the trace
```go
printTrace(ctx,a.prefix,params)
```

**GoaContext** is the guy that provides us with the **joinpoint details**.
You can find the full list of provided methods by GoaContext in section [The GoaContext API](/goacontext)

{: .text-yellow-300}
### > Write a function (or many) that returns the advice
To register a Before advice,  we need to provide functions that matches with the below signature
```go
func() Before
```

In this file we see the next functions:
```go
func NewTracingAdvice() api.Before {
  return &TracingAdvice{}
}

func NewTracingAdviceWithPrefix(prefix string) func() api.Before {
  return func() api.Before {
    return &TracingAdvice{
      prefix: prefix,
    }
  }
}
```

**IMPORTANT**: These functions must be `public`. 

{: .text-yellow-300}
### > Register the advice

Open file [cmd/before/main.go](https://github.com/wesovilabs/goa-examples/blob/master/cmd/before/main.go) and have a look at type `TracingAdvice`.

```go
func Goa() *api.Goa {
  return api.New().
    WithBefore(advice.NewTracingAdvice, "*.*(...)*").
    WithBefore(advice.NewTracingAdviceWithPrefix("[goa]"), "*.Bye(...)error")
}
```

Function `Bye` invocations will be traced with `[goa]` prefix and the function `Hello` won't.

We will dive into registering advices in [JoinPoint Expressions](/joinpoints)

{: .text-yellow-300}
### > Execution

This would be the normal behavior

```bash
>> go run cmd/before/main.go
Hey John
Bye John
```

but if you make use of Goa ...

```bash
>> goa run main.go
greeting.Hello(firstName:John)
Hey John
[goa] greeting.Bye(firstName:John)
Bye John
```

{: .text-green-300}
## Challenge

I purpose you to implement a new advice to put in practice what we learnt in this article.
 
1. Create a new advice that transforms the string params to uppercase or lowercase. 
2. This new advice will be applied to both `greeting.Hello` and `greeting.Bye`  functions. For the Hello function
the advice will transform the retrieved param to uppercase and for the function `Bye ` the param will be transformed
to lowercase.

The output must be 

```bash
>> goa run main.go
Hey JOHN
Bye john
```

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
