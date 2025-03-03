package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type FormAttributes struct {
	ID             uuid.UUID     `gorm:"type:uuid;primaryKey" json:"id"`
	UserID         uuid.UUID     `gorm:"type:uuid;not null" json:"user_id"`
	User           Users         `gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE" json:"user"`
	WorkspaceID    uuid.UUID     `gorm:"type:uuid;not null" json:"workspace_id"`
	Workspace      Workspaces    `gorm:"foreignKey:WorkspaceID;references:ID;constarint:OnDelete:CASCADE" json:"workspace"`
	CampaignID     uuid.UUID     `gorm:"type:uuid;not null" json:"campaign_id"`
	Campaign       Campaigns     `gorm:"foreignKey:CampaignID;references:ID;constraint:OnDelete:CASCADE" json:"campaign"`
	CampaignFormID uuid.UUID     `gorm:"type:uuid" json:"campaign_form_id"`
	CampaignForm   CampaignForms `gorm:"foreignKey:CampaignFormID;references:ID;constraint:OnDelete:CASCADE" json:"campaign_form"`
	Value          string        `gorm:"type:text" json:"value"`
	Deleted        bool          `gorm:"default:false" json:"deleted"`
	CreatedAt      time.Time     `gorm:"type:timestamp;autoCreateTime" json:"created_at"`
	UpdatedAt      time.Time     `gorm:"type:timestamp" json:"updated_at"`
}

func (FormAttributes) TableName() string {
	return "form_attributes" // name of table will created
}

func (form_attribute *FormAttributes) BeforeCreate(tx *gorm.DB) (err error) {
	form_attribute.ID = uuid.New()
	return
}
