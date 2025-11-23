package models

import "backend/pkg/common"

type Task struct {
	common.GormBaseModel

	UserID      int     `gorm:"not null" json:"userId"`
	Title       string  `gorm:"type:text" json:"title"`
	Description string  `gorm:"type:text" json:"description"`
	Completed   bool    `gorm:"default:false" json:"completed"`
	DueDate     *string `gorm:"type:date" json:"dueDate"`
}
