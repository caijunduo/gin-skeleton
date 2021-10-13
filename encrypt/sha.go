package encrypt

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
)

func HmacSha256ToString(str string, secret string) string {
	return base64.StdEncoding.EncodeToString(HmacSha256ToByte(str, secret))
}

func HmacSha256ToByte(str string, secret string) []byte {
	s := hmac.New(sha256.New, []byte(secret))
	s.Write([]byte(str))
	return s.Sum(nil)
}

func Sha256ToString(str string) string {
	return base64.StdEncoding.EncodeToString(Sha256ToByte(str))
}

func Sha256ToByte(str string) []byte {
	s := sha256.New()
	s.Write([]byte(str))
	return s.Sum(nil)
}
