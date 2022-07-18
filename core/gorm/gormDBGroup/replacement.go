package gormDBGroup

import (
	"github.com/ZbgIast/ZbgIast-agent-go/model/request"
	"gorm.io/gorm"
)

func Group(db *gorm.DB, value string) (tx *gorm.DB) {
	s := GroupT(db, value)
	request.FmtHookPool(request.PoolReq{
		Args:            request.Collect(value),
		Reqs:            request.Collect(tx),
		Source:          false,
		OriginClassName: "gorm.(*DB)",
		MethodName:      "Group",
		ClassName:       "gorm",
	})
	return s
}

func GroupT(db *gorm.DB, value string) (tx *gorm.DB) {
	return
}
