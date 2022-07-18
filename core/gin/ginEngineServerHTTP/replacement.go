package ginEngineServerHTTP

import (
	"bufio"
	"bytes"
	"encoding/base64"
	"github.com/ZbgIast/ZbgIast-agent-go/api"
	"github.com/ZbgIast/ZbgIast-agent-go/global"
	"github.com/ZbgIast/ZbgIast-agent-go/model/request"
	"github.com/ZbgIast/ZbgIast-agent-go/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"reflect"
	"strconv"
	"strings"
)

func MyServer(server *gin.Engine, w http.ResponseWriter, r *http.Request) {
	MyServerTemp(server, w, r)
	id := utils.CatGoroutineID()
	worker, _ := utils.NewWorker(global.AgentId)
	go func() {
		t := reflect.ValueOf(r.Body)
		var headerBase string
		body := ""
		for k, v := range r.Header {
			headerBase += k + ": " + strings.Join(v, ",") + "\n"
		}

		TraceId := global.TraceId + "-" + strconv.Itoa(int(worker.GetId()))
		global.TargetTraceId = TraceId
		tranceID := TraceId + "." + strconv.Itoa(global.AgentId) + ".0.0.0"
		headerBase += "dt-traceid:" + tranceID
		if t.Kind() == reflect.Ptr {
			buf := t.
				Elem().
				FieldByName("src").
				Elem().Elem().
				FieldByName("R").
				Elem().Elem().
				FieldByName("buf").Bytes()
			buf = buf[:bytes.IndexByte(buf, 0)]
			reader := bufio.NewReader(bytes.NewReader(buf))
			var reqArr []string
			for {
				line, _, err := reader.ReadLine()
				if err != nil {
					break
				}
				reqArr = append(reqArr, string(line))
			}
			body = reqArr[len(reqArr)-1]
		}
		header := base64.StdEncoding.EncodeToString([]byte(headerBase))
		scheme := "http"
		if r.TLS != nil {
			scheme = "https"
		}
		onlyKey := int(worker.GetId())
		HookGroup := &request.UploadReq{
			Type:     36,
			InvokeId: onlyKey,
			Detail: request.Detail{
				AgentId: global.AgentId,
				Function: request.Function{
					Method:        r.Method,
					Url:           scheme + "://" + r.Host + r.RequestURI,
					Uri:           r.URL.Path,
					Protocol:      r.Proto,
					ClientIp:      r.RemoteAddr,
					Language:      "GO",
					ReplayRequest: false,
					ReqHeader:     header,
					ReqBody:       body,
					QueryString:   r.URL.RawQuery,
					Pool:          []request.Pool{},
					TraceId:       tranceID,
				},
			},
		}
		var resBody string
		res, ok := global.ResponseMap.Load(id)
		if ok {
			global.ResponseMap.Delete(id)
			resBody = res.(string)
		}

		var resH string
		value2, ok1 := global.ResponseHeaderMap.Load(id)
		if ok1 {
			global.ResponseHeaderMap.Delete(id)
			resH = value2.(string)
		}
		for k, v := range w.Header() {
			resH += k + ": " + strings.Join(v, ",") + "\n"
		}
		resHeader := base64.StdEncoding.EncodeToString([]byte(resH))
		HookGroup.Detail.ResHeader = resHeader
		HookGroup.Detail.ResBody = resBody
		goroutineIDs := make(map[string]bool)
		global.PoolTreeMap.Range(func(key, value interface{}) bool {
			if value.(*request.PoolTree).IsThisBegin(id) {
				global.PoolTreeMap.Delete(key)
				value.(*request.PoolTree).FMT(&HookGroup.Detail.Function.Pool, worker, goroutineIDs, HookGroup.Detail.Function.TraceId)
				return false
			}
			return true
		})
		api.ReportUpload(*HookGroup)
		request.RunMapGCbYGoroutineID(goroutineIDs)
	}()
	return
}

func MyServerTemp(server *gin.Engine, w http.ResponseWriter, r *http.Request) {
	for i := 0; i < 100; i++ {

	}
	return
}
