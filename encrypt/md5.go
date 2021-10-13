package encrypt

import (
	"crypto/md5"
	"encoding/hex"
)

func Md5ToString(str string) string {
	return hex.EncodeToString(Md5ToByte(str))
}

func Md5ToByte(str string) []byte {
	s := md5.New()
	s.Write([]byte(str))
	return s.Sum(nil)
}
