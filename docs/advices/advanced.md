---
layout: post
parent: Advices
title: Advanced
nav_order: 4
---

{: .text-green-300}
# Advanced usage
{: .fs-9 }

{: .text-green-200}
Beyond the code ... your genius!!
{: .fs-6 .fw-300 }

---

## Skip Pointcut execution

Let's  break the flow! 

{: .text-yellow-300}
### Prerequisites

Let's check that our environment is ready to follow the tutorial!
 
- Install beyond tool & clone the beyond-examples repository
```bash
>> go get github.com/wesovilabs/beyond
>> git clone https://github.com/wesovilabs/beyond-examples.git
>> cd before
```

{: .text-green-200}
## Let's do it!

{: .text-yellow-300}
### > Define the advice


Open file [advice/breaking.go](https://github.com/wesovilabs/beyond-examples/blob/master/skip-pointcut/advice/breaking.go#L10) and have a look at type `BreakingAdvice`.

```go
type BreakingAdvice struct {
  prefix string
}

func (a *BreakingAdvice) Before(ctx *context.BeyondContext) {
  params := make([]string, ctx.Params().Count())
  ctx.Params().ForEach(func(index int, arg *context.Arg) {
    params[index] = fmt.Sprintf("%s:%v", arg.Name(), arg.Value())
  })
}
func (a *BreakingAdvice) Returning(ctx *context.BeyondContext) {
  fmt.Println("Returning")
}

func NewBreakingAdvice() api.Around {
  return &BreakingAdvice{}
}
```
- Register the above functions

Open file [cmd/main.go](https://github.com/wesovilabs/beyond-examples/blob/master/before/cmd/main.go#L9) and have a look at function `Beyond()`.


```go
func Beyond() *api.Beyond {
  return api.New().
    WithAround(advice.NewBreakingAdvice, "greeting.*(...)...")
}

func main() {
  if err:=greeting.Hello("John");err!=nil{
    fmt.Println(err.Error())
  }
  if err:=greeting.Bye("John");err!=nil{
    fmt.Println(err.Error())
  }
}
```

{: .text-yellow-300}
### > Beyond in action

```bash
>> beyond run cmd/main.go
Hey John
Returning
Bye John
Returning
```

That's a normal behavior... but if we modify function Before in file [advice/breaking.go](https://github.com/wesovilabs/beyond-examples/blob/master/skip-pointcut/advice/breaking.go#L10)
```go
func (a *BreakingAdvice) Before(ctx *context.BeyondContext) {
  params := make([]string, ctx.Params().Count())
  ctx.Params().ForEach(func(index int, arg *context.Arg) {
    params[index] = fmt.Sprintf("%s:%v", arg.Name(), arg.Value())
  })
  if ctx.Function()=="Hello"{
    fmt.Println("Leaving the flow")
    ctx.Results().SetAt(0,errors.New("this is awesome!"))
    ctx.Exit()
  }
}
```  

```bash
>> beyond run cmd/main.go
Leaving the flow
this is awesome!
Bye John
Returning
```

Wowwww... invocation to `greeting.Hello` was skipped!!

{: .text-green-300}
## Challenge

By making use of the above example, could you implement an advice that skip joincut
invocations when the given param in function `greeting.Hello`  has already been passed....?

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
