package webhook

import (
    "fmt"
    "github.com/spf13/viper"
    "log"
    "skeleton/pkg/cryptox"
    "time"
)

type dingTalk struct {
    *webhook
    secret string

    title   string
    mobiles []string
    userIds []string
    all     bool
}

func (d *dingTalk) SetTitle(title string) {
    d.title = title
}

func (d *dingTalk) SetAtMobiles(atMobiles []string) {
    d.mobiles = atMobiles
}

func (d *dingTalk) SetAtUserIds(atUserIds []string) {
    d.userIds = atUserIds
}

func (d *dingTalk) SetIsAtAll(isAtAll bool) {
    d.all = isAtAll
}

func (d dingTalk) Text() {
    d.url = d.addSign()
    d.send(map[string]interface{}{
        "msgtype": "text",
        "text": map[string]interface{}{
            "content": d.content,
        },
        "at": map[string]interface{}{
            "atMobiles": d.mobiles,
            "atUserIds": d.userIds,
            "isAtAll":   d.all,
        },
    })
}

func (d dingTalk) Markdown() {
    d.url = d.addSign()
    log.Println(d.url)
    d.send(map[string]interface{}{
        "msgtype": "markdown",
        "markdown": map[string]interface{}{
            "title": d.title,
            "text":  d.content,
        },
        "at": map[string]interface{}{
            "atMobiles": d.mobiles,
            "atUserIds": d.userIds,
            "isAtAll":   d.all,
        },
    })
}

func (d *dingTalk) addSign() string {
    t := time.Now().UnixNano() / 1e6
    return fmt.Sprintf("%s&timestamp=%d&sign=%s",
        d.url,
        t,
        cryptox.HmacSha256(fmt.Sprintf("%d\n%s", t, d.secret), d.secret),
    )
}

func NewDingTalk() *dingTalk {
    return &dingTalk{
        webhook: &webhook{
            content: "",
            url:     viper.GetString("webhook.dingtalk.url"),
        },
        secret:  viper.GetString("webhook.dingtalk.secret"),
        title:   "",
        mobiles: make([]string, 0),
        userIds: make([]string, 0),
        all:     false,
    }
}
