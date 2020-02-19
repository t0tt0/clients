package model

import dblayer "github.com/Myriad-Dreamin/go-ves/ves/model/db-layer"

func DecodeAddress(src string) []byte {
	return dblayer.DecodeAddress(src)
}
func EncodeAddress(src []byte) string {
	return dblayer.EncodeAddress(src)
}

func DecodeContent(src string) []byte {
	return dblayer.DecodeContent(src)
}
func EncodeContent(src []byte) string {
	return dblayer.EncodeContent(src)
}
