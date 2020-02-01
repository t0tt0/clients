package main

import (
	"github.com/Myriad-Dreamin/artisan"
	"github.com/Myriad-Dreamin/go-ves/central-ves/model"
)

type UserCategories struct {
	artisan.VirtualService
	List           artisan.Category
	Login          artisan.Category
	Register       artisan.Category
	ChangePassword artisan.Category
	Inspect        artisan.Category
	IdGroup        artisan.Category
}

func DescribeUserService(base string) artisan.ProposingService {
	var userModel = new(model.User)
	var vUserModel model.User
	svc := &UserCategories{
		List: artisan.Ink().
			Path("user-list").
			Method(artisan.POST, "ListUsers",
				artisan.QT("ListUsersRequest", model.Filter{}),
				artisan.Reply(
					codeField,
					artisan.ArrayParam(artisan.Param("users", artisan.Object(
						"ListUserReply",
						artisan.SPsC(&vUserModel.ID, &vUserModel.LastLogin),
					))),
				),
			),
		Login: artisan.Ink().
			Path("login").
			Method(artisan.POST, "Login",
				artisan.Request(
					artisan.SPsC(&userModel.ID, &userModel.Name),
					artisan.Param("password", artisan.String, required),
				),
				artisan.Reply(
					codeField,
					artisan.SPsC(&userModel.ID, &userModel.Name),
					artisan.Param("identity", artisan.Strings),
					artisan.Param("token", artisan.String),
					artisan.Param("refresh_token", artisan.String),
				),
			),
		Register: artisan.Ink().
			Path("register").
			Method(artisan.POST, "Register",
				artisan.Request(
					artisan.SPs(artisan.C(&userModel.Name), required),
					artisan.Param("password", artisan.String, required),
				),
				artisan.Reply(
					codeField,
					artisan.Param("id", &userModel.ID)),
			),
		ChangePassword: artisan.Ink().
			Path("user/:id/password").
			Method(artisan.PUT, "ChangePassword",
				artisan.Request(
					artisan.Param("old_password", artisan.String, required),
					artisan.Param("new_password", artisan.String, required),
				),
			),
		Inspect: artisan.Ink().Path("user/:id/inspect").
			Method(artisan.GET, "InspectUser",
				artisan.Reply(
					codeField,
					artisan.Param("user", &userModel),
				),
			),
		IdGroup: artisan.Ink().
			Path("user/:id").
			Method(artisan.GET, "GetUser",
				artisan.Reply(
					codeField,
					artisan.SPsC(&userModel.ID, &userModel.LastLogin),
				)).
			Method(artisan.PUT, "PutUser",
				artisan.Request(
					//artisan.Param("phone", &userModel.Phone),
				)).
			Method(artisan.DELETE, "Delete"),
	}
	svc.Name("UserService").Base(base).UseModel(
		artisan.Model(artisan.Name("user"), &userModel),
		artisan.Model(artisan.Name("vUser"), &vUserModel))
	return svc
}
