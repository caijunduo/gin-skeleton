package webhook

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
	"log"
	"skeleton/config"
	"time"
)

var WeCom = &weCom{}

type weCom struct {
	title   string
	content string
	items   string
}

func (w *weCom) reset() {
	w.title = ""
	w.content = ""
	w.items = ""
}

func (w *weCom) color() string {
	switch gin.Mode() {
	case gin.TestMode:
		return "info"
	case gin.ReleaseMode:
		return "warning"
	default:
		return "comment"
	}
}

func (w *weCom) before() string {
	return fmt.Sprintf("> Datetime：<font color=\"%s\">%s</font>\n", w.color(), time.Now().Format("2006-01-02 15:04:05")) +
		fmt.Sprintf("> Environment：<font color=\"%s\">%s</font>\n", w.color(), gin.Mode())
}

func (w *weCom) result() string {
	if w.content == "" {
		w.title += "\n"
	} else {
		w.title += "：" + w.content
	}
	return w.title + w.before() + w.items
}

func (w *weCom) Title(title string) *weCom {
	w.title = title
	return w
}

func (w *weCom) Content(content string) *weCom {
	w.content = content + "\n"
	return w
}

func (w *weCom) Item(key string, val string) *weCom {
	w.items += fmt.Sprintf("> %s：<font color=\"%s\">%s</font>\n", key, w.color(), val)
	return w
}

func (w *weCom) Send() {
	_, err := resty.New().R().
		SetBody(map[string]interface{}{
			"msgtype": "markdown",
			"markdown": map[string]interface{}{
				"content": w.result(),
			},
		}).Post(config.WebHook.WeCom.Url)
	w.reset()
	if err != nil {
		log.Println("[WebHook] 异常推送失败 ", err)
	}
}
