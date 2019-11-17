---
layout: post
title: GoaContext
nav_order: 4
---

{: .text-green-300}
# GoaContext
{: .fs-9 }

{: .text-green-200}
The cornerstone that will help you to build handy and useful advices.
{: .fs-6 .fw-300 }

---

{: .text-green-300}
## API

GoaContext provides us with methods to obtain all the information from the intercepted functions.

We can also make use of GoaContext to share data between Before and Returning methods (It's so useful when coding 
 Around advices).

{: .text-yellow-300}
### GoaContext

| Method                     | Description               |
|:---------------------------|:-------------------------|:
| Pkg():string               | It returns the package  path |
| Function():string          | It returns the name of the intercepted function|
| Type():interface{}         | It returns the value for the type that contains the intercepted method|
| Params():*Args             | It returns a pointer of [Args](#args)|
| Results():*Args            | It returns a pointer of [Args](#args)|
| Set(string,interface{})    | It saves a value that is shareable along the advice cycle|
| Get(string):interface{}    | It obtains a value from the GoaContext|

{: .text-yellow-300}
### Args

| Method                                    | Description               |
|:------------------------------------------|:-------------------------|:
| ForEach(fn func(int, *Arg))               | It exexutes the provided function for all the arguments |
| Find(func(int,*Arg) bool) (int,*Arg)      | It returns the first index and argument that match with the given function|
| Count() int                               | It returns the number or arguments|
| At(index int) *Arg                        | It returns the [Arg](#arg) int the given position|
| Get(name string):*Args                    | It returns the [Arg](#arg) with the given name|
| Set(string,interface{})                   | It update the value for the [Arg](#arg) with the given name|
| SetAt(int,interface{})                    | It updates the value for the [Arg](#arg) in the given position|

{: .text-yellow-300}
### Arg

| Method                                    | Description                          |
|:------------------------------------------|:-------------------------------------|:
| Name():string                             | It returns the name of the argument  |
| Value():interface{}                       | It returns the value of the argument |
| Kind():reflect.Type                       | It returns the type of the argument  |
