package handlers

import (
	"GO_gRPC/db"

	"gorm.io/gorm"
)

type DBHandler struct {
	DB *gorm.DB
}

func NewDBHandler() DBHandler {

	database := db.NewDB()

	return DBHandler{database}

}
