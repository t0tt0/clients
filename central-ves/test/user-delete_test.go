package tests

import (
	"github.com/Myriad-Dreamin/go-ves/central-ves/test/tester"
	types2 "github.com/Myriad-Dreamin/go-ves/types"
	"net/http"
	"reflect"
	"strconv"
)

func testUserDelete(t *tester.TesterContext) {
	id := reflect.ValueOf(srv.Get(normalUserIdKey)).Convert(intType).Interface().(int)
	t.Delete("/v1/user/" + strconv.Itoa(id))
	t.AssertNoError(false)
	resp := t.Get("/v1/user/" + strconv.Itoa(id))
	NotFound := t.FetchError(resp)
	t.Equal(NotFound.RespCode, http.StatusOK, "must result in 200")
	t.Equal(NotFound.Code, types2.CodeNotFound, "not be the ErrNotFound")
}
