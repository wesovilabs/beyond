---
layout: post
parent: Advices
grand_parent: Guides & Tutorials
title: GoaContext
nav_order: 4
---

{: .text-green-300}
# GoaContext

{: .text-green-200}
GoaContext is the cornerstone that will help you to build handy and useful advices.
{: .fs-6 .fw-300 }


{: .text-blue-200}
## API

GoaContext provides us with methods to obtain all the information from the intercepted functions.

We can also make use of GoaContext to share data between Before and Returning methods. This is so useful when we code 
an Around Advice.

{: .text-yellow-300}
#### GoaContext

| Method                     | Description               |
|:---------------------------|:-------------------------|:
| Pkg():string               | It returns the package  path |
| Function():string          | It returns the name of the intercepted function|
| Type():interface{}         | It returns the value for the type that contains the intercepted method|
| Params():*Args             | It returns a pointer of [Args](#Args)|
| Results():*Args            | It returns a pointer of [Args](#Args)|
| Set(string,interface{})    | It saves a value that is shareable along the advice cycle|
| Get(string):interface{}    | It obtains a value from the Goaontext|

{: .text-yellow-300}
#### Args

| Method                                    | Description               |
|:------------------------------------------|:-------------------------|:
| ForEach(fn func(int, *Arg))               | It exexutes the provided function for all the arguments |
| Find(func(int,*Arg) bool) (int,*Arg)      | It returns the first index and argument that match with the given function|
| Count() int                               | It returns the number or arguments|
| At(index int) *Arg                        | It returns the [Arg](#Arg) int the given position|
| Get(name string):*Args                    | It returns the [Arg](#Arg) with the given name|
| Set(string,interface{})                   | It update the value for the [Arg](#Arg) with the given name|
| SetAt(int,interface{})                    | It updates the value for the [Arg](#Arg) in the given position|

{: .text-yellow-300}
#### Arg

| Method                                    | Description                          |
|:------------------------------------------|:-------------------------------------|:
| Name():string                             | It returns the name of the argument  |
| Value():interface{}                       | It returns the value of the argument |
| Kind():reflect.Type                       | It returns the type of the argument  |
