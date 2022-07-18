package gin

import (
	_ "github.com/ZbgIast/ZbgIast-agent-go/core/gin/ginContextGetPostFormArray"
	_ "github.com/ZbgIast/ZbgIast-agent-go/core/gin/ginContextGetPostFormMap"
	_ "github.com/ZbgIast/ZbgIast-agent-go/core/gin/ginContextGetQueryArray"
	_ "github.com/ZbgIast/ZbgIast-agent-go/core/gin/ginContextGetQueryMap"
	_ "github.com/ZbgIast/ZbgIast-agent-go/core/gin/ginContextParam"
	_ "github.com/ZbgIast/ZbgIast-agent-go/core/gin/ginContextShouldBindBodyWith"
	_ "github.com/ZbgIast/ZbgIast-agent-go/core/gin/ginContextShouldBindUri"
	_ "github.com/ZbgIast/ZbgIast-agent-go/core/gin/ginContextShouldBindWith"
	_ "github.com/ZbgIast/ZbgIast-agent-go/core/gin/ginEngineServerHTTP"
	_ "github.com/ZbgIast/ZbgIast-agent-go/core/http/httpRequestCookie"
	"github.com/ZbgIast/ZbgIast-agent-go/global"
	"github.com/ZbgIast/ZbgIast-agent-go/hook"
)

func init() {
	g := new(hook.Gin)
	global.AllHooks = append(global.AllHooks, g)
	g.HookAll()
}
