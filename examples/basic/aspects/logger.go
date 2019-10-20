package aspects

import (
	"fmt"
	"github.com/wesovilabs/goa/context"
	"strings"
	"time"
)




func LogAspect(ctx *context.Ctx) {
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
