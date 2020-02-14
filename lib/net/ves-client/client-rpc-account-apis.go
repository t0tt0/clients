package vesclient

import (
	"github.com/HyperService-Consortium/go-uip/uiptypes"
	ginhelper "github.com/Myriad-Dreamin/go-ves/lib/gin-helper"
	"github.com/Myriad-Dreamin/go-ves/lib/miris"
	"github.com/Myriad-Dreamin/go-ves/types"
	"github.com/Myriad-Dreamin/minimum-lib/controller"
	"github.com/kataras/iris"
	"net/http"
)

func (vc *VesClient) buildAccountRPCApis(p iris.Party) {
	p.Get("account-list", vc.IrisListAccount)
	id := p.Party("account-list/{aid}")
	p.Post("account", miris.ToIrisHandler(vc.IrisPostAccount))
	id.Delete("", miris.ToIrisHandler(vc.IrisDeleteAccount))
	id.Put("", miris.ToIrisHandler(vc.IrisPutAccount))
}

type PostAccountRequest struct {
	ChainType uiptypes.ChainTypeUnderlyingType `json:"chain_type" form:"chain_type"`
	ChainID   uiptypes.ChainIDUnderlyingType   `json:"chain_id" form:"chain_id"`
	Address   []byte                           `json:"address" form:"address"`
}

func (vc *VesClient) IrisPostAccount(c controller.MContext) {
	var req PostAccountRequest
	if !ginhelper.BindRequest(c, &req) {
		return
	}

	var account Account
	account.ChainType = req.ChainType
	account.ChainID = req.ChainID
	account.Address = encodeAddress(req.Address)

	if _, err := vc.db.Create(&account); err != nil {
		c.JSON(http.StatusOK, errorSerializer(types.CodeInsertError, err))
		return
	}

	c.JSON(http.StatusOK, ginhelper.ResponseOK)
	return
}

func (vc *VesClient) IrisDeleteAccount(c controller.MContext) {
	id, ok := ginhelper.ParseUint(c, "aid")
	if ok {
		c.JSON(http.StatusOK, iris.Map{
			"code": types.CodeOK,
			"id":   id,
		})
	}
	return
}

func (vc *VesClient) IrisPutAccount(c controller.MContext) {

	c.JSON(http.StatusOK, iris.Map{
		"code": types.CodeToDo,
		"todo": 1,
	})
	return
}
