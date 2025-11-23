package common

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type GormBaseModel struct {
	ID uuid.UUID `gorm:"type:uuid; primarykey; default:gen_random_uuid()" json:"id"`

	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`

	DeletedAt gorm.DeletedAt `gorm:"index" json:"deletedAt"`
}
