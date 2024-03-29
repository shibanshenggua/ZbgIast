package ginContextGetPostFormArray

import (
	"github.com/ZbgIast/ZbgIast-agent-go/model/request"
	"github.com/gin-gonic/gin"
)

func GetPostFormArray(c *gin.Context, key string) ([]string, bool) {
	reArray, flag := GetPostFormArrayT(c, key)
	request.FmtHookPool(request.PoolReq{
		Args:            request.Collect(key),
		Reqs:            request.Collect(reArray),
		Source:          true,
		OriginClassName: "gin.(*Context)",
		MethodName:      "GetPostFormArray",
		ClassName:       "gin.(*Context)",
	})
	return reArray, flag
}

func GetPostFormArrayT(c *gin.Context, key string) ([]string, bool) {
	return []string{}, false
}
