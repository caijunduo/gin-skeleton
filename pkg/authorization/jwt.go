package authorizationPkg

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/uniplaces/carbon"
)

type JwtOption struct {
	Key           string
	Issuer        string
	Subject       string
	ExpireMinutes int
}

type JwtClaims struct {
	Opt  JwtOption
	Data map[string]interface{}
	jwt.StandardClaims
}

func (j *JwtClaims) SetData(data map[string]interface{}) *JwtClaims {
	j.Data = data
	return j
}

func (j *JwtClaims) Generate() (token string, err error) {
	j.ExpiresAt = carbon.Now().AddMinutes(j.Opt.ExpireMinutes).Unix()
	j.IssuedAt = carbon.Now().Unix()
	j.Issuer = j.Opt.Issuer
	j.Subject = j.Opt.Subject

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, j)
	if token, err = t.SignedString([]byte(j.Opt.Key)); err != nil {
		return
	}
	return
}

func (j *JwtClaims) Parse(token string) (bool, map[string]interface{}, error) {
	t, err := jwt.ParseWithClaims(token, &JwtClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(j.Opt.Key), nil
	})
	if err != nil {
		return false, nil, err
	}
	return t.Valid, t.Claims.(*JwtClaims).Data, err
}
