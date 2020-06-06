package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/myidapp/backend/models"
	"net/http"
)

func Clear(db *gorm.DB, w http.ResponseWriter, req *http.Request) {
	db.DropTableIfExists("signatures")
	db.AutoMigrate( new(models.Signature))
	ReturnResult(w, true)
}

func GetSign(db *gorm.DB, w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)

	uuid, ok := vars["uuid"]
	if !ok {
		ReturnError(w, http.StatusBadRequest, "Bad request")
		return
	}

	var dbSignature models.Signature
	db.First(&dbSignature, "uuid = ?", uuid)

	if dbSignature.ID == 0 {
		ReturnError(w, http.StatusNotFound, "Not found")
		return
	}


	ReturnResult(w, dbSignature)
}

func PostSign(db *gorm.DB, w http.ResponseWriter, req *http.Request) {

	var signature models.Signature
	var dbSignature models.Signature
	err := json.NewDecoder(req.Body).Decode(&signature)
	if err != nil {
		ReturnError(w, http.StatusBadRequest, "Bad request")
		return
	}


	if !db.Where("uuid = ?", signature.UUID).First(&dbSignature).RecordNotFound(){
		ReturnError(w, http.StatusForbidden, "Already exists")
		return
	}

	db.Create(&signature)

	ReturnResult(w, true)
}

