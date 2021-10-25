package vesclient

import (
	"github.com/HyperService-Consortium/go-ves/types"
	"github.com/kataras/iris"
)

func (vc *VesClient) IrisListAccount(ctx iris.Context) {
	//BindQuery(ctx, )
	accounts, err := vc.db.QueryChain().Query()
	if err != nil {
		vc.ContextJSON(ctx, errorSerializer(types.CodeSelectError, err))
		return
	}
	vc.ContextJSON(ctx, iris.Map{
		"code":   CodeOk,
		"result": accounts,
	})
	return
}
