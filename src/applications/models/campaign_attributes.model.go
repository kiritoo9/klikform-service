package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CampaignFormAttributes struct {
	ID             uuid.UUID     `gorm:"type:uuid;primaryKey" json:"id"`
	CampaignFormID uuid.UUID     `gorm:"type:uuid" json:"campaign_form_id"`
	CampaignForm   CampaignForms `gorm:"foreignKey:CampaignFormID;references:ID;constraint:OnDelete:CASCADE" json:"campaign_form"`
	Label          string        `gorm:"type:text" json:"label"`
	DefaultValue   string        `gorm:"type:text" json:"default_value"`
	Remark         string        `gorm:"type:text" json:"remark"`
	Deleted        bool          `gorm:"default:false" json:"deleted"`
	CreatedAt      time.Time     `gorm:"type:timestamp;autoCreateTime" json:"created_at"`
	UpdatedAt      time.Time     `gorm:"type:timestamp" json:"updated_at"`
}

func (CampaignFormAttributes) TableName() string {
	return "campaign_form_attributes" // name of table will created
}

func (campaign_form_attributes *CampaignFormAttributes) BeforeCreate(tx *gorm.DB) (err error) {
	campaign_form_attributes.ID = uuid.New()
	return
}
