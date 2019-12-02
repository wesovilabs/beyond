---
layout: page
title: Beyond
permalink: /beyond/
nav_order: 6
---

{: .text-green-300}
# >> beyond ...
{: .fs-9 }

{: .text-green-200}
The magic wand that will make it happen...
{: .fs-6 .fw-300 }

---

{: .text-green-200}
## Installation

```bash
go get -u github.com/wesovilabs/beyond
```

{: .text-green-200}
## But... how does beyond work?

Beyond could be understood like a go wrapper that takes the below steps:
- It enriches our code with the registered aspects
- It delegates the work to `go`.

{: .text-green-200}
## Beyond in action

`[env_vars] beyond [beyond_flags] go_command [go_flags]`

{: .text-yellow-300}
### Environment variables (env_vars)
Environment variables can be provided. They will be propagated to the go command too.


{: .text-yellow-300}
### Flags (beyond_flags)
Beyond provides some flags that we can use to customize default behavior 

|  Flag                         |Default                   |Description                                                     |
|:------------------------------|:-------------------------|:---------------------------------------------------------------|:
|`--project <projectname>`      | module name in go.mod    | only required if you don't use go.mod                          |
|`--output <directory>`         | a temporal directory     | directory used by beyond to copy de generated code                |
|`--path <directory>`           | current directory        | path to your project                                           |
|`--verbose`                    | false                    | enable verbose mode                                       |
|`--work`                       | false                    | print the name of the temporary work directory and do not delete it when exiting |
|`--config`                     | beyond.toml              | It loads the beyond configuration from the given toml file |


{: .text-yellow-300}
### Commands (go_command)
So far, allowed commands are: `build`, `run` and `generate`.

{: .text-yellow-300}
### Go glags (go_flags)

{: .text-green-200}
## Beyond in action by examples

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
>> CGO_ENABLED=0 beyond --project myapp build -ldflags "-X constants.Version=0.0.1" cmd/app/main.go
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
>> CGO_ENABLED=0 beyond build -ldflags "-X constants.Version=0.0.1" cmd/app/main.go
```

##  beyond.toml

As it was mentioned above, we can provide a configuration file to avoid pass arguments to the beyond command.

This file must have the below structure

```toml
project="github.com/wesovilabs/beyond-examples/settings"
output="generated"
verbose=true
work=true
excludes=[
    "go.sum",
    "vendor",
    ".git"
]
```

- project: The project name
- output: The directory in which our code will be generated.
- verbose: If true, we run beyond with verbose mode, false in other case
- work: If true,  print the name of the temporary work directory and do not delete it when exiting
- excludes: We can list files or directories which don't need to be taken in account by beyond.

To understand how the beyond.toml works, let's clone the beyond-examples repository
```bash
>> git clone https://github.com/wesovilabs/beyond-examples.git
>> cd settings
```
To run beyond with verbosity and leave the generated files in directory `generated` we should do the following

```bash
beyond --verbose --work --output generated run cmd/main.go
```

We can avoid passing always the flags if we create the following file to the root.

**beyond.toml**
```toml
project="github.com/wesovilabs/beyond-examples/settings"
output="generated"
verbose=true
work=true
excludes=[
    "go.sum",
    "vendor",
    ".git"
]
```

In case of we want to use a configuration file with a different name we can rename
beyond.toml to beyond-dev.toml and execute the below command

```bash
beyond --config beyond-dev.toml run cmd/main.go
```
