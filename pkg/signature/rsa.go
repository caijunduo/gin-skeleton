package signature

import (
	"crypto"
	"crypto/rand"
	rsa2 "crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"
	"github.com/spf13/cast"
	"github.com/uniplaces/carbon"
	"io/ioutil"
	"sort"
)

type RSA struct {
	Opt             Option
	data            map[string]interface{}
	beforeEncrypt   string
	beforeSignature string
	signature       string
}

func (r *RSA) SetData(data map[string]interface{}) *RSA {
	r.data = data
	return r
}

func (r *RSA) encrypt() *RSA {
	var key []string
	for k := range r.data {
		if k != "sn" {
			key = append(key, k)
		}
	}

	sort.Strings(key)

	for i := 0; i < len(key); i++ {
		if i == 0 {
			r.beforeEncrypt = fmt.Sprintf("%v=%v", key[i], r.data[key[i]])
		} else {
			r.beforeEncrypt += fmt.Sprintf("&%v=%v", key[i], r.data[key[i]])
		}
	}
	return r
}

func (r RSA) encryptSHA1WithRSA() (signature string, err error) {
	key, err := ioutil.ReadFile(r.Opt.PrivateKeyPath)
	if err != nil {
		return
	}
	block, _ := pem.Decode(key)
	if block == nil {
		err = errors.New("pem.Decode failed")
		return
	}
	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return
	}
	hash := sha256.New()
	_, err = hash.Write([]byte(r.beforeEncrypt))
	if err != nil {
		return
	}
	sign, err := rsa2.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, hash.Sum(nil))
	if err != nil {
		return
	}
	signature = base64.StdEncoding.EncodeToString(sign)
	return
}

func (r RSA) decryptSHA1WithRSA(sn string) (err error) {
	signature, err := base64.StdEncoding.DecodeString(sn)
	if err != nil {
		return
	}
	key, err := ioutil.ReadFile(r.Opt.PublicKeyPath)
	if err != nil {
		return
	}
	block, _ := pem.Decode(key)
	if block == nil {
		err = errors.New("pem.Decode failed")
		return
	}
	publicKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return
	}
	hash := sha256.New()
	if _, err = hash.Write([]byte(r.beforeEncrypt)); err != nil {
		return
	}

	err = rsa2.VerifyPKCS1v15(publicKey.(*rsa2.PublicKey), crypto.SHA256, hash.Sum(nil), signature)
	return
}

func (r *RSA) Generate() error {
	appSecret := r.Opt.AppSecret
	r.encrypt()
	r.beforeSignature = appSecret + r.beforeEncrypt + appSecret
	s, err := r.encryptSHA1WithRSA()
	if err != nil {
		return err
	}
	r.signature = s
	return nil
}

func (r *RSA) Verify() error {
	ak := cast.ToString(r.data["ak"])
	sn := cast.ToString(r.data["sn"])
	ts := cast.ToInt64(r.data["ts"])
	if ak == "" || sn == "" || ts <= 0 {
		return errors.New("signature field is missing")
	}

	if ak != r.Opt.AppKey {
		return errors.New("signature app key error")
	}

	expires := r.Opt.Expires
	if expires > 0 {
		now := carbon.Now()
		now.SetTimestamp(ts)
		if now.Between(carbon.Now(), carbon.Now().AddMinutes(expires), false) {
			return errors.New("signature expired")
		}
	}

	appSecret := r.Opt.AppSecret
	r.encrypt()
	r.beforeSignature = appSecret + r.beforeEncrypt + appSecret
	if err := r.decryptSHA1WithRSA(sn); err != nil {
		return err
	}

	return nil
}

func (r RSA) GetBeforeEncrypt() string {
	return r.beforeEncrypt
}

func (r RSA) GetBeforeSignature() string {
	return r.beforeSignature
}

func (r RSA) GetSignature() string {
	return r.signature
}
