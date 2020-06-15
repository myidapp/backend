package models

import "github.com/jinzhu/gorm"

type Schema struct {
	gorm.Model
	UUID       string `json:"uuid"`
	VendorName string `json:"vendor_name"`
	LogoURL    string `json:"logo_url"`
	Callback   string `json:"callback"`
	Fields     string `json:"fields"`
}
