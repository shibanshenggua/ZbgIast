package base

import (
	_ "github.com/ZbgIast/ZbgIast-agent-go/core/base/bufioWriterWrite"
	_ "github.com/ZbgIast/ZbgIast-agent-go/core/base/bufioWriterWriteString"
	_ "github.com/ZbgIast/ZbgIast-agent-go/core/base/execCmdStart"
	_ "github.com/ZbgIast/ZbgIast-agent-go/core/base/execCommand"
	_ "github.com/ZbgIast/ZbgIast-agent-go/core/base/fmtSprintf"
	_ "github.com/ZbgIast/ZbgIast-agent-go/core/base/htmlTemplateExecuteTemplate"
	_ "github.com/ZbgIast/ZbgIast-agent-go/core/base/httpClientDo"
	_ "github.com/ZbgIast/ZbgIast-agent-go/core/base/httpNewRequest"
	_ "github.com/ZbgIast/ZbgIast-agent-go/core/base/ioReadAll"
	_ "github.com/ZbgIast/ZbgIast-agent-go/core/base/jsonDecoderDecode"
	_ "github.com/ZbgIast/ZbgIast-agent-go/core/base/jsonNewDecoder"
	_ "github.com/ZbgIast/ZbgIast-agent-go/core/base/jsonUnmarshal"
	_ "github.com/ZbgIast/ZbgIast-agent-go/core/base/osOpenFile"
	_ "github.com/ZbgIast/ZbgIast-agent-go/core/base/regexpRegexpReplaceAllString"
	_ "github.com/ZbgIast/ZbgIast-agent-go/core/base/runtimeConcatstrings"
	_ "github.com/ZbgIast/ZbgIast-agent-go/core/base/runtimesSringtoslicebyte"
	_ "github.com/ZbgIast/ZbgIast-agent-go/core/base/sqlDBQuery"
	_ "github.com/ZbgIast/ZbgIast-agent-go/core/base/stringsBuilderWriteString"
	_ "github.com/ZbgIast/ZbgIast-agent-go/core/base/urlUrlString"
	"github.com/ZbgIast/ZbgIast-agent-go/hook"
	"github.com/ZbgIast/ZbgIast-agent-go/service"
	"github.com/ZbgIast/ZbgIast-agent-go/utils"
	"strconv"

	"github.com/ZbgIast/ZbgIast-agent-go/global"
)

func init() {
	b := new(hook.Base)
	global.AllHooks = append(global.AllHooks, b)
	global.InitViper()
	_ = service.AgentRegister()
	b.HookAll()
	worker, _ := utils.NewWorker(global.AgentId)
	global.TraceId = strconv.Itoa(int(worker.GetId()))
}
