package models

import (
	"github.com/jinzhu/gorm"
)

type Signature struct {
	gorm.Model
	UUID       string    `json:"uuid"`
	Signature  string    `json:"signature"`
	PublicKey  string 	 `json:"public_key"`
}
