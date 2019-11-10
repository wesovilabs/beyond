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


## Resource model

GoaContext contains information about the intercepted functions  by the advices. It means, that it contains
the path to the package where the function is found, the name of the function, the list of params and the list of results.

We can make use of GoaContext to share data between Before and Returning methods when you code an Around Advice.


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

#### Args


#### Arg

- Params() *Args: It returns the list of params


##



## Api


- **Pkg()**: It returns the package of the intercepted function.

- **Function()**: It returns the name of the intercepted function.

- **Type()**: It returns the object of the intercepted method.

- **In() \*Args**: It returns a pointer of Args. It contains details from the function arguments.

- **Out() \*Args**: It returns a pointer of Args. It contains details from the function results.

- **Set(key string, value interface{})**: It permits to set custom data to the Goa context.

- **Get(key string) interface{}**: It returns a data that is stored in the Goa Context.
