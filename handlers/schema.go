package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/myidapp/backend/models"
	"net/http"
)

func GetSchema(db *gorm.DB, w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)

	uuid, ok := vars["uuid"]
	if !ok {
		ReturnError(w, http.StatusBadRequest, "Bad request")
		return
	}

	var dbSchema models.Schema
	db.First(&dbSchema, "uuid = ?", uuid)

	if dbSchema.ID == 0 {
		ReturnError(w, http.StatusNotFound, "Not found")
		return
	}

	ReturnResult(w, dbSchema)
}

func PostSchema(db *gorm.DB, w http.ResponseWriter, req *http.Request) {

	var schema models.Schema
	var dbSchema models.Schema
	err := json.NewDecoder(req.Body).Decode(&schema)
	if err != nil {
		ReturnError(w, http.StatusBadRequest, "Bad request")
		return
	}

	if !db.Where("uuid = ?", schema.UUID).First(&dbSchema).RecordNotFound() {
		ReturnError(w, http.StatusForbidden, "Already exists")
		return
	}

	db.Create(&schema)

	ReturnResult(w, true)
}
