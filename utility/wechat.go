package utility

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gogf/gf/v2/frame/g"
)

// WechatSession 微信 code2session 返回结构
type WechatSession struct {
	OpenId     string `json:"openid"`
	SessionKey string `json:"session_key"`
	UnionId    string `json:"unionid"`
	ErrCode    int    `json:"errcode"`
	ErrMsg     string `json:"errmsg"`
}

// Code2Session 通过 code 换取 openid
func Code2Session(ctx context.Context, code string) (*WechatSession, error) {
	appId := g.Cfg().MustGet(ctx, "app.wechat.appId").String()
	appSecret := g.Cfg().MustGet(ctx, "app.wechat.appSecret").String()

	url := fmt.Sprintf(
		"https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code",
		appId, appSecret, code,
	)

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("请求微信接口失败: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("读取微信响应失败: %w", err)
	}

	var session WechatSession
	if err = json.Unmarshal(body, &session); err != nil {
		return nil, fmt.Errorf("解析微信响应失败: %w", err)
	}

	if session.ErrCode != 0 {
		return nil, fmt.Errorf("微信授权失败: %s", session.ErrMsg)
	}

	return &session, nil
}
