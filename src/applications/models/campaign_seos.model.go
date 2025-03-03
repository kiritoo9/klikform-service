package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CampaignSeos struct {
	ID         uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	CampaignID uuid.UUID `gorm:"type:uuid" json:"campaign_id"`
	Campaign   Campaigns `gorm:"foreignKey:CampaignID;references:ID;constraint:OnDelete:CASCADE" json:"campaign"`
	Platform   string    `gorm:"type:char(2);default:'NN';comment:'FB=FACEBOOK,GO=GOOGLE,TW=TWITTER,IG=INSTAGRAM'" json:"platform"`
	Type       string    `gorm:"type:varchar(50)" json:"type"`
	Code       string    `gorm:"type:text;" json:"code"`
	Remark     string    `gorm:"type:text;" json:"remark"`
	Deleted    bool      `gorm:"default:false" json:"deleted"`
	CreatedAt  time.Time `gorm:"type:timestamp;autoCreateTime" json:"created_at"`
	UpdatedAt  time.Time `gorm:"type:timestamp" json:"updated_at"`
}

func (CampaignSeos) TableName() string {
	return "campaign_seos" // name of table will created
}

func (campaign_seo *CampaignSeos) BeforeCreate(tx *gorm.DB) (err error) {
	campaign_seo.ID = uuid.New()
	return
}
