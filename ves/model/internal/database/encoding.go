package database

import "github.com/HyperService-Consortium/go-ves/lib/basic/encoding"

func DecodeAddress(src string) ([]byte, error) {
	return encoding.DecodeBase64(src)
}

func EncodeAddress(src []byte) string {
	return encoding.EncodeBase64(src)
}

func DecodeContent(src string) ([]byte, error) {
	return encoding.DecodeBase64(src)
}

func EncodeContent(src []byte) string {
	return encoding.EncodeBase64(src)
}
