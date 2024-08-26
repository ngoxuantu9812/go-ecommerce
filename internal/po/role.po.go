package po

import (
	"gorm.io/gorm"
)

type Role struct {
	gorm.Model
	ID       int64  `gorm:"column:id, primary_key;type:int;not null,idx_uuid unique, autoincrement"`
	RoleName string `gorm:"column:role_name; not null"`
	IsActive bool   `gorm:"column:is_active; type:bool;"`
	RoleNote string `gorm:"column:role_note;"`
}

func (Role) TableName() string {
	return "roles"
}
