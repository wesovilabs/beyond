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

We will go though a real Before advice. This advice prints a trace with the function
invocations. 


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

It contains the code to be executed before intercepted functions are invoked.

{: .text-yellow-300}
### > Register the advice 

- Write a function (or many) that returns the Before advice

The functions signature must be:
```go
func() Before
```

Check the following functions, in file [advice/tracing.go](https://github.com/wesovilabs/goa-examples/blob/master/advice/tracing.go),


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

Keep in mind that Goa ignores non-exported functions.

- Register the above functions

Open file [cmd/before/main.go](https://github.com/wesovilabs/goa-examples/blob/master/cmd/before/main.go) and have a look at function `Goa()`.

```go
func Goa() *api.Goa {
  return api.New().
    WithBefore(advice.NewTracingAdvice, "greeting.Hello(...)...").
    WithBefore(advice.NewTracingAdviceWithPrefix("goa"), "greeting.Bye(...)...")
}

func main() {
  greeting.Hello("John")
  greeting.Bye("John")
}
```
Two functions will be intercepted:

- Function `NewTracingAdvice` will be executed before function **Hello** in file [greeting/greeting.go](https://github.com/wesovilabs/goa-examples/blob/master/greeting/greeting.go) is invoked
- Function `NewTracingAdviceWithPrefix` will be executed before **Bye** in file [greeting/greeting.go](https://github.com/wesovilabs/goa-examples/blob/master/greeting/greeting.go) is invoked.

*We will learn more about how to register advices in section [JoinPoint Expressions](/joinpoints)*

{: .text-yellow-300}
### > Goa in action

This would be the normal behavior

```bash
>> go run cmd/before/main.go
Hey John
Bye John
```
but when we execute **goa** command ... 

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
 
- Create a new advice that transforms the string params to uppercase or lowercase. 
- This new advice will be applied to both `greeting.Hello` and `greeting.Bye`  functions. For the Hello function
the advice will transform the retrieved param to uppercase and for the function `Bye ` the param will be transformed
to lowercase.
- The result when running `goa run cmd/before/main.go` must be:
```bash
>> goa run cmd/before/main.go
Hey JOHN
Bye john
```

**Hint** *To face the challenge you could find useful the next functions*

- `ctx.Params().At(index int)`: It returns the `*Arg` in the provided position.
- `ctx.Params().SetAt(index int,value interface{})`: It updates the value for the argument in the provided position.
- `ctx.Params().Get(paramName string)`: It returns the `*Arg` with the provided name.
- `ctx.Params().Set(paramName string,paramValue interface{})`: It updates the value for the argument with the provided name.

*Check sections [GoaContext](/goacontext) for more details.*

If you find any problem to resolve this challenge, don't hesitate to drop me an email at `ivan.corrales.solera@gmail.com` and I will
be happy to give you some help.

---
If you enjoyed this article, I would really appreciate if you share it with your networks


<div class="socialme">
    <ul>
        <li class="twitter">
            <a href="https://twitter.com/intent/tweet?via={{site.data.social.twitter.username}}&url={{ site.data.social.twitter.url | uri_escape}}&text={{ site.data.social.twitter.message | uri_escape}}" target="_blank">
                {% include social/twitter.svg %}
            </a>
        </li>
    </ul>
</div>
