package dao

import (
	"github.com/ZhangMuran/GinMall/model"
)

func migrate() (err error) {
	err = _db.Set("gorm:table_options", "charset=utf8mb4").
		AutoMigrate(
			&model.User{},
			&model.Notice{},
		)
	return
}