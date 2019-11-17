---
layout: post
parent: Advices
title: Around
nav_order: 3
---

{: .text-green-300}
# Around Advice
{: .fs-9 }
 
{: .text-green-200}
The one that takes the full control of your functions.
{: .fs-6 .fw-300 }

---
 
{: .text-green-200}
## About
 
We will go though an advice that print the spent time by the functions.
 
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
>> git checkout feature/advice-around
 ```

{: .text-green-200}
## Let's do it!

{: .text-yellow-300}
### > Define the advice

Around advices must implement the interface Around (`github.com/wesovilabs/goa/api.Around`).  
```go
type Around interface {
  Before(ctx *context.GoaContext)
  Returning(ctx *context.GoaContext)
}
```

Let's a have a look at type `TimerAdvice` that is declared in file `advice/timer.go`.

```go
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

**Type TimerAdvice** 

This type has an attribute named `mode` that will be used to print the spent time in microseconds or nanoseconds.

**Method Before**:

It set to the GoaContext the time when the function is invoked.

```go
ctx.Set(timeStartKey, time.Now())
```

**Method Returning**:

It calculates the spent time by the function and it prints the result.

1. It recoveries the start time, that was set by the Before method
```go
start := ctx.Get(timeStartKey).(time.Time)
```

2. It calculates the spent time
```go
switch a.mode {
case Nanoseconds:
  timeDuration = fmt.Sprintf("%v nanoseconds\n", time.Since(start).Nanoseconds())
case Microseconds:
  timeDuration = fmt.Sprintf("%v microseconds\n", time.Since(start).Microseconds())
}
```

3. It prints the function invocation (`package.function(params)`) and the spent time by the function.
```go
params := make([]string, ctx.Params().Count())
ctx.Params().ForEach(func(index int, arg *context.Arg) {
  params[index] = fmt.Sprintf("%s:%v", arg.Name(), arg.Value())
})
fmt.Printf("%s.%s(%s) took %s", ctx.Pkg(), ctx.Function(), strings.Join(params, ","),timeDuration)
```

**GoaContext** is the guy that provides us with the **joinpoint details**.
You can find the full list of provided methods by GoaContext in section [The GoaContext API](/goacontext)


{: .text-yellow-300}
### > Write a function (or many) that returns the advice
To register an Around advice,  we need to provide functions that matches with the below signature
```go
func() Around
```

In file `advice/timer.go` we can find the function that will be used to register the advice.
```go
func NewTimerAdvice(mode TimerMode) func() api.Around {
	return func() api.Around{
		return &TimerAdvice{mode}
	}
}
```

This function must be `public`. If not, it will be ignored when registering the advice.

{: .text-yellow-300}
### > Register the advice

Function `Goa() * api.Goa` in file`main.go` is used to register the advices.

```go
func Goa() *api.Goa {
  return api.New().
    WithAround(advice.NewTimerAdvice(advice.Microseconds), "*.Greetings(...)...").
    WithAround(advice.NewTimerAdvice(advice.Nanoseconds), "*.hello(...)...").
    WithAround(advice.NewTimerAdvice(advice.Nanoseconds), "*.bye(...)..."
}
```

As you can guess, the spent time by functions `hello` and `bye` will be shown in nanoseconds and spent time by function
`Greetings` will be shown in microseconds. 


We will dive into registering advices in [JoinPoint Expressions](/joinpoints)

{: .text-yellow-300}
### > Execution

This would be the normal behavior

```bash
>> go run main.go
Hey John
Bye John
```

but if you make use of Goa  the output would look like this...

```bash
>> goa run main.go
greeting.hello(firstName:John) took 31803 nanoseconds
greeting.Greetings(mode:hello,firstName:John) took 49 microseconds
Bye John
greeting.bye(firstName:John) took 4876 nanoseconds
greeting.Greetings(mode:bye,firstName:John) took 12 microseconds
```


{: .text-green-300}
## Challenge

This time, the challenge must be decided by yourself!!! Extend the TimerAdvice or build a new one that you think it could
be useful for other developers too.

Why don't you post an article sharing your experience with Goa?  I would be very grateful! 

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
