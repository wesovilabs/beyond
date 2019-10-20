package aspect

import (
	"fmt"
	"github.com/wesovilabs/goa/context"
	"strings"
	"time"
)

func LoggerAround(pkg string, function string, in *context.Input) {
	t := time.Now()
	args := []string{}
	if in.Len() > 0 {
		if in != nil {
			for _, arg := range in.List() {
				args = append(args, fmt.Sprintf("%s:%v=%#v ", arg.Name(), arg.Kind(), arg.Value()))
			}
		}
	}
	fmt.Printf("    [%v] %s.%s with %s\n", t.Format("02/01/2006 15:04:05.999"), pkg, function, strings.Join(args, ","))
}
