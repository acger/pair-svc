package model

import (
	"gorm.io/gorm"
)

type Element struct {
	gorm.Model
	Uid  uint64 `gorm:"index"`
	Skill string `gorm:"size:6000;default:'[]';"`
	SkillNeed string `gorm:"size:6000;default:'[]';"`
	Star bool
	Boost int64
}
