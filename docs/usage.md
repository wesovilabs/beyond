---
layout: default
title: How does Goa work?
nav_order: 5
---

{: .text-green-300}
# How does Goa work?
{: .fs-9 }

{: .text-green-200}
Deep dive into Goa
{: .fs-6 .fw-300 }

---

## Taking registered expressions

Goa reads the registered expressions in your `Goa()` function.

Let's have a look at the below code:

```go
package main

func Goa() *api.Goa {
  return api.New().
    WithBefore(advice.NewTracingAdvice, "greeting.Hello(...)...").
    WithReturning(advice.NewErrorsEnrichAdviceAdvice, "*.*(...)error").
    WithAround(advice.NewTimerAdvice(advice.Nanoseconds), "greeting.Bye(string)...")
}
```

**Advices**

### TracingAdvice

- `advice.NewTracingAdvice` returns an instance of type **TracingAdvice**.
- *TracingAdvice is a Before advice*. It means that Advice that the advice must be executed before a function/method invocation.

The advice will find those functions that march with the expression `greeting.Hello(...)...`.




- `advice.NewErrorsEnrichAdviceAdvice` returns an instance of **ErrorsEnrichAdviceAdvice**.
- `advice.NewTimerAdvice returns` an instance of **TimerAdvice**.



*ErrorsEnrichAdviceAdvice is a Returning advice*. It means that the advice must be executed after a function/method invocation.

*TimerAdvice is an Around advice*. It means that the advice surrounds a function/method invocation.


- advice.NewTracingAdvice needs to be invoked before any function that matches with expression
`greeting.Hello(...)...`

- advice.NewErrorsEnrichAdviceAdvice needs to be invoked after a functions that matches with expression
`*.*(...)error` is executed.

- advice.NewTimerAdvice(advice.Nanoseconds) will be execut




2. Goa examines the AST of your golang project and It produces an unique expression for any function or method in your code

For example, 


2.  



## 
