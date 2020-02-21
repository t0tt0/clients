package model

import "github.com/Myriad-Dreamin/go-ves/ves/model/internal/database"

func DecodeAddress(src string) ([]byte, error) {
	return database.DecodeAddress(src)
}
func EncodeAddress(src []byte) string {
	return database.EncodeAddress(src)
}

func DecodeContent(src string) ([]byte, error) {
	return database.DecodeContent(src)
}
func EncodeContent(src []byte) string {
	return database.EncodeContent(src)
}
