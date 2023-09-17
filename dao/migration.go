package dao

import (
	"fmt"

	"github.com/ZhangMuran/GinMall/model"
)

func migrate() (err error) {
	err = _db.Set("gorm:table_options", "charset=utf8mb4").
		AutoMigrate(
			&model.User{},
		)
	if err != nil {
		fmt.Println("err =", err)
	}
	fmt.Println("creat success")
	return
}