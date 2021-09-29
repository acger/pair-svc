package model

import (
	"gorm.io/gorm"
)

type Element struct {
	gorm.Model
	Uid  uint64 `gorm:"index:id_mode,priority:1"`
	Name string `gorm:"size:60;index:name_mode,priority:1"`
	Mode string `gorm:"size:10;index:id_mode,priority:2;index:name_mode,priority:2"`
	Star bool
	Sort int64
}
