package httpRouter

import (
	_ "github.com/ZbgIast/ZbgIast-agent-go/core/http/httpRequestCookie"
	_ "github.com/ZbgIast/ZbgIast-agent-go/core/http/httpRequestFormValue"
	_ "github.com/ZbgIast/ZbgIast-agent-go/core/http/urlURLQuery"
	_ "github.com/ZbgIast/ZbgIast-agent-go/core/httpRouter/httpRouter"
	"github.com/ZbgIast/ZbgIast-agent-go/global"
	"github.com/ZbgIast/ZbgIast-agent-go/hook"
)

func init() {
	h := new(hook.HttpRouter)
	global.AllHooks = append(global.AllHooks, h)
	h.HookAll()
}
