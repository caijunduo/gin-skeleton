package jwtx

import (
    "github.com/dgrijalva/jwt-go"
    "github.com/spf13/cast"
    "github.com/uniplaces/carbon"
    "os"
)

var (
    jwtKey = []byte(os.Getenv("JWT_KEY"))
)

type Claims struct {
    UUID string
    jwt.StandardClaims
}

func New(uuid string) (token string, err error) {
    c := &Claims{
        UUID: uuid,
        StandardClaims: jwt.StandardClaims{
            ExpiresAt: carbon.Now().AddMinutes(cast.ToInt("JWT_EXPIRE_MINUTES")).Unix(), // 过期时间
            IssuedAt:  carbon.Now().Unix(),
            Issuer:    os.Getenv("JWT_ISSUER"),
            Subject:   os.Getenv("JWT_SUBJECT"),
        },
    }
    t := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
    token, err = t.SignedString(jwtKey)
    if err != nil {
        return
    }
    return
}

func Parse(token string) (t *jwt.Token, err error) {
    t, err = jwt.ParseWithClaims(token, &Claims{}, func(t *jwt.Token) (i interface{}, err error) {
        return jwtKey, nil
    })
    return
}
