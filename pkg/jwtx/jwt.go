package jwtx

import (
    "github.com/dgrijalva/jwt-go"
    "github.com/uniplaces/carbon"
)

type JwtClaims struct {
    opt  Option
    Data map[string]interface{}
    jwt.StandardClaims
}

func (j *JwtClaims) SetData(data map[string]interface{}) *JwtClaims {
    j.Data = data
    return j
}

func (j *JwtClaims) Generate() (token string, err error) {
    j.ExpiresAt = carbon.Now().AddMinutes(j.opt.ExpireMinutes).Unix()
    j.IssuedAt = carbon.Now().Unix()
    j.Issuer = j.opt.Issuer
    j.Subject = j.opt.Subject

    t := jwt.NewWithClaims(jwt.SigningMethodHS256, j)
    if token, err = t.SignedString([]byte(j.opt.Key)); err != nil {
        return
    }
    return
}

func (j *JwtClaims) Parse(token string) (bool, map[string]interface{}, error) {
    t, err := jwt.ParseWithClaims(token, &JwtClaims{}, func(t *jwt.Token) (interface{}, error) {
        return []byte(j.opt.Key), nil
    })
    if err != nil {
        return false, nil, err
    }
    return t.Valid, t.Claims.(*JwtClaims).Data, err
}

type Option struct {
    Key           string
    Issuer        string
    Subject       string
    ExpireMinutes int
}

func New(opt Option) *JwtClaims {
    return &JwtClaims{opt: opt}
}
