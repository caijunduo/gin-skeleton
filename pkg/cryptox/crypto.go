package cryptox

import (
    "crypto/hmac"
    "crypto/md5"
    "crypto/sha256"
    "encoding/base64"
    "encoding/hex"
)

func MD5(str string) string {
    s := md5.New()
    s.Write([]byte(str))
    return hex.EncodeToString(s.Sum(nil))
}

func HmacSha256(str string, secret string) string {
    s := hmac.New(sha256.New, []byte(secret))
    s.Write([]byte(str))
    return base64.StdEncoding.EncodeToString(s.Sum(nil))
}
