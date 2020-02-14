package vesclient

import (
	"fmt"
	ginhelper "github.com/Myriad-Dreamin/go-ves/lib/gin-helper"
	"github.com/Myriad-Dreamin/go-ves/types"
	"github.com/Myriad-Dreamin/minimum-lib/controller"
	"net/http"
	"strings"
)

type PostSessionRequest struct {
	Dependencies []string `json:"dependencies" form:"dependencies"`
	Intents      []string `json:"intents" form:"intents"`
}

func (vc *VesClient) IrisPostSession(c controller.MContext) {
	var req PostSessionRequest
	if !ginhelper.BindRequest(c, &req) {
		return
	}

	var session Session

	session.Intents = fmt.Sprintf("[%s]", strings.Join(req.Intents, ","))
	session.Dependencies = fmt.Sprintf("[%s]", strings.Join(req.Dependencies, ","))

	session.NSBHost = vc.nsbHost

	host, err := vc.getVESHost()
	if err != nil {
		c.JSON(http.StatusOK, errorSerializer(types.CodeGetVESHostError, err))
		return
	}
	session.VESHost = host

	if _, err := vc.sessionDB.Create(&session); err != nil {
		c.JSON(http.StatusOK, errorSerializer(types.CodeInsertError, err))
		return
	}

	if sessionID, err := vc.SendOpIntentsByStrings(
		host, req.Intents, req.Dependencies); err != nil {
		c.JSON(http.StatusOK, errorSerializer(types.CodeInsertError, err))
		return
	} else {
		session.ISCAddress = encodeAddress(sessionID)
	}

	c.JSON(http.StatusOK, ginhelper.ResponseOK)
	return
}
