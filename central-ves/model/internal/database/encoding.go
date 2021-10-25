package database

import (
	"encoding/base64"
	"github.com/HyperService-Consortium/go-ves/lib/basic/encoding"
)

func decodeBase64(src string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(src)
}

func EncodeAddress(src []byte) string {
	return encoding.EncodeBase64(src)
}

func DecodeAddress(src string) ([]byte, error) {
	return encoding.DecodeBase64(src)
}

