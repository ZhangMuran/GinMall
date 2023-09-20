package dao

import (
	"context"

	"github.com/ZhangMuran/GinMall/model"
	"gorm.io/gorm"
)

type NoticeDao struct {
	*gorm.DB
}

func NewNoticeDao(c context.Context) *NoticeDao {
	return &NoticeDao{NewDBClient(c)}
}

func NewNoticeDaobyDB(db *gorm.DB) *NoticeDao {
	return &NoticeDao{db}
}

// GetNoticeById 根据id获取Notice
func (dao *NoticeDao) GetNoticeById(id uint, notice *model.Notice) error {
	return dao.DB.Model(&model.Notice{}).Where("id=?", id).First(notice).Error
}