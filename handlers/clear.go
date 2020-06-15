package handlers

import (
	"github.com/jinzhu/gorm"
	"github.com/myidapp/backend/models"
	"net/http"
)

func Clear(db *gorm.DB, w http.ResponseWriter, req *http.Request) {
	db.DropTableIfExists("signatures")
	db.DropTableIfExists("schemas")
	db.AutoMigrate(new(models.Signature))
	db.AutoMigrate(new(models.Schema))
	ReturnResult(w, true)
}
