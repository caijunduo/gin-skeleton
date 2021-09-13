package jwt

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/uniplaces/carbon"
)

type Option struct {
	Key           string
	Issuer        string
	Subject       string
	ExpireMinutes int
}

type Claims struct {
	Opt  Option
	Data map[string]interface{}
	jwt.StandardClaims
}

func (j *Claims) SetData(data map[string]interface{}) *Claims {
	j.Data = data
	return j
}

func (j *Claims) Generate() (token string, err error) {
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

func (j *Claims) Parse(token string) (bool, map[string]interface{}, error) {
	t, err := jwt.ParseWithClaims(token, &Claims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(j.Opt.Key), nil
	})
	if err != nil {
		return false, nil, err
	}
	return t.Valid, t.Claims.(*Claims).Data, err
}
