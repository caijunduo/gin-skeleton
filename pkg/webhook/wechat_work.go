package webhook

import (
    "fmt"
    "github.com/gin-gonic/gin"
    "github.com/go-resty/resty/v2"
    "log"
    "skeleton/configs"
    "time"
)

func WeChatWork(title string) *weChatWork {
    w := &weChatWork{
        title: title,
    }
    w.color = w.getColor()
    return w
}

type weChatWork struct {
    title   string
    content string
    items   string

    color string
}

func (w *weChatWork) getColor() string {
    switch gin.Mode() {
    case gin.TestMode:
        return "info"
    case gin.ReleaseMode:
        return "warning"
    default:
        return "comment"
    }
}

func (w *weChatWork) Content(content string) *weChatWork {
    w.content = content + "\n"
    return w
}

func (w *weChatWork) before() string {
    return fmt.Sprintf("> Datetime：<font color=\"%s\">%s</font>\n", w.color, time.Now().Format("2006-01-02 15:04:05")) +
        fmt.Sprintf("> Environment：<font color=\"%s\">%s</font>\n", w.color, gin.Mode())
}

func (w *weChatWork) Item(key string, val string) *weChatWork {
    w.items += fmt.Sprintf("> %s：<font color=\"%s\">%s</font>\n", key, w.color, val)
    return w
}

func (w *weChatWork) result() string {
    title := ""
    if w.content == "" {
        title = w.title + "\n"
    } else {
        title = w.title + "：" + w.content
    }
    return title + w.before() + w.items
}

func (w *weChatWork) Send() {
    if gin.IsDebugging() {
        return
    }
    _, err := resty.New().R().
        SetBody(map[string]interface{}{
            "msgtype": "markdown",
            "markdown": map[string]interface{}{
                "content": w.result(),
            },
        }).Post(configs.Config.Webhook.WechatWork.Url)
    if err != nil {
        log.Println("[WebHook] 异常推送失败 ", err)
        return
    }
}
