package gormDBHaving

import (
	"github.com/ZbgIast/ZbgIast-agent-go/model/request"
	"gorm.io/gorm"
)

func Having(db *gorm.DB, query interface{}, args ...interface{}) (tx *gorm.DB) {
	s := HavingT(db, query, args...)
	request.FmtHookPool(request.PoolReq{
		Args:            request.Collect(query, args),
		Reqs:            request.Collect(tx),
		Source:          false,
		OriginClassName: "gorm.(*DB)",
		MethodName:      "Having",
		ClassName:       "gorm",
	})
	return s
}

func HavingT(db *gorm.DB, query interface{}, args ...interface{}) (tx *gorm.DB) {
	return
}
