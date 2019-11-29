---
layout: default
title: Home
description: "Beyond is a Go Library that provides AOP"
nav_order: 1
permalink: /
---

{: .text-green-300}
# Go Oriented to Aspects (Beyond)
{: .fs-9 }

{: .text-green-200}
A Go library that will drive you to the AOP paradigm.
{: .fs-6 .fw-300 }

[View it on GitHub](https://github.com/wesovilabs/beyond){: .btn .fs-5 .mb-4 .mb-md-0 }

---

{: .text-green-300}
## Getting started

{: .text-yellow-300}
### Installation

{: .text-green-200}
#### Install beyond
```bash
go get -u github.com/wesovilabs/beyond
```

{: .text-green-200}
#### Add beyond to your project 

Add beyond to go.mod. 


**go.mod**

```text
module github.com/wesovilabs/beyond-examples/greetings
...
require github.com/wesovilabs/beyond v0.0.1
...
```

Available Beyond releases can be found [here](https://github.com/wesovilabs/beyond/releases)

{: .text-yellow-300}
### Usage

{: .text-green-200}
#### Advices registration

Let's write a function to register the advices.
```go
package main

func Beyond()*api.Beyond{
  ...
}
```
- The function must be declared in **main package**.
- The function must be **named Beyond**.
- It must **not receive any argument**.
- It must **return a pointer of type Beyond** (`*github.com/wesovilabs/beyond/api.Beyond`).

Have a look at the below example:

```go
package main

import (
   "github.com/wesovilabs/beyond/api"
   "github.com/wesovilabs/beyond/api/advice"
)

func Beyond()*api.Beyond{
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
func (g *Beyond) WithBefore(func() Before, string) *Beyond{
  ...
}
```
You can learn more about writing Before advices [here](/advices/before)

- **withReturning**:
```go
func (g *Beyond) WithReturning(func() Returning, string) *Beyond {
  ...
}
```
You can learn more about writing Returning advices [here](/advices/returning)

- **withAround**:
```go
func (g *Beyond) WithAround(func() Around, string) *Beyond {
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
