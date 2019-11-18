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

goa could be understood like a go wrapper.. that enriches our code with the registered aspects and then
If delegates the work to go. 

{: .text-green-200}
## Goa in action

`[env_vars] goa [goa_flags] go_command [go_flags]`


{: .text-yellow-300}
### Flags

Goa provides some flags that we can use to customize default behavior 

|  Flag                                         |Default         |  Description              |
|:-----------------------------------------------|:----------------------------------------------|:-------------------------|:
|`--project <projectname>`      | module name in go.mod          |    only required if you don't use go.mod |
|`--verbose`                    | false                          | It displays the goa logs         |
|`--output <directory>`         | a temporal directory           | goa clone your path and generate code in it |
|`--path <directory>`           | current directory              | path where your project is |
|`--package <directory>`        | .                              | relative path where your main package is |
|`--work`                       | false                          | print the name of the temporary work directory and do not delete it when exiting |

{: .text-yellow-300}
### Commands

Any command provided by go can be used.

{: .text-yellow-300}
### Environment variables

Environment variables can be provided and they will passed to the go command too.

{: .text-yellow-300}
### Examples

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
    
The next command could be fine for building our application

```bash
CGO_ENABLED=0 goa --project myapp --package cmd/app \
    build -ldflags "-X constants.Version=0.0.1" cmd/app/main.go
```

However when we work with go modules command we con ignore flag project

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
CGO_ENABLED=0 goa --package cmd/app \
    build -ldflags "-X constants.Version=0.0.1" cmd/app/main.go
```

On the other hand,  if the main.go is in the root of the project, the command can be even more straightforward

- main.go
+ internal
    - files.go  
    + helper
        - strings.go
+ constants
    - app.go
- go.mod

```bash
CGO_ENABLED=0 goa build -ldflags "-X constants.Version=0.0.1" main.go
```
