package signature

import (
	"errors"
	"fmt"
	"github.com/spf13/cast"
	"github.com/uniplaces/carbon"
	"skeleton/encrypt"
	"sort"
)

type MD5 struct {
	Opt             Option
	data            map[string]interface{}
	beforeEncrypt   string
	beforeSignature string
	signature       string
}

func (m *MD5) SetData(data map[string]interface{}) *MD5 {
	m.data = data
	return m
}

func (m *MD5) encrypt() *MD5 {
	var key []string
	for k := range m.data {
		if k != "sn" {
			key = append(key, k)
		}
	}

	sort.Strings(key)

	for i := 0; i < len(key); i++ {
		if i == 0 {
			m.beforeEncrypt = fmt.Sprintf("%v=%v", key[i], m.data[key[i]])
		} else {
			m.beforeEncrypt += fmt.Sprintf("&%v=%v", key[i], m.data[key[i]])
		}
	}
	return m
}

func (m *MD5) Verify() error {
	ak := cast.ToString(m.data["ak"])
	sn := cast.ToString(m.data["sn"])
	ts := cast.ToInt64(m.data["ts"])
	if ak == "" || sn == "" || ts <= 0 {
		return errors.New("signature field is missing")
	}

	if ak != m.Opt.AppKey {
		return errors.New("signature app key error")
	}

	expires := m.Opt.Expires

	if expires > 0 {
		now := carbon.Now()
		now.SetTimestamp(ts)
		if now.Between(carbon.Now(), carbon.Now().AddMinutes(expires), false) {
			return errors.New("signature expired")
		}
	}

	m.Generate()

	if sn == "" || sn != m.signature {
		return errors.New("signature verification failed")
	}

	return nil
}

func (m *MD5) Generate() *MD5 {
	appSecret := m.Opt.AppSecret
	m.encrypt()
	m.beforeSignature = appSecret + m.beforeEncrypt + appSecret
	m.signature = encrypt.Md5ToString(m.beforeSignature)
	return m
}

func (m MD5) GetBeforeEncrypt() string {
	return m.beforeEncrypt
}

func (m MD5) GetBeforeSignature() string {
	return m.beforeSignature
}

func (m MD5) GetSignature() string {
	return m.signature
}
