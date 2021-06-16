package webhook

import (
    "encoding/json"
    "go.uber.org/zap"
    "io/ioutil"
    "net/http"
    "strings"
)

type webhook struct {
    content string
    url     string
}

func (h *webhook) SetContent(content string) {
    h.content = content
}

func (h webhook) parseParameters(parameters map[string]interface{}) string {
    b, err := json.Marshal(parameters)
    if err != nil {
        zap.L().Error("[WebHook] parameters编码失败", zap.Any("parameters", parameters), zap.Error(err))
        return ""
    }
    return string(b)
}

func (h webhook) send(parameters map[string]interface{}) {
    c := &http.Client{}
    req, err := http.NewRequest("POST", h.url, strings.NewReader(h.parseParameters(parameters)))
    if err != nil {
        zap.L().Error("[WebHook] http NewRequest failed", zap.Any("parameters", parameters), zap.Error(err))
        return
    }
    req.Header.Set("Content-Type", "application/json")
    r, err := c.Do(req)
    if err != nil {
        zap.L().Error("[WebHook] http client Do failed", zap.Any("parameters", parameters), zap.Error(err))
        return
    }
    defer r.Body.Close()
    body, err := ioutil.ReadAll(r.Body)
    if err != nil {
        zap.L().Error("[WebHook] http client Do failed", zap.Any("parameters", parameters), zap.String("response", string(body)), zap.Error(err))
        return
    }
    zap.L().Info("[WebHook] Success", zap.Any("parameters", parameters), zap.String("body", string(body)))
}
