package tests

import (
	"fmt"
	"github.com/HyperService-Consortium/go-ves/central-ves/control"
	"github.com/HyperService-Consortium/go-ves/central-ves/test/tester"
	"testing"
)

func TestChainID(t *testing.T) {
	_ = t.Run("Post", srv.HandleTestWithOutError(testChainIDPost)) &&
		t.Run("Put", srv.HandleTestWithOutError(testChainIDPut)) &&
		t.Run("Get", srv.HandleTestWithOutError(testChainIDGet)) &&
		t.Run("Delete", srv.HandleTestWithOutError(testChainIDDelete)) // &&
	//t.Run("Put", srv.HandleTestWithOutError(testChainID))

}

func testChainIDPut(t *tester.TesterContext) {

}

func testChainIDGet(t *tester.TesterContext) {
	resp := t.Get("/v1/chain_info/1")
	fmt.Println(resp)
}

func testChainIDDelete(t *tester.TesterContext) {
	//
}

func testChainIDPost(t *tester.TesterContext) {
	//_ = t.Post("/v1/problem/1/submission", control.PostSubmissionRequest{
	//                Language:    1,
	//                Information: "1",
	//                Shared:      0,
	//                Code:        "123",
	//        })
	resp := t.Post("/v1/chain_info", control.PostChainInfoRequest{
		UserId:  1,
		Address: "A=",
		ChainId: 3,
	})
	var x control.PostChainInfoReply
	t.HandlerError0(resp.JSON(&x))
	fmt.Println(x)
}
