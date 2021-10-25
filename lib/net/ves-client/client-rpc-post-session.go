package vesclient

import (
	"fmt"
	ginhelper "github.com/HyperService-Consortium/go-ves/lib/backend/gin-helper"
	"github.com/HyperService-Consortium/go-ves/lib/backend/serial"
	"github.com/HyperService-Consortium/go-ves/types"
	"github.com/Myriad-Dreamin/minimum-lib/controller"
	"net/http"
	"strings"
)

type PostSessionRequest struct {
	Dependencies []string `json:"dependencies" form:"dependencies"`
	Intents      []string `json:"intents" form:"intents"`
}

func (vc *VesClient) createSession(c controller.MContext, session *Session) {
	host, err := vc.getVESHost()
	if err != nil {
		c.JSON(http.StatusOK, errorSerializer(types.CodeGetVESHostError, err))
		return
	}
	session.NSBHost = vc.nsbHost
	session.VESHost = host
	return
}

func (vc *VesClient) insertSession(c controller.MContext, session *Session) {
	if _, err := vc.sessionDB.Create(session); err != nil {
		c.JSON(http.StatusOK, errorSerializer(types.CodeInsertError, err))
	} else {
		c.JSON(http.StatusOK, ginhelper.ResponseOK)
	}
	return
}

func (vc *VesClient) IrisPostSessionR(c controller.MContext) {
	var session Session
	vc.createSession(c, &session)
	if c.IsAborted() {
		return
	}
	b, err := c.GetRawData()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusOK, &serial.ErrorSerializer{
			Code: types.CodeInvalidParameters,
			Err:  err.Error(),
		})
		return
	}

	if sessionID, err := vc.SendOpIntentsR(
		session.VESHost, b); err != nil {
		//todo: CodeSendOpIntentError
		c.JSON(http.StatusOK, errorSerializer(types.CodeInsertError, err))
		return
	} else {
		session.ISCAddress = encodeAddress(sessionID)
	}

	vc.insertSession(c, &session)
}

func (vc *VesClient) IrisPostSession(c controller.MContext) {
	var session Session
	vc.createSession(c, &session)
	if c.IsAborted() {
		return
	}

	var req PostSessionRequest
	if !ginhelper.BindRequest(c, &req) {
		return
	}

	session.Intents = fmt.Sprintf("[%s]", strings.Join(req.Intents, ","))
	session.Dependencies = fmt.Sprintf("[%s]", strings.Join(req.Dependencies, ","))

	if sessionID, err := vc.SendOpIntentsByStrings(
		session.VESHost, req.Intents, req.Dependencies); err != nil {
		c.JSON(http.StatusOK, errorSerializer(types.CodeInsertError, err))
		return
	} else {
		session.ISCAddress = encodeAddress(sessionID)
	}

	vc.insertSession(c, &session)
}
