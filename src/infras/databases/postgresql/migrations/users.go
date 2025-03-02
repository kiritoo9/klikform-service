package migrations

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Users struct {
	ID       uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	Email    string    `gorm:"uniqueIndex" json:"email"`
	Password string    `gorm:"not null" json:"password"`
}

func (Users) TableName() string {
	return "users" // name of table will created
}

func (user *Users) BeforeCreate(tx *gorm.DB) (err error) {
	user.ID = uuid.New()
	return
}
