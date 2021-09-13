package pkg

import (
	"bytes"
	"github.com/go-redis/redis"
	"golang.org/x/sync/errgroup"
	"skeleton/pkg/crypto"
	"skeleton/pkg/gin"
	"skeleton/pkg/jwt"
	"skeleton/pkg/signature"
	"skeleton/pkg/validator"
	"upper.io/db.v3/lib/sqlbuilder"
)

var (
	Group   errgroup.Group
	Context = &gin.Context{}
	Writer  = &gin.ResponseWriter{
		Body: bytes.NewBufferString(""),
	}
	Validator  = validator.Validator
	Translator = validator.Translator
	Jwt        = func(opt jwt.Option) *jwt.Claims {
		return &jwt.Claims{Opt: opt}
	}
	SignatureMD5 = func(opt signature.Option) *signature.MD5 {
		return &signature.MD5{Opt: opt}
	}
	SignatureRSA = func(opt signature.Option) *signature.RSA {
		return &signature.RSA{Opt: opt}
	}
	RedisDefault *redis.Client
	MySQLDefault sqlbuilder.Database
	Crypto       = crypto.Crypto
)
