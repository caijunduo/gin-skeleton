package signaturex

import (
    "errors"
    "fmt"
    "github.com/spf13/cast"
    "github.com/spf13/viper"
    "github.com/uniplaces/carbon"
    "skeleton/pkg/cryptox"
    "sort"
)

type md5 struct {
    conn      string
    data      map[string]interface{}
    noEncrypt string
    noMd5     string
    md5       string
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
            m.noEncrypt = fmt.Sprintf("%v=%v", key[i], m.data[key[i]])
        } else {
            m.noEncrypt += fmt.Sprintf("&%v=%v", key[i], m.data[key[i]])
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

    key := "signature." + m.conn + "."

    if ak != viper.GetString(key+"appKey") {
        return errors.New("")
    }

    expires := viper.GetInt(key + "expires")

    if expires > 0 {
        now := carbon.Now()
        now.SetTimestamp(ts)
        if now.Between(carbon.Now(), carbon.Now().AddMinutes(expires), false) {
            return errors.New("signature app key error")
        }
    }

    m.Generate()

    if sn == "" || sn != m.md5 {
        return errors.New("signature verification failed")
    }

    return nil
}

func (m *md5) Generate() *md5 {
    appSecret := viper.GetString("signature." + m.conn + ".appSecret")
    m.encrypt()
    m.noMd5 = appSecret + m.noEncrypt + appSecret
    m.md5 = cryptox.MD5(m.noMd5)
    return m
}

func (m md5) GetNoEncrypt() string {
    return m.noEncrypt
}

func (m md5) GetNoMd5() string {
    return m.noMd5
}

func (m md5) GetMd5() string {
    return m.md5
}

func NewMd5(conn string) *md5 {
    return &md5{
        conn: conn,
    }
}
