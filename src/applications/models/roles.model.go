package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Roles struct {
	ID           uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	Name         string    `gorm:"type:varchar(50);uniqueIndex;not null" json:"name"`
	Descriptions string    `gorm:"type:text;" json:"descriptions"`
	Deleted      bool      `gorm:"default:false" json:"deleted"`
	CreatedAt    time.Time `gorm:"type:timestamp;autoCreateTime" json:"created_at"`
	UpdatedAt    time.Time `gorm:"type:timestamp" json:"updated_at"`
}

func (Roles) TableName() string {
	return "roles" // name of table will created
}

func (role *Roles) BeforeCreate(tx *gorm.DB) (err error) {
	role.ID = uuid.New()
	return
}
