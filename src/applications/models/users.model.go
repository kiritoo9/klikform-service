package models

import (
	"time"

	"github.com/google/uuid"
)

type Users struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	Email     string    `gorm:"type:varchar(100);uniqueIndex;not null" json:"email"`
	Password  string    `gorm:"type:varchar(60);not null" json:"password"`
	Fullname  string    `gorm:"type:varchar(200)" json:"fullname"`
	Phone     string    `gorm:"type:varchar(20)" json:"phone"`
	Address   string    `gorm:"type:text" json:"address"`
	Status    string    `gorm:"type:char(2);not null;default:'S1';comment: 'S1=NOT-ACTIVE,S2=ACTIVE'" json:"status"`
	Deleted   bool      `gorm:"default:false" json:"deleted"`
	CreatedAt time.Time `gorm:"type:timestamp;autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"type:timestamp" json:"updated_at"`
}

func (Users) TableName() string {
	return "users" // name of table will created
}
