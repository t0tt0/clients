package dblayer

import "github.com/Myriad-Dreamin/go-ves/lib/encoding"

func decodeBase64(src string) []byte {
	b, err := encoding.DecodeBase64(src)
	if err != nil {
		p.Logger.Debug("decode failed", "error", err, "source", src)
		return nil
	}
	return b
}

func DecodeAddress(src string) []byte {
	return decodeBase64(src)
}

func EncodeAddress(src []byte) string {
	return encoding.EncodeBase64(src)
}

func DecodeContent(src string) []byte {
	return decodeBase64(src)
}

func EncodeContent(src []byte) string {
	return encoding.EncodeBase64(src)
}

