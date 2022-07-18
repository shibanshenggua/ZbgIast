package grpc

import (
	_ "github.com/ZbgIast/ZbgIast-agent-go/core/grpc/clientConn"
	_ "github.com/ZbgIast/ZbgIast-agent-go/core/grpc/newServer"
	"github.com/ZbgIast/ZbgIast-agent-go/global"
	"github.com/ZbgIast/ZbgIast-agent-go/hook"
)

func init() {
	g := new(hook.Grpc)
	global.AllHooks = append(global.AllHooks, g)
	g.HookAll()
}
