package po

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UUID     uuid.UUID `gorm:"column:uuid;type:varchar(255);not null,idx_uuid unique"`
	Username string    `gorm:"column:user_name;"`
	Email    string    `gorm:"column:email;"`
	Password string    `gorm:"column:password;"`
	IsActive bool      `gorm:"column:is_active;"`
	Roles    []Role    `gorm:"many2many:user_roles;"`
}

func (User) TableName() string {
	return "users"
}
