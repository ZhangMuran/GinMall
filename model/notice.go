package model

import "gorm.io/gorm"

/*
CREATE TABLE `notice` (
	`id` bigint unsigned AUTO_INCREMENT,
	`created_at` datetime NULL,
	`updated_at` datetime NULL,
	`deleted_at` datetime NULL,
	`text` text,PRIMARY KEY (`id`),
	INDEX `idx_notice_deleted_at` (`deleted_at`)
)charset=utf8mb4
*/

type Notice struct {
	gorm.Model
	Text string `gorm:"type:text"`
}