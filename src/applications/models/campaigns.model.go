package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Campaigns struct {
	ID           uuid.UUID  `gorm:"type:uuid;primaryKey" json:"id"`
	WorkspaceID  uuid.UUID  `gorm:"type:uuid" json:"workspace_id"`
	Workspace    Workspaces `gorm:"foreignKey:WorkspaceID;references:ID;constraint:OnDelete:CASCADE" json:"workspace"`
	CampaignCode string     `gorm:"type:text;uniqueIndex;not null" json:"campaign_code"`
	Title        string     `gorm:"type:text;" json:"title"`
	Slug         string     `gorm:"type:text;" json:"slug"`
	Descriptions string     `gorm:"type:text;" json:"descriptions"`
	Remark       string     `gorm:"type:text;" json:"remark"`
	Status       string     `gorm:"type:char(2);default:'S1';comment:'S1=DRAFT,S2=PUBLISH'" json:"status"`
	Deleted      bool       `gorm:"default:false" json:"deleted"`
	CreatedAt    time.Time  `gorm:"type:timestamp;autoCreateTime" json:"created_at"`
	UpdatedAt    time.Time  `gorm:"type:timestamp" json:"updated_at"`
}

func (Campaigns) TableName() string {
	return "campaigns" // name of table will created
}

func (campaign *Campaigns) BeforeCreate(tx *gorm.DB) (err error) {
	campaign.ID = uuid.New()
	return
}
