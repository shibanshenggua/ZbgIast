package gorilla

import (
	_ "github.com/ZbgIast/ZbgIast-agent-go/core/gorilla/gorillaRpcServerHTTP"
	"github.com/ZbgIast/ZbgIast-agent-go/global"
	"github.com/ZbgIast/ZbgIast-agent-go/hook"
)

func init() {
	g := new(hook.Gorilla)
	global.AllHooks = append(global.AllHooks, g)
	g.HookAll()
}
