package main

import (
	"fmt"
	"regexp"
)

const pkgChars = `[a-zA-Z0-9_\-*/]+`
const modChars = `[a-zA-Z0-9_\-*]+\.`
const funcChars = `[a-zA-Z0-9_\-*]+`
const inChars = `[a-zA-Z0-9\[\]_\-*\,]*`
const outChars = `[a-zA-Z0-9_\[\]\-*\,]*|\([a-zA-Z0-9_\-*\,]*\)`

func exprStr() string {
	out := `^(` + `?P<pkg>` + pkgChars + `)\.`
	out += `(` + `?P<mod>` + modChars + `)?`
	out += `(` + `?P<func>` + funcChars + `)`
	out += `\((` + `?P<in>` + inChars + `)\)`
	out += `[ ]*`
	out += `(` + `?P<out>` + outChars + `)?$`
	return out
}

var expRegExp = regexp.MustCompile(exprStr())

func processExprStr(text string) {
	fmt.Println()
	items := expRegExp.FindStringSubmatch(text)
	fmt.Println(len(items))
	if len(items) == 0 {
		fmt.Println("ERROR")
		fmt.Println(text)
		return
	}
	fmt.Println("-----")
	fmt.Printf("str : %s\n", text)
	fmt.Printf("pkg : %s\n", items[1])
	fmt.Printf("mod : %s\n", items[2])
	fmt.Printf("func: %s\n", items[3])
	fmt.Printf("in  : %s\n", items[4])
	fmt.Printf("out : %s\n", items[5])
}

var texts = []string{
	"test.demo.vyd(d)string",
	"test.demo.vyd(d)(as,asd)",
	"a.build()",
	"test/test/test.build()",
	"a.b()",
	"a.b.c()",
	"a/k.d.buld()",
	"*.*()",
	"a.b(string,demo)*",
	"goa/demo.build(string,*) (*,int)",
	"goa/demo.object.build(string,*) (*,int)",
}

func main() {
	fmt.Println(exprStr())
	for _, text := range texts {
		processExprStr(text)

	}
}
