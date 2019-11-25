---
layout: default
title: Joinpoint
nav_order: 3
---


{: .text-green-300}
# Joinpoint Expressions
{: .fs-9 }

{: .text-green-200}
Where the magic happens...  Keep open your eyes!
{: .fs-6 .fw-300 }

---

{: .text-green-300}
## Syntax

It’s important to keep in mind that Goa provides a mechanism to intercept both functions and methods invocations.

Goa interprets  the provided expressions to decide which functions must be intercepted by the advices. The expressions have the 
following format.

`<package>.<type>?.<function>(<params)<results>`

*`<type>.` is only required when the advice needs to intercept a method instead of a function.* 

{: .text-yellow-300}
### A brief cheat sheet 

The table contains some examples, that could help us to get a better understanding.

| Expression                               | Intercepted               |
|:-----------------------------------------|:--------------------------|:
| `*.*(...)...`                            | Any function invocation with 0 or N params and 0 or N results |
| `*.*.*(...)...`                          | Any method invocation with 0 or N params and 0 or N results |
| `model.*(...)...`                      | Any function invocation, in package `model`,  with 0 or N params and 0 or N results |
| `handlers/employee.*(...)...`          | Any function invocation, in package `handlers/employee`,  with 0 or N params and 0 or N results |
| `model.*.*(...)...`                      | Any method invocation, in package `model`,  with 0 or N params and 0 or N results |
| `model.person.*(...)...`                 | Any method invocation, for type `person` in package `model`,  with 0 or N params and 0 or N results |
| `database.*(string)...`                  | Any function in package `database`, with 1 param of type string and 0 or N results |
| `database.*(string,*int32)...`           | Any function in package `database`, with 2 params of types string and *int32, and 0 or N results |
| `database.*(string,*)...`                | Any function in package `database`, with 2 params of types string and the second param of any type, and 0 or N results |
| `database.*(string,...)...`              | Any function in package `database`, with 1 string param and 1 or more params of any type, and 0 or N results |
| `database.*(string,...)func()string`     | Any function in package `database`, with 1 string param and 1 or more params of any type, and 1 result whose type is `func()string`|
| `database.set*(*model.Person)`        | Any function whose name `starts with set` in package `database`, with 1 params of type `*model.Person`, and 0 results |

---

{: .text-yellow-300}
### Let's practice

Let's check that our environment is ready to follow the tutorial!
 
- Install goa tool & clone the goaexamples repository
```bash
>> go get github.com/wesovilabs/goa
>> git clone https://github.com/wesovilabs/goaexamples.git
>> cd joinpoints
```

