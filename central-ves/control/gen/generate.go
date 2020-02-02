package main

import (
	"fmt"
	"github.com/Myriad-Dreamin/artisan"
	"github.com/Myriad-Dreamin/go-ves/central-ves/types"
)

var codeField = artisan.Param("code", new(types.CodeRawType))
var required = artisan.Tag("binding", "required")

func main() {
	v1 := "v1"

	//instantiate
	chainInfoCate := DescribeChainInfoService(v1)
	userCate := DescribeUserService(v1)
	objectCate := DescribeObjectService(v1)

	//to files
	chainInfoCate.ToFile("chain-info.go")
	userCate.ToFile("user.go")
	objectCate.ToFile("object.go")

	err := artisan.NewService(
		chainInfoCate,
		userCate,
		objectCate,
	).Publish()
	if err != nil {
		fmt.Println(err)
	}
}
