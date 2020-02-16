package main

import (
	"bytes"
	"fmt"
	opintent "github.com/HyperService-Consortium/go-uip/op-intent"
	"github.com/HyperService-Consortium/go-uip/uiptypes"
	"github.com/Myriad-Dreamin/artisan"
	"github.com/Myriad-Dreamin/go-ves/ves/model"
	"github.com/Myriad-Dreamin/go-ves/ves/model/fset"
	"github.com/Myriad-Dreamin/minimum-lib/sugar"
	"log"
	"os"
	"reflect"

	"golang.org/x/tools/imports"
)

//var codeField = artisan.Param("code", new(types.CodeRawType))
//var required = artisan.Tag("binding", "required")

type Struct struct {
	structType, elemStructType reflect.Type
	methods                    []reflect.Method
	valueMethods               []reflect.Value
}

func getElements(i interface{}) (reflect.Value, reflect.Type) {
	return getReflectElements(reflect.ValueOf(i))
}

func getReflectElements(v reflect.Value) (reflect.Value, reflect.Type) {
	t := v.Type()
	for t.Kind() == reflect.Ptr {
		v, t = v.Elem(), t.Elem()
	}
	return v, t
}

func getReflectTypeElementType(t reflect.Type) reflect.Type {
	for t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	return t
}

func newStruct(i interface{}) *Struct {
	return newStructByType(reflect.TypeOf(i))
}

func newStructByType(t reflect.Type) *Struct {
	var (
		et = getReflectTypeElementType(t)
		s     = &Struct{
			structType:     t,
			elemStructType: et,
			methods:        nil,
		}
	)
	for i := 0; i < t.NumMethod(); i++ {
		s.methods = append(s.methods, t.Method(i))
	}
	return s
}

func (s Struct) MockInterface() string {
	return fmt.Sprintf(
		`type %sI interface {
%s}`, s.elemStructType.Name(), s.interfaceList(1))
}

func FormatCode(code string) ([]byte, error) {
	opts := &imports.Options{
		TabIndent: true,
		TabWidth:  2,
		Fragment:  true,
		Comments:  true,
	}
	return imports.Process("", []byte(code), opts)
}

func (s Struct) importList(pkg artisan.PackageSet) {
	for _, method := range s.methods {
		t := method.Type
		for i := hasRecv(s.elemStructType); i < t.NumIn(); i++ {
			in := getReflectTypeElementType(t.In(i))
			if len(in.PkgPath()) != 0 {
				pkg[in.PkgPath()] = true
			}
		}
		for i := 0; i < t.NumOut(); i++ {
			out := getReflectTypeElementType(t.Out(i))
			if len(out.PkgPath()) != 0 {
				pkg[out.PkgPath()] = true
			}
		}
	}
}

func hasRecv(elemStructType reflect.Type) int {
	if elemStructType.Kind() == reflect.Interface {
		return 0
	}
	return 1
}

func (s Struct) interfaceListToStream(indentCount int, stream *bytes.Buffer) {
	for _, method := range s.methods {
		t := method.Type
		writeIndent(stream, indentCount)
		stream.WriteString(method.Name)
		stream.WriteByte('(')
		for i := hasRecv(s.elemStructType); i < t.NumIn(); i++ {
			in := t.In(i)
			if i == t.NumIn()-1 {
				if t.IsVariadic() {
					stream.WriteString("...")
					stream.WriteString(in.Elem().String())
				} else {
					stream.WriteString(in.String())
				}
			} else {
				stream.WriteString(in.String())
			}
			if i != t.NumIn()-1 {
				stream.WriteByte(',')
			}
		}
		stream.WriteByte(')')
		if t.NumOut() > 1 {
			stream.WriteByte('(')
		}
		for i := 0; i < t.NumOut(); i++ {
			out := t.Out(i)
			stream.WriteString(out.String())
			if i != t.NumOut()-1 {
				stream.WriteByte(',')
			}
		}
		if t.NumOut() > 1 {
			stream.WriteByte(')')
		}
		stream.WriteByte('\n')
	}
}

func writeIndent(buffer *bytes.Buffer, indentCount int) {
	for i := 0; i < indentCount; i++ {
		buffer.WriteString("    ")
	}
}

func (s Struct) interfaceList(indentCount int) string {
	var b = bytes.NewBuffer(nil)
	s.interfaceListToStream(indentCount, b)
	return b.String()
}

var codeField = artisan.Param("code", artisan.Int)

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
	mockList := []*Struct{
		newStruct(&model.SessionDB{}),
		newStruct(&model.SessionAccountDB{}),
		newStruct(&fset.SessionFSet{}),
	}

	var x uiptypes.BlockChainInterface
	externMockList := []*Struct{
		newStruct(&opintent.OpIntentInitializer{}),
		newStructByType(reflect.TypeOf(&x).Elem()),
	}
	for _, s := range [] struct{
		l []*Struct
		fp string
	}{
		{mockList, "./gen-model-interface.go"}, {externMockList, "./gen-external-interface.go"},
	} {
		printMock(s.l, s.fp)
	}
}

func printMock(mockList []*Struct, filepath string) {
	var pkg = make(artisan.PackageSet)

	for _, s := range mockList {
		s.importList(pkg)
	}
	code := `
package control

import(
`
	code += printPkgPaths(pkg)
	code += ")\n"

	for _, s := range mockList {
		code += s.MockInterface() + "\n\n"
	}
	c, err := FormatCode(code)
	if err != nil {
		log.Fatal(err)
	}
	sugar.WithWriteFile(func(f *os.File) {
		_, err := f.Write(c)
		if err != nil {
			log.Fatal(err)
		}
	}, filepath)
}

func printPkgPaths(pkg artisan.PackageSet) string {
	var b = bytes.NewBuffer(nil)
	for k := range pkg {
		writeIndent(b, 1)
		b.WriteByte('"')
		b.WriteString(k)
		b.WriteByte('"')
		b.WriteByte('\n')
	}
	return b.String()
}
