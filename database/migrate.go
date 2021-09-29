package database

import (
	"github.com/acger/pair-svc/model"
	"gorm.io/gorm"
)

func migrate(db *gorm.DB) {
	db.AutoMigrate(model.Element{})
}
