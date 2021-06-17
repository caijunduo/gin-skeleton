package webhook

import (
    "github.com/spf13/viper"
    "os"
)

type wechatWork struct {
    *webhook
    mentionedList       []string
    mentionedMobileList []string
}

func (w *wechatWork) SetMentionedList(mentionedList []string) {
    w.mentionedList = mentionedList
}

func (w *wechatWork) SetMentionedMobileList(mentionedMobileList []string) {
    w.mentionedMobileList = mentionedMobileList
}

func (w wechatWork) Text() {
    w.send(map[string]interface{}{
        "msgtype": "text",
        "text": map[string]interface{}{
            "content":               w.content,
            "mentioned_list":        w.mentionedList,
            "mentioned_mobile_list": w.mentionedMobileList,
        },
    })
}

func (w wechatWork) Markdown() {
    w.url = os.Getenv("WEBHOOK_WECHAT_WORK_URL")
    w.send(map[string]interface{}{
        "msgtype": "markdown",
        "markdown": map[string]interface{}{
            "content":               w.content,
            "mentioned_list":        w.mentionedList,
            "mentioned_mobile_list": w.mentionedMobileList,
        },
    })
}

func NewWeChatWork() *wechatWork {
    return &wechatWork{
        webhook: &webhook{
            content: "",
            url:     viper.GetString("webhook.wechat_work.url"),
        },
        mentionedList:       make([]string, 0),
        mentionedMobileList: make([]string, 0),
    }
}
