package tests

import (
	"github.com/Myriad-Dreamin/go-ves/central-ves/control"
	"github.com/Myriad-Dreamin/go-ves/central-ves/test/tester"
)

func testUserRegisterLogin(t *tester.TesterContext) {
	var (
		name  = "chan tan"
		pswd  = normalUserPassword
	)
	resp := t.Post("/v1/user", control.RegisterRequest{
		Name:     name,
		Password: pswd,
	})
	id := t.DecodeJSON(resp.Body(),
		new(control.RegisterReply)).(*control.RegisterReply).Id
	resp = t.Post("/v1/login", control.LoginRequest{
		Id:       id,
		Password: pswd,
	})
	resp = t.Post("/v1/login", control.LoginRequest{
		Name: name,
		Password: pswd,
	})

	srv.Set(normalUserIdKey, id)
}
