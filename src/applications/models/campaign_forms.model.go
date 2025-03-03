package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CampaignForms struct {
	ID           uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	CampaignID   uuid.UUID `gorm:"type:uuid" json:"campaign_id"`
	Campaign     Campaigns `gorm:"foreignKey:CampaignID;references:ID;constraint:OnDelete:CASCADE" json:"campaign"`
	IsRequired   bool      `gorm:"default:false" json:"is_required"`
	IsReadonly   bool      `gorm:"default:false" json:"is_readonly"`
	Label        string    `gorm:"type:text" json:"label"`
	Descriptions string    `gorm:"type:text" json:"descriptions"`
	Remark       string    `gorm:"type:text" json:"remark"`
	Deleted      bool      `gorm:"default:false" json:"deleted"`
	CreatedAt    time.Time `gorm:"type:timestamp;autoCreateTime" json:"created_at"`
	UpdatedAt    time.Time `gorm:"type:timestamp" json:"updated_at"`
}

func (CampaignForms) TableName() string {
	return "campaign_forms" // name of table will created
}

func (campaign_form *CampaignForms) BeforeCreate(tx *gorm.DB) (err error) {
	campaign_form.ID = uuid.New()
	return
}
