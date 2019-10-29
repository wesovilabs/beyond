package aspects

import (
	"fmt"
	"github.com/wesovilabs/goa/api"
	"github.com/wesovilabs/goa/api/context"
)

type NormalizeID struct {

}

func (a *NormalizeID) Returning(ctx *context.Context){
	id:=ctx.Out().Get("result0").(string)
	ctx.Out().Set("result0",fmt.Sprintf("ID:%s",id))
}

func NewNormalizeID() api.Returning{
	return &NormalizeID{

	}
}
