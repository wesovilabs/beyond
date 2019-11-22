---
layout: page
title: Goa
permalink: /goa/
nav_order: 6
---

{: .text-green-300}
# >> goa ...
{: .fs-9 }

{: .text-green-200}
The magic wand that will make it happen...
{: .fs-6 .fw-300 }

---

{: .text-green-200}
## Installation

```bash
go get -u github.com/wesovilabs/goa
```

{: .text-green-200}
## But... how does goa work?

Goa could be understood like a go wrapper that takes the below steps:
- It enriches our code with the registered aspects
- It delegates the work to `go`.

{: .text-green-200}
## Goa in action

`[env_vars] goa [goa_flags] go_command [go_flags]`

{: .text-yellow-300}
### Environment variables (env_vars)
Environment variables can be provided. They will be propagated to the go command too.


{: .text-yellow-300}
### Flags (goa_flags)
Goa provides some flags that we can use to customize default behavior 

|  Flag                         |Default                   |Description                                                     |
|:------------------------------|:-------------------------|:---------------------------------------------------------------|:
|`--project <projectname>`      | module name in go.mod    | only required if you don't use go.mod                          |
|`--output <directory>`         | a temporal directory     | directory used by goa to copy de generated code                |
|`--path <directory>`           | current directory        | path to your project                                           |
|`--verbose`                    | false                    | enable verbose mode                                       |
|`--work`                       | false                    | print the name of the temporary work directory and do not delete it when exiting |


{: .text-yellow-300}
### Commands (go_command)
So far, allowed commands are: `build`, `run` and `generate`.

{: .text-yellow-300}
### Go glags (go_flags)

{: .text-green-200}
## Goa in action by examples

Let's suppose we have a code organization like the one shown below

+ cmd
    + app
        - main.go
+ internal
    - files.go  
    + helper
        - strings.go
+ constants
    - app.go
    
The next command could be fine to build the application

```bash
>> CGO_ENABLED=0 goa --project myapp build -ldflags "-X constants.Version=0.0.1" cmd/app/main.go
```

However when we work with go modules command we con ignore th flag `--project`.

+ cmd
    + app
        - main.go
+ internal
    - files.go  
    + helper
        - strings.go
+ constants
    - app.go
- go.mod

```bash
>> CGO_ENABLED=0 goa build -ldflags "-X constants.Version=0.0.1" cmd/app/main.go
```
