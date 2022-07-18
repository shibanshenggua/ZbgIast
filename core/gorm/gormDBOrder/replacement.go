package gormDBOrder

import (
	"github.com/ZbgIast/ZbgIast-agent-go/model/request"
	"gorm.io/gorm"
)

func Order(db *gorm.DB, value interface{}) (tx *gorm.DB) {

	s := OrderT(db, value)
	request.FmtHookPool(request.PoolReq{
		Args:            request.Collect(value),
		Reqs:            request.Collect(tx),
		Source:          false,
		OriginClassName: "gorm.(*DB)",
		MethodName:      "Order",
		ClassName:       "gorm",
	})
	return s
}

func OrderT(db *gorm.DB, value interface{}) (tx *gorm.DB) {
	return
}
