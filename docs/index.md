---
layout: default
title: Home
description: "Goa is a Go Library that provides AOP"
nav_order: 1
permalink: /
---

{: .text-green-300}
# Go Oriented to Aspects (Goa)
{: .fs-9 }

{: .text-green-200}
A Go library that will drive you to the AOP paradigm.
{: .fs-6 .fw-300 }

[View it on GitHub](https://github.com/wesovilabs/goa){: .btn .fs-5 .mb-4 .mb-md-0 }

---

{: .text-green-300}
## Getting started

{: .text-yellow-300}
### Installation

{: .text-green-200}
#### Install goa
```bash
go get -u github.com/wesovilabs/goa
```

{: .text-green-200}
#### Add goa to your project 

Add goa to go.mod. 


**go.mod**

```text
module github.com/wesovilabs/goa-examples/greetings
...
require github.com/wesovilabs/goa v0.0.1
...
```

Available Goa releases can be found [here](https://github.com/wesovilabs/goa/releases)

{: .text-yellow-300}
### Usage

{: .text-green-200}
#### Advices registration

Let's write a function to register the advices.
```go
package main

func Goa()*api.Goa{
  ...
}
```
- The function must be declared in **main package**.
- The function must be **named Goa**.
- It must **not receive any argument**.
- It must **return a pointer of type Goa** (`*github.com/wesovilabs/goa/api.Goa`).

Have a look at the below example:

```go
package main

import (
   "github.com/wesovilabs/goa/api"
   "github.com/wesovilabs/goa/api/advice"
)

func Goa()*api.Goa{
   return api.New().
      WithBefore(advice.NewTracingAdvice,"*.*(...)...")      		
}
```
There're three types of supported advices:

- **Before**: Advice that is executed before a function/method invocation.
- **Returning**: Advice that is executed after a function/method invocation.
- **Around**: Advice that surrounds a function/method invocation.

We can register as many advices as we need by making use of these methods:

- **withBefore**:
```go
func (g *Goa) WithBefore(func() Before, string) *Goa{
  ...
}
```
You can learn more about writing Before advices [here](/advices/before)

- **withReturning**:
```go
func (g *Goa) WithReturning(func() Returning, string) *Goa {
  ...
}
```
You can learn more about writing Returning advices [here](/advices/returning)

- **withAround**:
```go
func (g *Goa) WithAround(func() Around, string) *Goa {
  ....	
}
```
You can learn more about writing Around advices [here](/advices/around)


As you could realize, the above methods retrieve two params:

-  The first param must be **a function that returns** an
object of type **Before, Returning or Advice**. 
- The second argument must be an expression that will
be used to find the **joinpoints**. In other words, the expressions will be used to define 
which functions or methods must be intercepted by the advices. 

[Write your own advices is very straightforward!](/advices)  
 
---

<div class="socialme">
    <ul>
        <li class="twitter">
            <a href="https://twitter.com/intent/tweet?via={{site.data.social.twitter.username}}&url={{ site.data.social.twitter.url | uri_escape}}&text={{ site.data.social.twitter.message | uri_escape}}" target="_blank">
                {% include social/twitter.svg %}
            </a>
        </li>
    </ul>
</div>
