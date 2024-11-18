package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID        uint   `json:"id" gorm:"column:id;PRIMARY_KEY;type:integer;autoIncrement"`
	PublicID  string `gorm:"type:uuid;uniqueIndex"`
	FullName  string
	Email     string
	Password  string
	CreatedAt time.Time
}

func (p *User) BeforeCreate(tx *gorm.DB) (err error) {
	p.PublicID = uuid.New().String()
	return
}
