package api

import (
	"encoding/json"
	"github.com/ZbgIast/ZbgIast-agent-go/global"
	"github.com/ZbgIast/ZbgIast-agent-go/utils"
	"github.com/parnurzeal/gorequest"
	"time"
)

// NewRequest
// 空白的请求工具 预留做后续的统一参数处理
func NewRequest() *gorequest.SuperAgent {
	request := gorequest.New()
	request.Timeout(30 * time.Second)
	return request
}

// POST
// 统一的POST前置处理
// url 请求api
// token 身份标识
func POST(url string, body interface{}) *gorequest.SuperAgent {
	s, _ := json.Marshal(body)
	jsonStr := utils.GzipStr(string(s))
	request := NewRequest()
	request.Post(global.Config.DongtaiGoOpenapi+url).
		Set("Content-Type", "application/json").
		Set("Authorization", "Token "+global.Config.DongtaiGoToken).
		Send(jsonStr)
	return request
}

// GET
// 统一的GET前置处理
// url 请求api
// token 身份标识
func GET(url string, body interface{}) *gorequest.SuperAgent {
	request := NewRequest()
	request.Get(global.Config.DongtaiGoOpenapi+url).Set("Authorization", "Token "+global.Config.DongtaiGoToken).Query(body)
	return request
}
