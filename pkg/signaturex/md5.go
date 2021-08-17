package signaturex

import (
    "errors"
    "fmt"
    "github.com/spf13/cast"
    "github.com/uniplaces/carbon"
    "skeleton/pkg/helper"
    "sort"
)

type md5 struct {
    opt             Option
    data            map[string]interface{}
    beforeEncrypt   string
    beforeSignature string
    signature       string
}

func (m *md5) SetData(data map[string]interface{}) *md5 {
    m.data = data
    return m
}

func (m *md5) encrypt() *md5 {
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

func (m *md5) Verify() error {
    ak := cast.ToString(m.data["ak"])
    sn := cast.ToString(m.data["sn"])
    ts := cast.ToInt64(m.data["ts"])
    if ak == "" || sn == "" || ts <= 0 {
        return errors.New("signature field is missing")
    }

    if ak != m.opt.AppKey {
        return errors.New("signature app key error")
    }

    expires := m.opt.Expires

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

func (m *md5) Generate() *md5 {
    appSecret := m.opt.AppSecret
    m.encrypt()
    m.beforeSignature = appSecret + m.beforeEncrypt + appSecret
    m.signature = helper.Crypto.MD5(m.beforeSignature)
    return m
}

func (m md5) GetBeforeEncrypt() string {
    return m.beforeEncrypt
}

func (m md5) GetBeforeSignature() string {
    return m.beforeSignature
}

func (m md5) GetSignature() string {
    return m.signature
}

func NewMd5(opt Option) *md5 {
    return &md5{opt: opt}
}
