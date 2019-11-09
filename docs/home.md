---
layout: default
title: Home
description: "Goa is a Golang Library that provides AOP"
nav_order: 1
permalink: /home
---

{: .text-green-300}
# Golang Oriented to Aspects (Goa)
{: .fs-9 }

{: .text-green-200}
A Golang library that drives us to the AOP paradigm.
{: .fs-6 .fw-300 }

[Get started now](#getting-started){: .btn .btn-primary .fs-5 .mb-4 .mb-md-0 .mr-2 } [View it on GitHub](https://github.com/wesovilabs/goa){: .btn .fs-5 .mb-4 .mb-md-0 }

---

{: .text-green-300}
## Getting started

{: .text-yellow-300}
### Installation

Add goa to go.mod. 


**go.mod**

```text
module github.com/wesovilabs/goa-examples/greetings
...
require github.com/wesovilabs/goa master
...
```

Available Goa releases can be found [here](https://github.com/wesovilabs/goa/releases)

{: .text-yellow-300}
### Usage

{: .text-green-200}
#### Aspects registration

Aspects are registered programmatically in a function whose signature will be the below
```go
func Goa()*api.Goa{...}
```
The function could be found in any package in our project. On the other hand, 
and just for convection, I suggest you to write this function in your main package 

- The function must be named Goa.
- It won't retrieve any argument.
- It returns a pointer of Goa (`*github.com/wesovilabs/goa/api.Goa`).

A full example is shown in this example

```go
package main

import (
   "github.com/wesovilabs/goa/api"
   "github.com/wesovilabs/goa/aspects"
)

func Goa()*api.Goa{
   return api.New().
      WithBefore(aspects.TracingBasic,"*.*(...)...")      		
}
```

There're three types of advices:

- **Before**: Advice that executes before a function/method invocation.
- **Returning**: Advice that executes after a function/method invocation.
- **Around**: Advice that surrounds a function/method invocation.

[Tracing]() and some others advices are provided out of the box by Goa. The full list of provided aspects
can be found in section [Advices by Goa]().

Write your own advices is very straightforward!  

[Write your own advice](#getting-started){: .btn .btn-secondary .fs-5 .mb-4 .mb-md-0 .mr-2 } 
---

