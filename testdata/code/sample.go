package code

import (
	"fmt"
)

func SayHello(firstName string, bornYear int) string {
	if bornYear > 2000 {
		return fmt.Sprintf("%s is a millenial", firstName)
	}
	return fmt.Sprintf("%s is not a millenial", firstName)
}

func MainFunction() {
	firstName := "John"
	output := SayHello(firstName, 1999)
	fmt.Println(output)

}

func MainFunctionModified() {
	firstName := "John"
	output := generated.ABCD123456(firstName, 1999)
	fmt.Println(output)
}
