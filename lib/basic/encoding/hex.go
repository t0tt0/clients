package encoding

import "encoding/hex"

func DecodeHex(src string) ([]byte, error) {
	return hex.DecodeString(src)
}

func EncodeHex(src []byte) string {
	return hex.EncodeToString(src)
}
