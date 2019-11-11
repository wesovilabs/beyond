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


We will go through some examples to understand how the aspects expressions work.

| Expression                  | Intercepted               |
|:----------------------------|:--------------------------|:
| `*.*(...)...`               | Any function invocation with 0 or N params and 0 or N results |
| `*.*.*(...)...`             | Any method invocation with 0 or N params and 0 or N results |

