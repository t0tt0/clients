package database

import (
	"encoding/base64"
	"github.com/Myriad-Dreamin/go-ves/lib/encoding"
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

