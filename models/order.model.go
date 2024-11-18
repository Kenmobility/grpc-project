package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Order struct {
	ID        uint   `json:"id" gorm:"column:id;PRIMARY_KEY;type:integer;autoIncrement"`
	PublicID  string `gorm:"type:uuid;uniqueIndex"`
	Status    string
	UserId    uint  `json:"user_id"`
	User      *User `json:"user" gorm:"foreignkey:UserId"`
	CreatedAt time.Time
}

func (p *Order) BeforeCreate(tx *gorm.DB) (err error) {
	p.PublicID = uuid.New().String()
	return
}
