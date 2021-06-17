package jwtx

import (
    "github.com/dgrijalva/jwt-go"
    "github.com/spf13/viper"
    "github.com/uniplaces/carbon"
)

type JwtClaims struct {
    conn string
    key  []byte
    Data map[string]interface{}
    jwt.StandardClaims
}

func (j *JwtClaims) SetData(data map[string]interface{}) *JwtClaims {
    j.Data = data
    return j
}

func (j JwtClaims) getConfigKey(key string) string {
    return j.conn + key
}

func (j *JwtClaims) Generate() (token string, err error) {
    j.ExpiresAt = carbon.Now().AddMinutes(viper.GetInt(j.getConfigKey("expireMinutes"))).Unix()
    j.IssuedAt = carbon.Now().Unix()
    j.Issuer = viper.GetString(j.getConfigKey("issuer"))
    j.Subject = viper.GetString(j.getConfigKey("subject"))

    t := jwt.NewWithClaims(jwt.SigningMethodHS256, j)
    if token, err = t.SignedString(j.key); err != nil {
        return
    }
    return
}

func (j *JwtClaims) Parse(token string) (bool, map[string]interface{}, error) {
    t, err := jwt.ParseWithClaims(token, &JwtClaims{}, func(t *jwt.Token) (interface{}, error) {
        return j.key, nil
    })
    if err != nil {
        return false, nil, err
    }
    return t.Valid, t.Claims.(*JwtClaims).Data, err
}

func New(conn string) *JwtClaims {
    key := "jwt." + conn + "."
    return &JwtClaims{
        conn: key,
        key:  []byte(viper.GetString(key + "key")),
    }
}
