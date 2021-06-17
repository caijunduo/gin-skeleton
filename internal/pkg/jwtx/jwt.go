package jwtx

import (
    "github.com/dgrijalva/jwt-go"
    "github.com/spf13/viper"
    "github.com/uniplaces/carbon"
)

var (
    jwtKey = []byte(viper.GetString("jwt.default.key"))
)

type Claims struct {
    UUID string
    jwt.StandardClaims
}

func New(uuid string) (token string, err error) {
    c := &Claims{
        UUID: uuid,
        StandardClaims: jwt.StandardClaims{
            ExpiresAt: carbon.Now().AddMinutes(viper.GetInt("jwt.default.expireMinutes")).Unix(), // 过期时间
            IssuedAt:  carbon.Now().Unix(),
            Issuer:    viper.GetString("jwt.default.issuer"),
            Subject:   viper.GetString("jwt.default.subject"),
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
