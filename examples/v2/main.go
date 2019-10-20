package main

import (
	"fmt"
	"github.com/wesovilabs/goa/api"
	"github.com/wesovilabs/goa/context"
	"strings"
	"time"
)

func LoggerAround(ctx *context.AroundCtx) {
	t := time.Now()
	args := []string{}
	if ctx.In().Len() > 0 {
		if ctx.In() != nil {
			for _, arg := range ctx.In().List() {
				args = append(args, fmt.Sprintf("%s:%v=%#v ", arg.Name(), arg.Kind(), arg.Value()))
			}
		}
	}
	fmt.Printf("    [%v] %s.%s with %s\n", t.Format("02/01/2006 15:04:05.999"), ctx.Pkg(), ctx.Name(), strings.Join(args, ","))
}


func Goa() *api.Goa {
	return api.Init().
		WithAround(`main\.test1\(string\)`, LoggerAround)
}

func main() {
	fmt.Println("\n[test]")
	test1("pepe")
	fmt.Println("\n[test2]")
	test2(2)
	fmt.Println("\n[test3]")
	test3("John", "Doe")
	fmt.Println("----------")
	test5("John", fmt.Println)
}

func test1(name string) {
	fmt.Printf("    name is %v\n", name)
}

func test2(value int) {
	fmt.Printf("    value is %v\n", value)
}

func test3(name string, surname string) {
	fmt.Printf("    %s %s\n", name, surname)
}

func test5(value string, fn func(...interface{}) (int, error)) {
	fn(value)
}

