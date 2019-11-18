---
layout: default
title: Joinpoint
nav_order: 3
---


{: .text-green-300}
# Joinpoint
{: .fs-9 }

{: .text-green-200}
Where the magic happens...  Keep open your eyes!
{: .fs-6 .fw-300 }

---


{: .text-green-300}
## Introduction

It's important to keep in mind that Goa provides a mechanism to intercept both functions and methods invocations.

---

{: .text-green-300}
## Syntax

Goa interprets  the provided expressions in order to decide which functions must be intercepted by the advices.

`<package>.<type>?.<function>(<params)<results>`

* `<type>.` is only required when the advice needs to intercept a method instead of a function. 

{: .text-yellow-300}
### A brief cheat sheet 

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
| `database.set*(*model.Person)...`        | Any function whose name `starts with set` in package `database`, with 1 params of type `*model.Person`, and 0 or N results |

---

{: .text-green-300}
## Let’s practice!

{: .text-yellow-300}
### Prerequisites
 
Let's check that our environment is ready to follow the tutorial!
  
- Install goa tool 
```bash
>> go get github.com/wesovilabs/goa
```

- Clone the [goa-examples repository](https://github.com/wesovilabs/goa-examples.git)
```bash
>> git clone https://github.com/wesovilabs/goa-examples.git
>> cd goa-examples
>> git checkout feature/joinpoints
 ```

{: .text-yellow-300}
### Test application

The application provices a Rest API with the below methods:

- Create Employee
- Get employee
- Delete employee
- List employees

The application contains a Before advice that prints the functions/methods invocations.

{: .text-green-300}
*Intercept handler CreateEmployee*

- Replace the regExp in function `Goa` in file `main.go` by on these:
```go
handler.CreateEmployee(...)...
*.CreateEmployee(...)...
handler.CreateEmployee(net/http.ResponseWriter,*net/http.Request,github.com/wesovilabs/goaexamples/storage.Database)
handler.CreateEmployee(...,*net/http.Request,...)
handler.Create*(...,*net/http.Request,...)
```
  
- Run the server application
```bash
>> goa run main.go
```
- Run the client 
```go
>> go run test/main.go
```  
- Check the server stdout 
```bash
>> goa run main.go
Launching server on localhost:8000
handler.CreateEmployee
```

{: .text-green-300}
*Intercept handlers*

- Replace the regExp in function `Goa` in file `main.go` by on these:
```go
handler.*(...)...
*.*Employee(...)...
handler.*Employee(net/http.ResponseWriter,*net/http.Request,github.com/wesovilabs/goaexamples/storage.Database)
handler.*(...,*net/http.Request,...)
*.*(net/http.ResponseWriter,*net/http.Request,...)
```
  
- Run the server application
```bash
>> goa run main.go
```
- Run the client 
```go
>> go run test/main.go
```  
- Check the server stdout 
```bash
>> goa run main.go
Launching server on localhost:8000
handler.CreateEmployee
handler.GetEmployee
handler.ListEmployees
handler.DeleteEmployee
```

{: .text-green-300}
*Intercept SaveEmployee method for type memDBClient*

- Replace the regExp in function `Goa` in file `main.go` by on these:
```go
storage.*memDBClient.SaveEmployee(*github.com/wesovilabs/goaexamples/model.Employee)error
storage.*memDBClient.Save*(*github.com/wesovilabs/goa/model.Employee)error
storage.*memDBClient.Save*(...)...
*.*.SaveEmployee(...)...
*.*memDBClient.Save*(...)...
```
  
- Run the server application
```bash
>> goa run main.go
```
- Run the client 
```go
>> go run test/main.go
```  
- Check the server stdout 
```bash
>> goa run main.go
Launching server on localhost:8000
storage.*storage.memDBClient.SaveEmployee
```


{: .text-green-300}
*Intercept  function RespondWithJSON in internal package handler/internal*

- Replace the regExp in function `Goa` in file `main.go` by on these:
```go
handler/internal.RespondWithJSON(net/http.ResponseWriter,int,interface{})
*/internal.RespondWithJSON(net/http.ResponseWriter,int,interface{})
handler/*.RespondWithJSON(net/http.ResponseWriter,int,interface{})
*.RespondWithJSON(net/http.ResponseWriter,...)
*.RespondWithJSON(...,int,...)
```
  
- Run the server application
```bash
>> goa run main.go
```
- Run the client 
```go
>> go run test/main.go
```  
- Check the server stdout 
```bash
>> goa run main.go
Launching server on localhost:8000
internal.RespondWithJSON
internal.RespondWithJSON
internal.RespondWithJSON
internal.RespondWithJSON
```

{: .text-green-300}
## Challenge

Find valid expressions for intercepting...

- All the memDBClient methods
- Function `RandomString` in package helper
- Function `RespondWithError`. By the way, you will need to force this error. In file `test/main.go`,  set an empty email. 
```go
res,_ := api.CreateEmployee(&model.Employee{
   Email:    "",
   Fullname: "John Doe",
})
```


If you found any problem to resolve this challenge, don't hesitate to drop me an email at `ivan.corrales.solera@gmail.com` and I will
be happy to give you some help.


If you enjoyed this article, I would really appreciate if you shared it with your networks


<div class="socialme">
    <ul>
        <li class="twitter">
            <a href="https://twitter.com/intent/tweet?via={{site.data.social.twitter.username}}&url={{ site.data.social.twitter.url | uri_escape}}&text={{ site.data.social.twitter.message | uri_escape}}" target="_blank">
                {% include social/twitter.svg %}
            </a>
        </li>
    </ul>
</div>
