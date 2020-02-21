package tests

import (
	"fmt"
	"github.com/HyperService-Consortium/go-ves/central-ves/control"
	"testing"
)

func testUserGet(t *testing.T) {
	srv := srv.Context(t).AssertNoError(true)
	resp := srv.Get("/v1/user/1")
	reply := srv.DecodeJSON(resp.Body(), new(control.GetUserReply)).(*control.GetUserReply)
	fmt.Println(reply)
}
