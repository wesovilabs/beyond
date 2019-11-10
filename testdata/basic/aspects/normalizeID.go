package aspects

import (
	"fmt"
	"github.com/wesovilabs/goa/api"
	"github.com/wesovilabs/goa/api/context"
)

type NormalizeID struct {

}

func (a *NormalizeID) Returning(ctx *context.GoaContext){
	id:=ctx.Results().Get("result0").(string)
	ctx.Results().Set("result0",fmt.Sprintf("ID:%s",id))
}

func NewNormalizeID() api.Returning{
	return &NormalizeID{

	}
}
