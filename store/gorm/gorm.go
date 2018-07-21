package gorm

import (
	"errors"

	"github.com/jinzhu/gorm"
)

// SwitchDB 切换数据库
func SwitchDB(tran *gorm.DB, db *gorm.DB) *gorm.DB {
	if tran != nil {
		return tran
	}
	if db != nil {
		return db
	}
	panic(errors.New("事务切换失败"))
}
