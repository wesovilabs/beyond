---
layout: page
title: Getting Started
permalink: /getting-started/
nav_order: 2
---

# Getting Started

## Installation

Add Goa dependency to your go.mod

**go.mod**
```text
module github.com/wesovilabs/goa-examples/greetings
...
require github.com/wesovilabs/goa <goa.version>
...
```

Replace `<goa.version>` by the one that you desire. Available Goa versions are available [here](https://github.com/wesovilabs/goa/releases)

Stable Version is:

v0.0.1 
{: .label .label-green } 


In case of your project dependencies were managed with Glide or Go Dep you could add Goas as below: 

**Glide**
```bash
glide get github.com/wesovilabs/goa
```

**Go dep**
```bash
go get github.com/wesovilabs/goa
```

## Usage

### Registering aspects

Define a function named Goa that returns a `*api.Goa`. In this function
we will define which aspects (before,returning or/and around) will be applied
to the functions that match with the provided expression.

```go
package main

import "github.com/wesovilabs/goa/api"

func Goa()*api.Goa{
	return api.Init().
		WithBefore("*.*(...)...",TracingAspect).
		WithAround("*.StringUtils.*(string)string",MemorizeAspect)
}
```

To learn more about aspects, go to section [Aspects]
