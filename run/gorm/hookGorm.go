package gorm

import (
	_ "github.com/ZbgIast/ZbgIast-agent-go/core/gorm/gormDBExec"
	_ "github.com/ZbgIast/ZbgIast-agent-go/core/gorm/gormDBGroup"
	_ "github.com/ZbgIast/ZbgIast-agent-go/core/gorm/gormDBHaving"
	_ "github.com/ZbgIast/ZbgIast-agent-go/core/gorm/gormDBOrder"
	_ "github.com/ZbgIast/ZbgIast-agent-go/core/gorm/gormDBPluck"
	_ "github.com/ZbgIast/ZbgIast-agent-go/core/gorm/gormDBRaw"
	_ "github.com/ZbgIast/ZbgIast-agent-go/core/gorm/gormDBSelect"
	"github.com/ZbgIast/ZbgIast-agent-go/global"
	"github.com/ZbgIast/ZbgIast-agent-go/hook"
)

func init() {
	g := new(hook.Gorm)
	global.AllHooks = append(global.AllHooks, g)
	g.HookAll()
}
