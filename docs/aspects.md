---
layout: default
title: Expressions
nav_order: 3
---


{: .text-blue-300}
# Expressions

{: .text-green-200}
Goa provides us a powerful syntax to define which functions will be intercepted by the advices. 
{: .fs-6 .fw-300 }

{: .text-blue-200}
## Introduction

It's important to keep in mind that Goa provides a mechanism to intercept both functions and methods invocations.



{: .text-blue-200}
## Syntax

Goa interprets  the provided expressions in order to decide which functions must be intercepted by the advices.

`<package>.<type>?.<function>(<params)<results>`

* `<type>.` is only required when the advice needs to intercept a method instead of a function. 

We will go through some examples to understand how the aspects expressions work.

| Expression                               | Intercepted               |
|:-----------------------------------------|:--------------------------|:
| `*.*(...)...`                            | Any function invocation with 0 or N params and 0 or N results |
| `*.*.*(...)...`                          | Any method invocation with 0 or N params and 0 or N results |
| `model.*.*(...)...`                      | Any function in invocation, in package `model`,  with 0 or N params and 0 or N results |
| `handlers/employee.*.*(...)...`          | Any function in invocation, in package `handlers/employee`,  with 0 or N params and 0 or N results |
| `model.*.*(...)...`                      | Any method invocation, in package `model`,  with 0 or N params and 0 or N results |
| `model.person.*(...)...`                 | Any method invocation, for type `person` in package `model`,  with 0 or N params and 0 or N results |
| `database.*(string)...`                  | Any function in package `database`, with 1 param of type string and 0 or N results |
| `database.*(string,*int32)...`           | Any function in package `database`, with 2 params of types string and *int32, and 0 or N results |
| `database.*(string,*)...`                | Any function in package `database`, with 2 params of types string and the second param of any type, and 0 or N results |
| `database.*(string,...)...`              | Any function in package `database`, with 2 params of types string and the second param of any type, and 0 or N results |
| `database.*(string,...)func()string`     | Any function in package `database`, with 2 params of types string and the second param of any type, and 1 result whose type is `func()string`|
| `database.set*(*model.Person)...`        | Any function whose name `starts with set` in package `database`, with 1 params of type `*model.Person`, and 1 result whose type is `func()string`|

