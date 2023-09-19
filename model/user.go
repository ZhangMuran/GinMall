package model

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

/*
CREATE TABLE `user` (
	`id` bigint unsigned AUTO_INCREMENT,
	`created_at` datetime NULL,
	`updated_at` datetime NULL,
	`deleted_at` datetime NULL,
	`user_name` varchar(256) UNIQUE COMMENT '用户名',
	`email` varchar(256) UNIQUE COMMENT '电子邮件地址',
	`password_digest` varchar(256) COMMENT '密码摘要',
	`nickname` varchar(256) NOT NULL COMMENT '昵称',
	`status` varchar(256) COMMENT '用户状态',
	`avatar` varchar(1000) COMMENT '头像',
	`money` bigint COMMENT '金钱数量',
	PRIMARY KEY (`id`),
	INDEX `idx_user_deleted_at` (`deleted_at`))charset=utf8mb4
*/

type User struct {
	gorm.Model
	UserName       string `gorm:"unique;comment:用户名"`
	Email          string `gorm:"unique;comment:电子邮件地址"`
	PasswordDigest string `gorm:"comment:密码摘要"`
	Nickname       string `gorm:"not null;comment:昵称"`
	Status         string `gorm:"comment:用户状态"`
	Avatar         string `gorm:"size:1000;comment:头像"`
	Money          int    `gorm:"comment:金钱数量"`
}

const (
	PassWordCost        = 12         //密码加密难度
	Active       string = "active"   //激活用户
)

//SetPassword 设置密码
func (user *User) SetPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), PassWordCost)
	if err != nil {
		return err
	}
	user.PasswordDigest = string(bytes)
	return nil
}

// CheckPassword 校验密码
func (u *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.PasswordDigest), []byte(password))
	return err == nil
}