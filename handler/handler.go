package handler

import (
	"github.com/ramown/gopportunites/config"
	"gorm.io/gorm"
)

var (
	logger *config.Logger
	db     *gorm.DB
)

func InitializeHanlder() {
	logger = config.NewLogger("handler")
	db = config.GetSQLite()
}