-  The application provides a Rest API to interact with employee resources. A test purposed advice is 
registered in file [cmd/main.go](https://github.com/wesovilabs/goaexamples/blob/master/cmd/joinpoints/main.go#L14).

To test the joinpoints expreessions we need to launch the server with command `goa run cmd/main.go` and then, 
run `go test/main.go` to make some requests to the server. 

Server main is found in file [cmd/main.go](https://github.com/wesovilabs/goaexamples/blob/master/joinpoints/cmd/main.go#L19) 
and client main in [test/main.go](https://github.com/wesovilabs/goaexamples/blob/master/joinpoints/test/main.go)


{: .text-green-300}
**Intercepting function [CreateEmployee](https://github.com/wesovilabs/goaexamples/blob/master/joinpoints/handler/employee.go#L12)**

Any of these expressions would be valid to intercept the above function.
- `handler.CreateEmployee(...)...`
- `*.CreateEmployee(...)...`
- `handler.CreateEmployee(net/http.ResponseWriter,*net/http.Request,github.com/wesovilabs/goaexamples/storage.Database)`
- `handler.CreateEmployee(...,*net/http.Request,...)`
- `handler.Create*(...,*net/http.Request,...)`

To check the expressions we just need to modify the registered expression in [cmd/main.go](https://github.com/wesovilabs/goaexamples/blob/master/joinpoints/cmd/joinpoints/main.go#L19)
```go
func Goa() *api.Goa {
  return api.New().
  WithBefore(advice.NewSimpleTracingAdvice, `handler.Create*(...,*net/http.Request,...)`)
}
```
and then, run the server and the client.
```bash
>> goa run main.go
Launching server on localhost:8000
handler.CreateEmployee
```
```bash
>> go run test/main.go
```

{: .text-green-300}
**Intercepting any function in package [handler](https://github.com/wesovilabs/goaexamples/blob/master/joinpoints/handler/employee.go)**

Any of these expressions would be valid to intercept the above function.
- `handler.*(...)...`
- `handler.*(...,*net/http.Request,...)`
- `*.*(net/http.ResponseWriter,*net/http.Request,...)`

To check the expressions we just need to modify the registered expression in [cmd/main.go](https://github.com/wesovilabs/goaexamples/blob/master/joinpoints/cmd/main.go#L19)
```go
func Goa() *api.Goa {
  return api.New().
  WithBefore(advice.NewSimpleTracingAdvice, `handler.*(...,*net/http.Request,...)`)
}
```
and then, run the server and the client.
```bash
>> goa run main.go
Launching server on localhost:8000
handler.CreateEmployee
handler.GetEmployee
handler.ListEmployees
handler.DeleteEmployee
```
```bash
>> go run test/main.go
```

Why don't these other expressions print the same output?
- `*.*Employee(...)...`
- `handler.*Employee(net/http.ResponseWriter,*net/http.Request,github.com/wesovilabs/goaexamples/joinpoints/storage.Database)`
  

{: .text-green-300}
**Intercepting method SaveEmployee of type [memDBClient](https://github.com/wesovilabs/goaexamples/blob/master/joinpoints/storage/mem.go#L20)**

Any of these expressions would be valid to intercept the above function.
- `storage.*memDBClient.SaveEmployee(*github.com/wesovilabs/goaexamples/model.Employee)error`
- `storage.*memDBClient.Save*(*github.com/wesovilabs/goaexamples/model.Employee)error`
- `storage.*memDBClient.Save*(...)...`
- `*.*.SaveEmployee(...)...`
- `*.*memDBClient.Save*(...)...`
  
To check the expressions we just need to modify the registered expression in [cmd/joinpoints/main.go](https://github.com/wesovilabs/goaexamples/blob/master/joinpoints/cmd/main.go#L19)
```go
func Goa() *api.Goa {
  return api.New().
  WithBefore(advice.NewSimpleTracingAdvice, `handler.*(...,*net/http.Request,...)`)
}
```
and then, run the server and the client.
```bash
>> goa run main.go
Launching server on localhost:8000
storage.*storage.memDBClient.SaveEmployee
```
```bash
>> go run test/main.go
```

{: .text-green-300}
**Intercepting function RespondWithJSON in package [handler/internal](https://github.com/wesovilabs/goaexamples/blob/master/joinpoints/handler/internal/helper.go#L14)**

Any of these expressions would be valid to intercept the above function.
- `handler/internal.RespondWithJSON(net/http.ResponseWriter,int,interface{})`
- `*/internal.RespondWithJSON(net/http.ResponseWriter,int,interface{})`
- `handler/*.RespondWithJSON(net/http.ResponseWriter,int,interface{})`
- `*.RespondWithJSON(net/http.ResponseWriter,...)`
- `*.RespondWithJSON(...,int,...)`

To check the expressions we just need to modify the registered expression in [cmd/main.go](https://github.com/wesovilabs/goaexamples/blob/master/joinpoints/cmd/main.go#L19)
```go
func Goa() *api.Goa {
  return api.New().
  WithBefore(advice.NewSimpleTracingAdvice, `*.RespondWithJSON(...,int,...)`)
}
```
and then, run the server and the client.
```bash
>> goa run main.go
Launching server on localhost:8000
internal.RespondWithJSON
internal.RespondWithJSON
internal.RespondWithJSON
internal.RespondWithJSON
```
```bash
>> go run test/main.go
```

{: .text-green-300}
## Challenge

Find valid expressions that intercept the following:

- All the invocations to memDBClient methods
- Function `RandomString` in package helper
- Function `RespondWithError`. To check it, you will need to force an error. You can do it in file [test/main.go](https://github.com/wesovilabs/goaexamples/blob/master/joinpoints/test/main.go#L22) by making the below change. 
```go
res,_ := api.CreateEmployee(&model.Employee{
   Email:    "",
   Fullname: "John Doe",
})
```

If you found any problem to resolve this challenge, don't hesitate to drop me an email at `ivan.corrales.solera@gmail.com` and I will
be happy to give you some help.

---

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
