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
 
We will go though a real Around advice. This advice prints the taken time by the function.
 
{: .text-yellow-300}
### Prerequisites

Let's check that our environment is ready to follow the tutorial!
 
- Install beyond tool & clone the beyondexamples repository
```bash
>> go get github.com/wesovilabs/beyond
>> git clone https://github.com/wesovilabs/beyondexamples.git
>> cd around
```

{: .text-green-200}
## Let's do it!

{: .text-yellow-300}
### > Define the advice

Around advices must implement the interface Around (`github.com/wesovilabs/beyond/api.Around`).  
```go
type Around interface {
  Before(ctx *context.BeyondContext)
  Returning(ctx *context.BeyondContext)
}
```

Open file [advice/timer.go](https://github.com/wesovilabs/beyondexamples/blob/master/around/advice/timer.go#L20).

```go
const timeStartKey = "time.start"

type TimerMode int32

const (
  Nanoseconds TimerMode = iota
  Microseconds
)

type TimerAdvice struct {
  mode TimerMode
}

func (a *TimerAdvice) Before(ctx *context.BeyondContext) {
  ctx.Set(timeStartKey, time.Now())
}

func (a *TimerAdvice) Returning(ctx *context.BeyondContext) {
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

This is our advice. It implements `Around` interface.

**Method Before**:

It contains the code to be executed before intercepted functions are executed.

**Method Returning**:

It contains the code to be executed after intercepted functions are executed.


{: .text-yellow-300}
### > Register the advice 

- Write a function (or many) that returns the Returning advice

The function signature must be:

```go
func() Around
```

Check the following functions, in file [advice/timer.go](https://github.com/wesovilabs/beyondexamples/blob/master/around/advice/timer.go#L44),

```go
func NewTimerAdvice(mode TimerMode) func() api.Around {
	return func() api.Around{
		return &TimerAdvice{mode}
	}
}
```

Keep in mind that Beyond ignores non-exported functions.

- Register the above function

Open file [cmd/main.go](https://github.com/wesovilabs/beyondexamples/blob/master/around/cmd/main.go#L9) and have a look at function `Beyond()`.

```go
func Beyond() *api.Beyond {
  return api.New().
    WithAround(advice.NewTimerAdvice(advice.Microseconds), "greeting.Hello(string)...").
    WithAround(advice.NewTimerAdvice(advice.Nanoseconds), "greeting.Bye(string)...")
}

func main() {
  greeting.Greetings("Hello", "John")
  greeting.Greetings("Bye", "John")
}
```
Two functions will be intercepted:

- Taken time by function **Hello** in file [greeting/greeting.go](https://github.com/wesovilabs/beyondexamples/blob/master/around/greeting/greeting.go#L8) will be shown in microseconds.
- Taken time by function **Bye** in file [greeting/greeting.go](https://github.com/wesovilabs/beyondexamples/blob/master/around/greeting/greeting.go#L16) will be shown in nanoseconds.

*We will learn more about how to register advices in section [JoinPoint Expressions](/joinpoints)*

{: .text-yellow-300}
### > Beyond in action

This would be the normal behavior

```bash
>> go run cmd/main.go
Hey John
Bye John
```

but when we execute **beyond** command ... (time won't be exactly the same)

```bash
>> beyond run cmd/main.go
Hey John
greeting.Hello(firstName:John) took 37 microseconds
Bye John
greeting.Bye(firstName:John) took 4102 nanoseconds
```

{: .text-green-300}
## Challenge

This time, the challenge must be decided by yourself!!! Extend the TimerAdvice or build a new one that you think it could
be useful for other developers too.

When you complete this challenge, why dont you post an article sharing your experience with Beyond!  I would be very grateful! 

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
