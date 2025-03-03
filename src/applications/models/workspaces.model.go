package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Workspaces struct {
	ID           uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	Title        string    `gorm:"type:text;not null" json:"title"`
	Descriptions string    `gorm:"type:text;" json:"descriptions"`
	Thumbnail    string    `gorm:"type:varchar(100);" json:"thumbnail"`
	Status       string    `gorm:"type:char(2);default:'S1';comment:'S1=DRAFT,S2=PUBLISH'" json:"status"`
	Remark       string    `gorm:"type:text;" json:"remark"`
	Deleted      bool      `gorm:"default:false" json:"deleted"`
	CreatedAt    time.Time `gorm:"type:timestamp;autoCreateTime" json:"created_at"`
	UpdatedAt    time.Time `gorm:"type:timestamp" json:"updated_at"`
}

func (Workspaces) TableName() string {
	return "workspaces" // name of table will created
}

func (workspace *Workspaces) BeforeCreate(tx *gorm.DB) (err error) {
	workspace.ID = uuid.New()
	return
}
