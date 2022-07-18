package kafkaGo

import (
	_ "github.com/ZbgIast/ZbgIast-agent-go/core/kafkaGo/kafkaWriter"
	"github.com/ZbgIast/ZbgIast-agent-go/global"
	"github.com/ZbgIast/ZbgIast-agent-go/hook"
)

func init() {
	g := new(hook.KafkaGo)
	global.AllHooks = append(global.AllHooks, g)
	g.HookAll()
}
