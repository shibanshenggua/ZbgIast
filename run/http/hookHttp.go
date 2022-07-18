package http

import (
	_ "github.com/ZbgIast/ZbgIast-agent-go/core/http/httpHeaderGet"
	_ "github.com/ZbgIast/ZbgIast-agent-go/core/http/httpRequestCookie"
	_ "github.com/ZbgIast/ZbgIast-agent-go/core/http/httpRequestFormValue"
	_ "github.com/ZbgIast/ZbgIast-agent-go/core/http/httpServeHTTP"
	_ "github.com/ZbgIast/ZbgIast-agent-go/core/http/urlURLQuery"
	"github.com/ZbgIast/ZbgIast-agent-go/global"
	"github.com/ZbgIast/ZbgIast-agent-go/hook"
)

func init() {
	h := new(hook.Http)
	global.AllHooks = append(global.AllHooks, h)
	h.HookAll()
}
