package main

import (
	"fmt"
	"github.com/Myriad-Dreamin/artisan"
	"github.com/Myriad-Dreamin/go-ves/vesx/types"
)

var codeField = artisan.Param("code", new(types.CodeRawType))
var required = artisan.Tag("binding", "required")

func main() {
	v1 := "v1"

	//instantiate
	objectCate := DescribeObjectService(v1)

	//to files
	objectCate.ToFile("object.go")

	err := artisan.NewService(
		objectCate,
	).Publish()
	if err != nil {
		fmt.Println(err)
	}
}
